package backend

import (
	"context"
	"os"
	"path/filepath"
)

func LDirectoryScan(folderPath string) ([]string, error) {
	return LDirectoryLocalScan(context.Background(), folderPath, false)
}

func LDirectoryTreeScan(folderPath string) ([]string, error) {
	return LDirectoryLocalScan(context.Background(), folderPath, true)
}

func LDirectoryLocalScan(LRuntimeContext context.Context, folderPath string, includeSubfolders bool) ([]string, error) {
	var paths []string

	if includeSubfolders {
		err := filepath.WalkDir(folderPath, func(path string, entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if LRuntimeContext.Err() != nil {
				return LRuntimeContext.Err()
			}

			if entry.IsDir() {
				return nil
			}

			if LClipCheck(path) {
				paths = append(paths, path)
			}

			return nil
		})

		if err != nil {
			return nil, err
		}

		return paths, nil
	}

	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if LRuntimeContext.Err() != nil {
			return nil, LRuntimeContext.Err()
		}

		if entry.IsDir() {
			continue
		}

		fullPath := filepath.Join(folderPath, entry.Name())

		if LClipCheck(fullPath) {
			paths = append(paths, fullPath)
		}
	}

	return paths, nil
}
