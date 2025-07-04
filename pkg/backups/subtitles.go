// file: pkg/backups/subtitles.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174011

package backups

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/security"
	"github.com/spf13/viper"
)

// SubtitleBackupData represents subtitle file backup information.
type SubtitleBackupData struct {
	Files    []SubtitleFileInfo        `json:"files"`
	Metadata []database.SubtitleRecord `json:"metadata"`
	Created  time.Time                 `json:"created"`
}

// SubtitleFileInfo contains information about a backed up subtitle file.
type SubtitleFileInfo struct {
	RelativePath string    `json:"relative_path"`
	OriginalPath string    `json:"original_path"`
	Size         int64     `json:"size"`
	ModTime      time.Time `json:"mod_time"`
	Hash         string    `json:"hash"`
}

// SubtitleBackupper provides subtitle file backup and restore functionality.
type SubtitleBackupper struct {
	store database.SubtitleStore
}

// NewSubtitleBackupper creates a new subtitle backup instance.
func NewSubtitleBackupper(store database.SubtitleStore) *SubtitleBackupper {
	return &SubtitleBackupper{
		store: store,
	}
}

// CreateSubtitleBackup creates a backup of subtitle files and their metadata.
func (sb *SubtitleBackupper) CreateSubtitleBackup(ctx context.Context, includePaths []string) ([]byte, error) {
	backupData := &SubtitleBackupData{
		Files:   []SubtitleFileInfo{},
		Created: time.Now(),
	}

	// Get subtitle metadata from database
	subtitles, err := sb.store.ListSubtitles()
	if err != nil {
		return nil, fmt.Errorf("failed to get subtitle metadata: %w", err)
	}
	backupData.Metadata = subtitles

	// Create tar archive with subtitle files
	var archiveData []byte
	if len(includePaths) > 0 {
		archiveData, err = sb.createSubtitleArchive(ctx, includePaths, &backupData.Files)
		if err != nil {
			return nil, fmt.Errorf("failed to create subtitle archive: %w", err)
		}
	}

	// Create backup metadata
	metadataBytes, err := json.MarshalIndent(backupData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to serialize backup metadata: %w", err)
	}

	// Combine metadata and archive into a single backup
	backup := map[string]interface{}{
		"metadata": string(metadataBytes),
		"archive":  archiveData,
	}

	result, err := json.Marshal(backup)
	if err != nil {
		return nil, fmt.Errorf("failed to create backup: %w", err)
	}

	return result, nil
}

// RestoreSubtitleBackup restores subtitle files and metadata from backup.
func (sb *SubtitleBackupper) RestoreSubtitleBackup(ctx context.Context, data []byte, restorePath string) error {
	var backup map[string]interface{}
	if err := json.Unmarshal(data, &backup); err != nil {
		return fmt.Errorf("failed to parse backup: %w", err)
	}

	// Restore metadata
	if metadataStr, ok := backup["metadata"].(string); ok {
		var backupData SubtitleBackupData
		if err := json.Unmarshal([]byte(metadataStr), &backupData); err != nil {
			return fmt.Errorf("failed to parse backup metadata: %w", err)
		}

		// Restore subtitle records to database
		for _, subtitle := range backupData.Metadata {
			if err := sb.store.InsertSubtitle(&subtitle); err != nil {
				// Log error but continue with other records
				fmt.Printf("Warning: failed to restore subtitle record %s: %v\n", subtitle.File, err)
			}
		}
	}

	// Restore files if archive exists and restore path is provided
	if archiveData, ok := backup["archive"]; ok && restorePath != "" {
		if archiveBytes, ok := archiveData.([]byte); ok && len(archiveBytes) > 0 {
			if err := sb.extractSubtitleArchive(ctx, archiveBytes, restorePath); err != nil {
				return fmt.Errorf("failed to restore subtitle files: %w", err)
			}
		}
	}

	return nil
}

