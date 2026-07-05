package backend

import (
	"os"
	"path/filepath"
	"strings"
)

type LTemporaryResult struct {
	LTemporaryPath  string `json:"LTemporaryPath"`
	LTemporaryCount int    `json:"LTemporaryCount"`
}

func LTemporaryRootRead(preference LPreference) string {
	root := strings.TrimSpace(preference.LPreferenceTemporary)
	if root == "" {
		root = os.TempDir()
	}

	return filepath.Join(root, "VideoMergerPreview")
}

func LTemporaryCreate(preference LPreference, pattern string) (*os.File, error) {
	root := LTemporaryRootRead(preference)
	if err := os.MkdirAll(root, 0o755); err != nil {
		return nil, err
	}

	return os.CreateTemp(root, pattern)
}

func LTemporaryClean(preference LPreference) (LTemporaryResult, error) {
	root := LTemporaryRootRead(preference)
	result := LTemporaryResult{LTemporaryPath: root}
	entries, err := os.ReadDir(root)
	if os.IsNotExist(err) {
		return result, nil
	}
	if err != nil {
		return result, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !LTemporaryFileCheck(name) {
			continue
		}

		if err := os.Remove(filepath.Join(root, name)); err == nil {
			result.LTemporaryCount++
		}
	}

	return result, nil
}

func LTemporaryFileCheck(name string) bool {
	return strings.HasSuffix(name, ".mp4") || strings.HasPrefix(name, "video-merger-list-")
}
