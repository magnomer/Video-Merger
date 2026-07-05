package backend

import (
	"fmt"
	"path/filepath"
	"strings"
)

func LConcatCreate(options LPreference, group LBatch) (string, error) {
	tempFile, err := LTemporaryCreate(options, "video-merger-list-*.txt")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	for _, file := range group.LBatchClip {
		escapedPath := LConcatEscape(file.LClipPath)

		_, err := tempFile.WriteString(fmt.Sprintf("file '%s'\n", escapedPath))
		if err != nil {
			return "", err
		}
	}

	return tempFile.Name(), nil
}

func LConcatEscape(path string) string {
	normalizedPath := filepath.ToSlash(path)
	return strings.ReplaceAll(normalizedPath, "'", "'\\''")
}