// createSubtitleArchive creates a tar.gz archive of subtitle files.
func (sb *SubtitleBackupper) createSubtitleArchive(ctx context.Context, includePaths []string, fileInfos *[]SubtitleFileInfo) ([]byte, error) {
	// Create tar.gz archive in memory
	var buf bytes.Buffer
	gzw := gzip.NewWriter(&buf)
	defer gzw.Close()

	tw := tar.NewWriter(gzw)
	defer tw.Close()

	for _, basePath := range includePaths {
		err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip directories and non-subtitle files
			if info.IsDir() || !isSubtitleFile(path) {
				return nil
			}

			// Calculate relative path
			relPath, err := filepath.Rel(basePath, path)
			if err != nil {
				relPath = filepath.Base(path)
			}

			// Read file
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("failed to open file %s: %w", path, err)
			}
			defer file.Close()

			// Calculate hash
			hash, err := calculateFileHash(file)
			if err != nil {
				return fmt.Errorf("failed to calculate hash for %s: %w", path, err)
			}

			// Reset file pointer
			file.Seek(0, 0)

			// Add to tar archive
			header := &tar.Header{
				Name:    relPath,
				Size:    info.Size(),
				Mode:    int64(info.Mode()),
				ModTime: info.ModTime(),
			}

			if err := tw.WriteHeader(header); err != nil {
				return fmt.Errorf("failed to write header for %s: %w", path, err)
			}

			if _, err := io.Copy(tw, file); err != nil {
				return fmt.Errorf("failed to write file %s: %w", path, err)
			}

			// Add to file info list
			*fileInfos = append(*fileInfos, SubtitleFileInfo{
				RelativePath: relPath,
				OriginalPath: path,
				Size:         info.Size(),
				ModTime:      info.ModTime(),
				Hash:         hash,
			})

			return nil
		})

		if err != nil {
			return nil, fmt.Errorf("failed to walk directory %s: %w", basePath, err)
		}
	}

	if err := tw.Close(); err != nil {
		return nil, fmt.Errorf("failed to close tar writer: %w", err)
	}

	if err := gzw.Close(); err != nil {
		return nil, fmt.Errorf("failed to close gzip writer: %w", err)
	}

	return buf.Bytes(), nil
}

// extractSubtitleArchive extracts a tar.gz archive of subtitle files.
func (sb *SubtitleBackupper) extractSubtitleArchive(ctx context.Context, archiveData []byte, restorePath string) error {
	sanitizedRestore, err := security.ValidateAndSanitizePath(restorePath)
	if err != nil {
		return fmt.Errorf("invalid restore path: %w", err)
	}

	// Create restore directory
	if err := os.MkdirAll(sanitizedRestore, 0755); err != nil {
		return fmt.Errorf("failed to create restore directory: %w", err)
	}

	// Open gzip reader
	gzr, err := gzip.NewReader(bytes.NewReader(archiveData))
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzr.Close()

	// Open tar reader
	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar header: %w", err)
		}

		// Create sanitized full path and ensure it stays within the restore directory
		fullPath, err := security.ValidateAndSanitizePath(filepath.Join(sanitizedRestore, header.Name))
		if err != nil || !strings.HasPrefix(fullPath, sanitizedRestore+string(os.PathSeparator)) {
			return fmt.Errorf("invalid file path in archive: %s", header.Name)
		}

		// Create directory if needed
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return fmt.Errorf("failed to create directory for %s: %w", fullPath, err)
		}

		// Extract file
		file, err := os.Create(fullPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", fullPath, err)
		}

		if _, err := io.Copy(file, tr); err != nil {
			file.Close()
			return fmt.Errorf("failed to extract file %s: %w", fullPath, err)
		}

		// Set file permissions and modification time
		if err := file.Chmod(os.FileMode(header.Mode)); err != nil {
			file.Close()
			return fmt.Errorf("failed to set permissions for %s: %w", fullPath, err)
		}

		if err := file.Close(); err != nil {
			return fmt.Errorf("failed to close file %s: %w", fullPath, err)
		}

		if err := os.Chtimes(fullPath, header.ModTime, header.ModTime); err != nil {
			return fmt.Errorf("failed to set times for %s: %w", fullPath, err)
		}
	}

	return nil
}

// isSubtitleFile checks if a file is a subtitle file based on extension.
func isSubtitleFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	subtitleExts := []string{".srt", ".vtt", ".ass", ".ssa", ".sub", ".idx", ".sup"}

	for _, validExt := range subtitleExts {
		if ext == validExt {
			return true
		}
	}

	return false
}

// calculateFileHash calculates SHA-256 hash of a file.
func calculateFileHash(file *os.File) (string, error) {
	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

// GetSubtitleBackupPaths returns the configured paths for subtitle backup.
func GetSubtitleBackupPaths() []string {
	pathsStr := viper.GetString("backup_subtitle_paths")
	if pathsStr == "" {
		return []string{}
	}

	// Split by comma and trim spaces
	paths := strings.Split(pathsStr, ",")
	for i, path := range paths {
		paths[i] = strings.TrimSpace(path)
	}

	return paths
}
