package backend

import "context"

func LClipFileResolve(
	LRuntimeContext context.Context,
	path string,
	batchDirectory string,
	seen map[string]bool,
	marker LMarker,
) (LClip, bool, error) {
	if !LClipCheck(path) {
		return LClip{}, false, nil
	}

	if seen[path] {
		return LClip{}, false, nil
	}
	seen[path] = true

	mediaFile, ok := LClipParse(path, marker)
	if !ok {
		return LClip{}, false, nil
	}

	if err := LClipSet(LRuntimeContext, &mediaFile); err != nil {
		return LClip{}, false, err
	}

	mediaFile.LBatchDirectory = batchDirectory
	return mediaFile, true, nil
}
