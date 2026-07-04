package backend

import "path/filepath"

func LDirectoryPathRead(path string) string {
	folder := filepath.Dir(path)
	if folder == "" {
		return "."
	}

	return folder
}
