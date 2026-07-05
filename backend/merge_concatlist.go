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
	success := false
	defer func() {
		tempFile.Close()
		if !success {
			LTemporaryOwnedRemove(tempFile.Name())
		}
	}()

	for _, file := range group.LBatchClip {
		cleanPath, err := LConcatPathResolve(file.LClipPath)
		if err != nil {
			return "", err
		}

		escapedPath := LConcatEscape(cleanPath)

		_, err = tempFile.WriteString(fmt.Sprintf("file '%s'\n", escapedPath))
		if err != nil {
			return "", err
		}
	}

	success = true
	return tempFile.Name(), nil
}

func LConcatEscape(path string) string {
	normalizedPath := filepath.ToSlash(path)
	return strings.ReplaceAll(normalizedPath, "'", "'\\''")
}

func LConcatPathResolve(path string) (string, error) {
	cleanPath := filepath.Clean(strings.TrimSpace(path))
	if cleanPath == "." || cleanPath == "" {
		return "", fmt.Errorf("concat input path is empty")
	}
	if strings.ContainsAny(cleanPath, "\r\n") {
		return "", fmt.Errorf("concat input path contains a line break")
	}
	if LConcatProtocolCheck(cleanPath) {
		return "", fmt.Errorf("concat input path must be a local file")
	}

	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return "", err
	}

	if !LFileRegularCheck(absPath) {
		return "", fmt.Errorf("concat input path is not a regular file")
	}

	return absPath, nil
}

func LConcatProtocolCheck(path string) bool {
	if filepath.VolumeName(path) != "" {
		return false
	}

	separator := strings.IndexAny(path, `/\`)
	colon := strings.Index(path, ":")
	return colon > 0 && (separator == -1 || colon < separator)
}
