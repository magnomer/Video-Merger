package backend

import "context"

func LClipInputResolve(
	LRuntimeContext context.Context,
	preference LPreference,
	cleanInputPath string,
	isDirectory bool,
	includeSubfolders bool,
	seen map[string]bool,
	marker LMarker,
) ([]LClip, error) {
	if isDirectory {
		return LClipFolderResolve(LRuntimeContext, preference, cleanInputPath, includeSubfolders, seen, marker)
	}

	file, ok, err := LClipFileResolve(LRuntimeContext, preference, cleanInputPath, "", seen, marker)
	if err != nil || !ok {
		return nil, err
	}

	return []LClip{file}, nil
}
