package backend

import (
	"os"
	"path/filepath"
	"strings"
)

func LCommandFFmpegRead(preference LPreference) string {
	path := strings.TrimSpace(preference.LPreferenceFFmpeg)
	if path == "" {
		return "ffmpeg"
	}

	return path
}

func LCommandFFprobeRead(preference LPreference) string {
	ffmpegPath := strings.TrimSpace(preference.LPreferenceFFmpeg)
	if ffmpegPath == "" {
		return "ffprobe"
	}

	folder := filepath.Dir(ffmpegPath)
	name := "ffprobe"
	if strings.EqualFold(filepath.Ext(ffmpegPath), ".exe") {
		name = "ffprobe.exe"
	}

	candidate := filepath.Join(folder, name)
	info, err := os.Stat(candidate)
	if err == nil && !info.IsDir() {
		return candidate
	}

	return "ffprobe"
}
