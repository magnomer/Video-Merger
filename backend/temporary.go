package backend

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const LTemporaryOwnerText = "VideoMergerPreview\n"

var LTemporaryPreviewPattern = regexp.MustCompile(`^[a-f0-9]{40}\.mp4$`)

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

	file, err := os.CreateTemp(root, pattern)
	if err != nil {
		return nil, err
	}
	if err := LTemporaryOwnerWrite(file.Name()); err != nil {
		file.Close()
		os.Remove(file.Name())
		return nil, err
	}

	return file, nil
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
		path := filepath.Join(root, name)
		removed, err := LTemporaryPathClean(path, name)
		if err != nil {
			return result, err
		}
		if removed {
			result.LTemporaryCount++
		}
	}

	return result, nil
}

func LTemporaryPathClean(path string, name string) (bool, error) {
	if LTemporaryPreviewFileCheck(name) {
		return LTemporaryPreviewRemove(path)
	}
	if LTemporaryListFileCheck(name) && LTemporaryOwnerCheck(path) {
		return true, LTemporaryOwnedRemove(path)
	}

	return false, nil
}

func LTemporaryPreviewFileCheck(name string) bool {
	return LTemporaryPreviewPattern.MatchString(name)
}

func LTemporaryListFileCheck(name string) bool {
	return strings.HasPrefix(name, "video-merger-list-") && strings.HasSuffix(name, ".txt")
}

func LTemporaryOwnerPathRead(path string) string {
	return path + ".owner"
}

func LTemporaryOwnerWrite(path string) error {
	return os.WriteFile(LTemporaryOwnerPathRead(path), []byte(LTemporaryOwnerText), 0o600)
}

func LTemporaryOwnerCheck(path string) bool {
	if !LFileRegularCheck(path) || !LFileRegularCheck(LTemporaryOwnerPathRead(path)) {
		return false
	}
	data, err := os.ReadFile(LTemporaryOwnerPathRead(path))
	return err == nil && string(data) == LTemporaryOwnerText
}

func LTemporaryOwnedRemove(path string) error {
	if err := LFileRemoveIfExists(path); err != nil {
		return err
	}
	return LFileRemoveIfExists(LTemporaryOwnerPathRead(path))
}

func LTemporaryPreviewRemove(path string) (bool, error) {
	if !LAssetPreviewCacheCheckByPath(path) {
		return false, nil
	}
	return true, LAssetPreviewCacheRemove(path)
}
