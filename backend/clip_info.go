package backend

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var LClipSupportedExtensionMap = map[string]bool{
	".mp4":  true,
	".m4v":  true,
	".mov":  true,
	".3gp":  true,
	".3g2":  true,
	".mkv":  true,
	".webm": true,
	".avi":  true,
	".ts":   true,
	".mts":  true,
	".m2ts": true,
	".mpeg": true,
	".mpg":  true,
	".vob":  true,
	".flv":  true,
	".f4v":  true,
	".ogv":  true,
	".ogg":  true,
	".mxf":  true,
	".wmv":  true,
	".asf":  true,
	".wtv":  true,
	".nut":  true,
	".dv":   true,
}

var LClipSupportedExtensionList = []string{
	".mp4",
	".m4v",
	".mov",
	".3gp",
	".3g2",
	".mkv",
	".webm",
	".avi",
	".ts",
	".mts",
	".m2ts",
	".mpeg",
	".mpg",
	".vob",
	".flv",
	".f4v",
	".ogv",
	".ogg",
	".mxf",
	".wmv",
	".asf",
	".wtv",
	".nut",
	".dv",
}

func LClipCheck(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))
	return LClipSupportedExtensionMap[extension]
}

func LClipSet(LRuntimeContext context.Context, preference LPreference, file *LClip) error {
	info, err := os.Stat(file.LClipPath)
	if err == nil {
		file.LMetricSize = info.Size()
	}

	probe, err := LProbeRun(LRuntimeContext, preference, file.LClipPath)
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
