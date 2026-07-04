package backend

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func LClipCheck(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))

	supportedExtensions := map[string]bool{
		".mp4": true,
		".mov": true,
		".mkv": true,
		".m4v": true,
	}

	return supportedExtensions[extension]
}

func LClipSet(LRuntimeContext context.Context, file *LClip) error {
	info, err := os.Stat(file.LClipPath)
	if err == nil {
		file.LMetricSize = info.Size()
	}

	probe, err := LProbeRun(LRuntimeContext, file.LClipPath)
	if err != nil {
		if LRuntimeContext.Err() != nil {
			return LRuntimeContext.Err()
		}

		return nil
	}

	duration, err := strconv.ParseFloat(probe.LProbeFormat.LMetricDuration, 64)
	if err == nil {
		file.LMetricDuration = duration
	}

	return nil
}
