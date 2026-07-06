package backend

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var (
	LAssetPreviewLock        = make(chan struct{}, 1)
	lAssetPreviewTaskLock    sync.Mutex
	lAssetPreviewTaskCounter uint64
	lAssetPreviewTaskList    []LAssetPreviewTask
	lAssetPreviewStopMap     = map[string]struct{}{}
)

type LAssetPreviewTask struct {
	LAssetPreviewKey     uint64
	LAssetPreviewSession string
	LAssetPreviewCancel  context.CancelFunc
	LAssetPreviewDone    chan struct{}
}

type LAssetPreviewMarker struct {
	LAssetSourcePath string `json:"LAssetSourcePath"`
	LAssetSourceSize int64  `json:"LAssetSourceSize"`
	LAssetSourceTime int64  `json:"LAssetSourceTime"`
}

func LAssetPreviewResolve(ctx context.Context, session string, sourcePath string, sourceInfo os.FileInfo) (string, error) {
	preference, _ := LPreferenceLoad()
	cachePath, err := LAssetPreviewPathRead(preference, sourcePath, sourceInfo)
	if err != nil {
		return "", err
	}

	if LAssetPreviewCacheCheck(cachePath, sourcePath, sourceInfo) {
		return cachePath, nil
	}

	previewContext, cancel := context.WithCancel(ctx)
	previewStop := LAssetPreviewStart(session, cancel)
	defer previewStop()

	previewUnlock, err := LAssetPreviewLockSet(previewContext)
	if err != nil {
		return "", err
	}
	defer previewUnlock()

	if err := previewContext.Err(); err != nil {
		return "", err
	}

	if LAssetPreviewCacheCheck(cachePath, sourcePath, sourceInfo) {
		return cachePath, nil
	}

	if err := os.MkdirAll(filepath.Dir(cachePath), 0o755); err != nil {
		return "", err
	}

	if err := LAssetPreviewCacheRemove(cachePath); err != nil {
		return "", err
	}

	temporaryFile, err := os.CreateTemp(filepath.Dir(cachePath), strings.TrimSuffix(filepath.Base(cachePath), ".mp4")+".building-*.mp4")
	if err != nil {
		return "", err
	}
	temporaryPath := temporaryFile.Name()
	if err := temporaryFile.Close(); err != nil {
		os.Remove(temporaryPath)
		return "", err
	}
	defer os.Remove(temporaryPath)

	ffmpegPath, err := LCommandFFmpegRead(preference)
	if err != nil {
		return "", err
	}

	cmd := exec.CommandContext(
		previewContext,
		ffmpegPath,
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
		if previewContext.Err() != nil {
			return "", previewContext.Err()
		}

		return "", fmt.Errorf("preview conversion failed: %v %s", err, strings.TrimSpace(string(output)))
	}

	if err := previewContext.Err(); err != nil {
		return "", err
	}

	if !LFileRegularCheck(temporaryPath) {
		return "", fmt.Errorf("preview conversion did not create a regular file")
	}
	if err := os.Rename(temporaryPath, cachePath); err != nil {
		return "", err
	}
	if err := LAssetPreviewMarkerWrite(cachePath, sourcePath, sourceInfo); err != nil {
		os.Remove(cachePath)
		return "", err
	}

	return cachePath, nil
}

func LAssetPreviewStart(session string, cancel context.CancelFunc) func() {
	done := make(chan struct{})

	lAssetPreviewTaskLock.Lock()
	lAssetPreviewTaskCounter++
	key := lAssetPreviewTaskCounter
	lAssetPreviewTaskList = append(lAssetPreviewTaskList, LAssetPreviewTask{
		LAssetPreviewKey:     key,
		LAssetPreviewSession: session,
		LAssetPreviewCancel:  cancel,
		LAssetPreviewDone:    done,
	})
	_, stopped := lAssetPreviewStopMap[session]
	lAssetPreviewTaskLock.Unlock()

	if stopped {
		cancel()
	}

	return func() {
		lAssetPreviewTaskLock.Lock()
		for index, item := range lAssetPreviewTaskList {
			if item.LAssetPreviewKey == key {
				lAssetPreviewTaskList = append(lAssetPreviewTaskList[:index], lAssetPreviewTaskList[index+1:]...)
				break
			}
		}
		lAssetPreviewTaskLock.Unlock()

		close(done)
	}
}

