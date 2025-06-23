package renamer

import (
	"os"
	"path/filepath"
	"strings"
)

// Rename updates subtitle filenames so they match the given video path.
// It searches the video directory for files named "*.{lang}.srt" and
// renames them to use the video base name. Existing matching files are
// left untouched.
func Rename(videoPath, lang string) error {
	base := strings.TrimSuffix(videoPath, filepath.Ext(videoPath))
	dir := filepath.Dir(videoPath)
	pattern := filepath.Join(dir, "*."+lang+".srt")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}
	newName := base + "." + lang + ".srt"
	for _, f := range files {
		if f == newName {
			continue
		}
		if err := os.Rename(f, newName); err != nil {
			return err
		}
	}
	return nil
}
