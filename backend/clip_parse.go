package backend

import (
	"path/filepath"
	"strconv"
	"strings"
)

func LClipParse(path string, marker LMarker) (LClip, bool) {
	name := filepath.Base(path)
	extension := filepath.Ext(name)
	stem := strings.TrimSuffix(name, extension)

	matches := marker.LMarkerPattern.FindStringSubmatch(stem)
	if matches != nil {
		partNumber, err := strconv.Atoi(matches[2])
		if err != nil {
			return LClip{}, false
		}

		return LClip{
			LClipPath:      path,
			LClipName:      name,
			LBatchName:     strings.TrimSpace(matches[1]),
			LClipNumber:    partNumber,
			LClipExtension: strings.ToLower(extension),
		}, true
	}

	if marker.LMarkerUnnumbered {
		return LClip{
			LClipPath:      path,
			LClipName:      name,
			LBatchName:     stem,
			LClipNumber:    0,
			LClipExtension: strings.ToLower(extension),
		}, true
	}

	return LClip{}, false
}
