package bridge

import "video-merger/backend"

func (a *LProgram) LAssetPreviewStop() string {
	backend.LAssetPreviewStop()
	return "Preview loading stopped."
}
