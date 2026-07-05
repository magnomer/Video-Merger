package backend

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	lAssetLock     sync.Mutex
	lAssetSecret   = LAssetSecretCreate()
	lAssetPathByID = map[string]string{}
	lAssetIDByPath = map[string]string{}
)

func LAssetSecretCreate() []byte {
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		panic(err)
	}

	return secret
}

func LAssetIDRead(path string) string {
	cleanPath, info, err := LAssetFileResolve(path)
	if err != nil {
		return ""
	}

	seed := fmt.Sprintf("%s\x00%d\x00%d", cleanPath, info.Size(), info.ModTime().UnixNano())
	signature := hmac.New(sha256.New, lAssetSecret)
	_, _ = signature.Write([]byte(seed))
	id := hex.EncodeToString(signature.Sum(nil))

	lAssetLock.Lock()
	defer lAssetLock.Unlock()

	lAssetIDByPath[cleanPath] = id
	lAssetPathByID[id] = cleanPath

	return id
}

func LAssetFileOpen(ctx context.Context, id string, compatibility bool) (*os.File, os.FileInfo, error) {
	path, info, ok := LAssetPathRead(id)
	if !ok {
		return nil, nil, os.ErrNotExist
	}

	servedPath := path
	if compatibility {
		previewPath, err := LAssetPreviewResolve(ctx, path, info)
		if err == nil {
			servedPath = previewPath
		}
	}

	return LAssetPathOpen(servedPath)
}

func LAssetPathOpen(path string) (*os.File, os.FileInfo, error) {
	cleanPath, info, err := LAssetFileResolve(path)
	if err != nil {
		return nil, nil, err
	}

	file, err := os.Open(cleanPath)
	if err != nil {
		return nil, nil, err
	}

	currentInfo, err := file.Stat()
	if err != nil {
		_ = file.Close()
		return nil, nil, err
	}

	if currentInfo.IsDir() || !currentInfo.Mode().IsRegular() || currentInfo.Size() != info.Size() || !currentInfo.ModTime().Equal(info.ModTime()) {
		_ = file.Close()
		return nil, nil, os.ErrInvalid
	}

	return file, currentInfo, nil
}

func LAssetPathRead(id string) (string, os.FileInfo, bool) {
	if !LAssetIDCheck(id) {
		return "", nil, false
	}

	lAssetLock.Lock()
	path := lAssetPathByID[id]
	lAssetLock.Unlock()

	if path == "" {
		return "", nil, false
	}

	cleanPath, info, err := LAssetFileResolve(path)
	if err != nil || cleanPath != path {
		return "", nil, false
	}

	return cleanPath, info, true
}

func LAssetIDCheck(id string) bool {
	if len(id) != sha256.Size*2 {
		return false
	}

	for _, letter := range id {
		if (letter >= '0' && letter <= '9') || (letter >= 'a' && letter <= 'f') {
			continue
		}

		return false
	}

	return true
}

func LAssetFileResolve(path string) (string, os.FileInfo, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return "", nil, os.ErrInvalid
	}

	absolutePath, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return "", nil, err
	}

	info, err := os.Stat(absolutePath)
	if err != nil {
		return "", nil, err
	}

	if info.IsDir() || !info.Mode().IsRegular() {
		return "", nil, os.ErrInvalid
	}

	return absolutePath, info, nil
}
