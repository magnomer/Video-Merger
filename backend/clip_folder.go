package backend

import (
	"context"
	"path/filepath"
)

func LClipFolderResolve(
	LRuntimeContext context.Context,
	cleanInputPath string,
	includeSubfolders bool,
	seen map[string]bool,
	marker LMarker,
) ([]LClip, error) {
	folderFiles, err := LDirectoryLocalScan(LRuntimeContext, cleanInputPath, includeSubfolders)
	if err != nil {
		return nil, err
	}

	parentOfSelectedFolder := filepath.Dir(cleanInputPath)
	mediaFiles := []LClip{}

	for _, filePath := range folderFiles {
		if LRuntimeContext.Err() != nil {
			return nil, LRuntimeContext.Err()
		}

		normalizedPath := filepath.Clean(filePath)
		relativeDir, err := filepath.Rel(parentOfSelectedFolder, filepath.Dir(normalizedPath))
		if err != nil || relativeDir == "." {
			relativeDir = ""
		}

		file, ok, err := LClipFileResolve(LRuntimeContext, normalizedPath, relativeDir, seen, marker)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		mediaFiles = append(mediaFiles, file)
	}

	return mediaFiles, nil
}
