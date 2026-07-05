package backend

import "context"

func LInputResolve(inputPaths []string) ([]string, error) {
	mediaFiles, err := LClipResolve(context.Background(), LPreference{}, inputPaths, false, LMarkerDefaultCreate())
	if err != nil {
		return nil, err
	}

	var paths []string

	for _, file := range mediaFiles {
		paths = append(paths, file.LClipPath)
	}

	return paths, nil
}
