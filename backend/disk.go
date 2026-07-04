package backend

import (
	"context"
	"io"
	"os"
)

func LDiskCheck(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func LDiskCancelCopy(LRuntimeContext context.Context, sourcePath string, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	buffer := make([]byte, 1024*1024)

	for {
		if LRuntimeContext.Err() != nil {
			destinationFile.Close()
			os.Remove(destinationPath)
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

	return nil
}
