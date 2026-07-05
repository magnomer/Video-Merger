package backend

import (
	"context"
	"path/filepath"
)

func LClipFileResolve(
	LRuntimeContext context.Context,
	preference LPreference,
	path string,
	batchDirectory string,
	seen map[string]bool,
	marker LMarker,
) (LClip, bool, error) {
	if !LClipCheck(path) {
		return LClip{}, false, nil
	}

	identity := LClipIdentityRead(path)
	if seen[identity] {
		return LClip{}, false, nil
	}
	seen[identity] = true

	mediaFile, ok := LClipParse(path, marker)
	if !ok {
		return LClip{}, false, nil
	}

	if err := LClipSet(LRuntimeContext, preference, &mediaFile); err != nil {
		return LClip{}, false, err
	}

	mediaFile.LBatchDirectory = batchDirectory
	return mediaFile, true, nil
}

func LClipIdentityRead(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		absPath = filepath.Clean(path)
	}

	resolvedPath, err := filepath.EvalSymlinks(absPath)
	if err == nil {
		return filepath.Clean(resolvedPath)
	}

	return filepath.Clean(absPath)
}
