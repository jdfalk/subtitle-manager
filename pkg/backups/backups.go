// Package backups provides utilities for managing comprehensive backup and restore operations.
// It supports creating, listing, and storing backup metadata for subtitle-manager.
//
// The package provides a comprehensive backup system supporting:
// - Database backups (all tables)
// - Configuration file backups
// - Subtitle file backups (optional)
// - Backup compression and encryption
// - Cloud storage integration
// - Backup scheduling and rotation
// - Selective restore capabilities

package backups

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Errors
var (
	ErrBackupNotFound    = errors.New("backup not found")
	ErrChecksumMismatch  = errors.New("backup checksum mismatch")
)

// BackupType represents the type of backup.
type BackupType string

const (
	BackupTypeDatabase      BackupType = "database"
	BackupTypeConfiguration BackupType = "configuration"
	BackupTypeSubtitles     BackupType = "subtitles"
	BackupTypeFull         BackupType = "full"
)

// Backup represents a comprehensive backup with metadata.
type Backup struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	CreatedAt   time.Time  `json:"created_at"`
	Size        int64      `json:"size"`
	Type        BackupType `json:"type"`
	Contents    []string   `json:"contents"`
	Compressed  bool       `json:"compressed"`
	Encrypted   bool       `json:"encrypted"`
	CloudStored bool       `json:"cloud_stored"`
	FilePath    string     `json:"file_path"`
	Checksum    string     `json:"checksum"`
}

// Storage interface abstracts backup storage operations.
type Storage interface {
	// Store saves backup data to storage and returns the stored path.
	Store(ctx context.Context, data []byte, filename string) (string, error)
	// Retrieve loads backup data from storage.
	Retrieve(ctx context.Context, path string) ([]byte, error)
	// Delete removes backup data from storage.
	Delete(ctx context.Context, path string) error
	// List returns all backup files in storage.
	List(ctx context.Context) ([]string, error)
}

// Compression interface abstracts compression operations.
type Compression interface {
	// Compress compresses the input data.
	Compress(data []byte) ([]byte, error)
	// Decompress decompresses the input data.
	Decompress(data []byte) ([]byte, error)
}

// Encryption interface abstracts encryption operations.
type Encryption interface {
	// Encrypt encrypts the input data.
	Encrypt(data []byte) ([]byte, error)
	// Decrypt decrypts the input data.
	Decrypt(data []byte) ([]byte, error)
}

// BackupManager provides comprehensive backup and restore operations.
type BackupManager struct {
	storage     Storage
	compression Compression
	encryption  Encryption
	mu          sync.RWMutex
	backups     map[string]*Backup
}

// NewBackupManager creates a new backup manager with the provided interfaces.
func NewBackupManager(storage Storage, compression Compression, encryption Encryption) *BackupManager {
	return &BackupManager{
		storage:     storage,
		compression: compression,
		encryption:  encryption,
		backups:     make(map[string]*Backup),
	}
}

// CreateBackup creates a new backup with the specified type and contents.
func (bm *BackupManager) CreateBackup(ctx context.Context, backupType BackupType, contents []string, data []byte) (*Backup, error) {
	backup := &Backup{
		ID:        generateBackupID(),
		Name:      generateBackupName(backupType),
		CreatedAt: time.Now(),
		Type:      backupType,
		Contents:  contents,
		Size:      int64(len(data)),
	}

	// Apply compression if available
	if bm.compression != nil {
		compressed, err := bm.compression.Compress(data)
		if err != nil {
			return nil, err
		}
		data = compressed
		backup.Compressed = true
		backup.Size = int64(len(data))
	}

	// Apply encryption if available
	if bm.encryption != nil {
		encrypted, err := bm.encryption.Encrypt(data)
		if err != nil {
			return nil, err
		}
		data = encrypted
		backup.Encrypted = true
		backup.Size = int64(len(data))
	}

	// Calculate checksum
	backup.Checksum = calculateChecksum(data)

	// Store backup
	path, err := bm.storage.Store(ctx, data, backup.Name)
	if err != nil {
		return nil, err
	}
	backup.FilePath = path

	// Store backup metadata
	bm.mu.Lock()
	bm.backups[backup.ID] = backup
	bm.mu.Unlock()

	return backup, nil
}

