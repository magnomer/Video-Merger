package backend

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func LDiskCheck(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func LFileRegularCheck(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return false
	}
	return info.Mode().IsRegular() && info.Size() > 0
}

func LFileRemoveIfExists(path string) error {
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func LDiskOutputTemporaryRead(destinationPath string) (string, error) {
	folder := filepath.Dir(destinationPath)
	base := filepath.Base(destinationPath)
	extension := filepath.Ext(base)
	if extension == "" {
		extension = ".tmp"
	}
	stem := strings.TrimSuffix(base, filepath.Ext(base))
	if stem == "" {
		stem = "output"
	}
	if err := os.MkdirAll(folder, 0o755); err != nil {
		return "", err
	}

	file, err := os.CreateTemp(folder, "."+stem+".building-*"+extension)
	if err != nil {
		return "", err
	}
	path := file.Name()
	if err := file.Close(); err != nil {
		os.Remove(path)
		return "", err
	}
	if err := os.Remove(path); err != nil {
		return "", err
	}

	return path, nil
}

func LDiskCancelCopy(LRuntimeContext context.Context, sourcePath string, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	if err := os.MkdirAll(filepath.Dir(destinationPath), 0o755); err != nil {
		return err
	}

	destinationFile, err := os.OpenFile(destinationPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return err
	}
	completed := false
	defer func() {
		destinationFile.Close()
		if !completed {
			os.Remove(destinationPath)
		}
	}()

	buffer := make([]byte, 1024*1024)

	for {
		if LRuntimeContext.Err() != nil {
			return LRuntimeContext.Err()
		}

		n, readErr := sourceFile.Read(buffer)
		if n > 0 {
			_, writeErr := destinationFile.Write(buffer[:n])
			if writeErr != nil {
				return writeErr
			}
		}

		if readErr == io.EOF {
			break
		}

		if readErr != nil {
			return readErr
		}
	}

	if err := destinationFile.Close(); err != nil {
		return err
	}
	completed = true
	return nil
}

func LDiskPublishMove(sourcePath string, destinationPath string) error {
	if !LFileRegularCheck(sourcePath) {
		return fmt.Errorf("temporary output is not a regular file")
	}
	if err := os.MkdirAll(filepath.Dir(destinationPath), 0o755); err != nil {
		return err
	}

	if err := os.Link(sourcePath, destinationPath); err != nil {
		return err
	}
	return os.Remove(sourcePath)
}
