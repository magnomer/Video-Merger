package backend

import (
	"context"
	"os"
	"path/filepath"
	"strings"
)

func LClipResolve(LRuntimeContext context.Context, inputPaths []string, includeSubfolders bool, marker LMarker) ([]LClip, error) {
	mediaFiles := []LClip{}
	seen := map[string]bool{}

	for _, inputPath := range inputPaths {
		if LRuntimeContext.Err() != nil {
			return nil, LRuntimeContext.Err()
		}

		inputPath = strings.TrimSpace(inputPath)
		if inputPath == "" {
			continue
		}

		cleanInputPath := filepath.Clean(inputPath)
		info, err := os.Stat(cleanInputPath)
		if err != nil {
			return nil, err
		}

		resolvedFiles, err := LClipInputResolve(LRuntimeContext, cleanInputPath, info.IsDir(), includeSubfolders, seen, marker)
		if err != nil {
			return nil, err
		}

		mediaFiles = append(mediaFiles, resolvedFiles...)
	}

	return mediaFiles, nil
}
