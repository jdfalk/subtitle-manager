package renamer

import (
	"fmt"
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
	if len(files) > 1 {
		return fmt.Errorf("multiple subtitle files for %s", videoPath)
	}
	if len(files) == 0 {
		return nil
	}
	newName := base + "." + lang + ".srt"
	if files[0] == newName {
		return nil
	}
	return os.Rename(files[0], newName)
}