// ListBackups returns all available backups.
func (bm *BackupManager) ListBackups() []*Backup {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	
	backups := make([]*Backup, 0, len(bm.backups))
	for _, backup := range bm.backups {
		backups = append(backups, backup)
	}
	return backups
}

// GetBackup retrieves a specific backup by ID.
func (bm *BackupManager) GetBackup(id string) (*Backup, bool) {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	
	backup, exists := bm.backups[id]
	return backup, exists
}

// RestoreBackup restores data from a specific backup.
func (bm *BackupManager) RestoreBackup(ctx context.Context, id string) ([]byte, error) {
	backup, exists := bm.GetBackup(id)
	if !exists {
		return nil, ErrBackupNotFound
	}

	// Retrieve backup data
	data, err := bm.storage.Retrieve(ctx, backup.FilePath)
	if err != nil {
		return nil, err
	}

	// Verify checksum
	if calculateChecksum(data) != backup.Checksum {
		return nil, ErrChecksumMismatch
	}

	// Decrypt if necessary
	if backup.Encrypted && bm.encryption != nil {
		decrypted, err := bm.encryption.Decrypt(data)
		if err != nil {
			return nil, err
		}
		data = decrypted
	}

	// Decompress if necessary
	if backup.Compressed && bm.compression != nil {
		decompressed, err := bm.compression.Decompress(data)
		if err != nil {
			return nil, err
		}
		data = decompressed
	}

	return data, nil
}

// DeleteBackup removes a backup from storage.
func (bm *BackupManager) DeleteBackup(ctx context.Context, id string) error {
	backup, exists := bm.GetBackup(id)
	if !exists {
		return ErrBackupNotFound
	}

	// Delete from storage
	if err := bm.storage.Delete(ctx, backup.FilePath); err != nil {
		return err
	}

	// Remove from metadata
	bm.mu.Lock()
	delete(bm.backups, id)
	bm.mu.Unlock()

	return nil
}

// Helper functions
func generateBackupID() string {
	return uuid.New().String()
}

func generateBackupName(backupType BackupType) string {
	timestamp := time.Now().Format("20060102-150405")
	return fmt.Sprintf("%s-%s.bak", string(backupType), timestamp)
}

func calculateChecksum(data []byte) string {
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

// Global manager instance for backward compatibility
var defaultManager *BackupManager

// Initialize default manager with no-op implementations for backward compatibility
func init() {
	defaultManager = NewBackupManager(&noOpStorage{}, nil, nil)
}

// Backward compatibility functions - these maintain the original API

// Create simulates creating a new backup file and stores metadata in memory.
func Create() Backup {
	// For backward compatibility, create a simple backup with the old structure
	backup := Backup{
		ID:        generateBackupID(),
		Name:      time.Now().Format("20060102-150405") + ".bak",
		CreatedAt: time.Now(),
		Type:      BackupTypeDatabase,
		Size:      0,
		Contents:  []string{},
	}
	
	defaultManager.mu.Lock()
	defaultManager.backups[backup.ID] = &backup
	defaultManager.mu.Unlock()
	
	return backup
}

// List returns all known backups.
func List() []Backup {
	backups := defaultManager.ListBackups()
	result := make([]Backup, len(backups))
	for i, backup := range backups {
		result[i] = *backup
	}
	return result
}

// Restore pretends to restore the latest backup and returns its name.
func Restore() string {
	backups := defaultManager.ListBackups()
	if len(backups) == 0 {
		return ""
	}
	// Return the most recent backup
	var latest *Backup
	for _, backup := range backups {
		if latest == nil || backup.CreatedAt.After(latest.CreatedAt) {
			latest = backup
		}
	}
	return latest.Name
}

// No-op storage implementation for backward compatibility
type noOpStorage struct{}

func (s *noOpStorage) Store(ctx context.Context, data []byte, filename string) (string, error) {
	return filename, nil
}

func (s *noOpStorage) Retrieve(ctx context.Context, path string) ([]byte, error) {
	return []byte{}, nil
}

func (s *noOpStorage) Delete(ctx context.Context, path string) error {
	return nil
}

func (s *noOpStorage) List(ctx context.Context) ([]string, error) {
	return []string{}, nil
}
