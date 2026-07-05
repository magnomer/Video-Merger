package backend

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var LAssetPreviewLock sync.Mutex

func LAssetPreviewResolve(sourcePath string, sourceInfo os.FileInfo) (string, error) {
	preference, _ := LPreferenceLoad()
	cachePath, err := LAssetPreviewPathRead(preference, sourcePath, sourceInfo)
	if err != nil {
		return "", err
	}

	if LAssetFileCheck(cachePath) {
		return cachePath, nil
	}

	LAssetPreviewLock.Lock()
	defer LAssetPreviewLock.Unlock()

	if LAssetFileCheck(cachePath) {
		return cachePath, nil
	}

	if err := os.MkdirAll(filepath.Dir(cachePath), 0o755); err != nil {
		return "", err
	}

	temporaryPath := strings.TrimSuffix(cachePath, ".mp4") + ".building.mp4"
	os.Remove(temporaryPath)
	defer os.Remove(temporaryPath)

	cmd := exec.Command(
		LCommandFFmpegRead(preference),
		"-y",
		"-hide_banner",
		"-loglevel", "error",
		"-i", sourcePath,
		"-map", "0:v:0?",
		"-map", "0:a:0?",
		"-vf", "format=yuv420p",
		"-c:v", "libx264",
		"-preset", "veryfast",
		"-crf", "23",
		"-c:a", "aac",
		"-b:a", "160k",
		"-movflags", "+faststart",
		temporaryPath,
	)
	LCommandHide(cmd)

	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("preview conversion failed: %v %s", err, strings.TrimSpace(string(output)))
	}

	if err := os.Rename(temporaryPath, cachePath); err != nil {
		return "", err
	}

	return cachePath, nil
}

func LAssetPreviewPathRead(preference LPreference, sourcePath string, sourceInfo os.FileInfo) (string, error) {
	absolutePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return "", err
	}

	seed := fmt.Sprintf("%s\x00%d\x00%d", absolutePath, sourceInfo.Size(), sourceInfo.ModTime().UnixNano())
	sum := sha1.Sum([]byte(seed))
	name := hex.EncodeToString(sum[:]) + ".mp4"

	return filepath.Join(LTemporaryRootRead(preference), name), nil
}

func LAssetFileCheck(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir() && info.Size() > 0
}
