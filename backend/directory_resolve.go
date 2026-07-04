package backend

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func LDirectorySourceResolve(sourceText string) (string, error) {
	for _, sourceLine := range strings.Split(sourceText, "\n") {
		sourcePath := strings.TrimSpace(sourceLine)
		if sourcePath == "" {
			continue
		}

		return LDirectoryExistingResolve(sourcePath)
	}

	return "", errors.New("source path is empty")
}

func LDirectoryOutputResolve(outputText string, sourceText string, sameAsInput bool) (string, error) {
	if sameAsInput {
		return LDirectorySourceResolve(sourceText)
	}

	outputPath := strings.TrimSpace(outputText)
	if outputPath == "" {
		return "", errors.New("output path is empty")
	}

	return LDirectoryExistingResolve(outputPath)
}

func LDirectoryExistingResolve(path string) (string, error) {
	cleanPath := filepath.Clean(path)
	pathInfo, err := os.Stat(cleanPath)
	if err != nil {
		return "", err
	}

	if pathInfo.IsDir() {
		return cleanPath, nil
	}

	return filepath.Dir(cleanPath), nil
}
