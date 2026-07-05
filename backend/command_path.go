package backend

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func LCommandFFmpegRead(preference LPreference) (string, error) {
	path := strings.TrimSpace(preference.LPreferenceFFmpeg)
	if path == "" {
		return "ffmpeg", nil
	}

	return LCommandExecutableResolve(path, "ffmpeg")
}

func LCommandFFprobeRead(preference LPreference) (string, error) {
	ffmpegPath := strings.TrimSpace(preference.LPreferenceFFmpeg)
	if ffmpegPath == "" {
		return "ffprobe", nil
	}

	resolvedFFmpeg, err := LCommandExecutableResolve(ffmpegPath, "ffmpeg")
	if err != nil {
		return "", err
	}

	folder := filepath.Dir(resolvedFFmpeg)
	candidate := filepath.Join(folder, LCommandExecutableNameRead("ffprobe"))
	return LCommandExecutableResolve(candidate, "ffprobe")
}

func LCommandExecutableResolve(path string, expected string) (string, error) {
	cleanPath := filepath.Clean(strings.TrimSpace(path))
	if cleanPath == "." || cleanPath == "" {
		return "", fmt.Errorf("%s path is empty", expected)
	}

	if !LCommandExecutableNameCheck(cleanPath, expected) {
		return "", fmt.Errorf("configured executable must be %s", LCommandExecutableNameRead(expected))
	}

	info, err := os.Lstat(cleanPath)
	if err != nil {
		return "", err
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return "", fmt.Errorf("configured %s path must not be a symlink", expected)
	}
	if info.IsDir() {
		return "", fmt.Errorf("configured %s path is a folder", expected)
	}
	if !info.Mode().IsRegular() {
		return "", fmt.Errorf("configured %s path is not a regular file", expected)
	}
	if runtime.GOOS != "windows" && info.Mode()&0o111 == 0 {
		return "", fmt.Errorf("configured %s path is not executable", expected)
	}

	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return "", err
	}

	return absPath, nil
}

func LCommandExecutableNameRead(name string) string {
	if runtime.GOOS == "windows" {
		return name + ".exe"
	}

	return name
}

func LCommandExecutableNameCheck(path string, expected string) bool {
	base := strings.ToLower(filepath.Base(path))
	if base == expected {
		return true
	}

	return base == expected+".exe"
}