func LAssetPreviewStop(session string) {
	lAssetPreviewTaskLock.Lock()
	if session != "" {
		lAssetPreviewStopMap[session] = struct{}{}
	}
	items := append([]LAssetPreviewTask(nil), lAssetPreviewTaskList...)
	lAssetPreviewTaskLock.Unlock()

	for _, item := range items {
		if session == "" || item.LAssetPreviewSession == session {
			item.LAssetPreviewCancel()
		}
	}
}

func LAssetPreviewLockSet(ctx context.Context) (func(), error) {
	select {
	case LAssetPreviewLock <- struct{}{}:
		return func() { <-LAssetPreviewLock }, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func LAssetPreviewPathRead(preference LPreference, sourcePath string, sourceInfo os.FileInfo) (string, error) {
	absolutePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return "", err
	}

	return LAssetPreviewPathByMarkerRead(preference, LAssetPreviewMarker{
		LAssetSourcePath: absolutePath,
		LAssetSourceSize: sourceInfo.Size(),
		LAssetSourceTime: sourceInfo.ModTime().UnixNano(),
	})
}

func LAssetPreviewPathByMarkerRead(preference LPreference, marker LAssetPreviewMarker) (string, error) {
	absolutePath, err := filepath.Abs(marker.LAssetSourcePath)
	if err != nil {
		return "", err
	}
	if absolutePath != marker.LAssetSourcePath {
		return "", fmt.Errorf("preview marker source path is not absolute")
	}
	if marker.LAssetSourceSize < 0 || marker.LAssetSourceTime == 0 {
		return "", fmt.Errorf("preview marker source metadata is invalid")
	}

	seed := LAssetPreviewSeedRead(marker)
	sum := sha1.Sum([]byte(seed))
	name := hex.EncodeToString(sum[:]) + ".mp4"

	return filepath.Join(LTemporaryRootRead(preference), name), nil
}

func LAssetPreviewSeedRead(marker LAssetPreviewMarker) string {
	return fmt.Sprintf("%s\x00%d\x00%d", marker.LAssetSourcePath, marker.LAssetSourceSize, marker.LAssetSourceTime)
}

func LAssetPreviewMarkerPathRead(path string) string {
	return path + ".json"
}

func LAssetPreviewMarkerWrite(cachePath string, sourcePath string, sourceInfo os.FileInfo) error {
	absolutePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return err
	}

	marker := LAssetPreviewMarker{
		LAssetSourcePath: absolutePath,
		LAssetSourceSize: sourceInfo.Size(),
		LAssetSourceTime: sourceInfo.ModTime().UnixNano(),
	}
	data, err := json.Marshal(marker)
	if err != nil {
		return err
	}

	markerPath := LAssetPreviewMarkerPathRead(cachePath)
	return os.WriteFile(markerPath, data, 0o600)
}

func LAssetPreviewCacheCheck(cachePath string, sourcePath string, sourceInfo os.FileInfo) bool {
	if !LFileRegularCheck(cachePath) {
		return false
	}

	markerPath := LAssetPreviewMarkerPathRead(cachePath)
	if !LFileRegularCheck(markerPath) {
		return false
	}

	data, err := os.ReadFile(markerPath)
	if err != nil {
		return false
	}

	var marker LAssetPreviewMarker
	if err := json.Unmarshal(data, &marker); err != nil {
		return false
	}

	absolutePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return false
	}

	return marker.LAssetSourcePath == absolutePath &&
		marker.LAssetSourceSize == sourceInfo.Size() &&
		marker.LAssetSourceTime == sourceInfo.ModTime().UnixNano()
}

func LAssetPreviewCacheRemove(cachePath string) error {
	if err := LFileRemoveIfExists(cachePath); err != nil {
		return err
	}
	return LFileRemoveIfExists(LAssetPreviewMarkerPathRead(cachePath))
}

func LAssetPreviewCacheCheckByPath(cachePath string) bool {
	if !LFileRegularCheck(cachePath) {
		return false
	}
	markerPath := LAssetPreviewMarkerPathRead(cachePath)
	if !LFileRegularCheck(markerPath) {
		return false
	}

	data, err := os.ReadFile(markerPath)
	if err != nil {
		return false
	}

	var marker LAssetPreviewMarker
	if err := json.Unmarshal(data, &marker); err != nil {
		return false
	}
	preference, _ := LPreferenceLoad()
	expectedPath, err := LAssetPreviewPathByMarkerRead(preference, marker)
	if err != nil || expectedPath != cachePath {
		return false
	}

	return true
}
