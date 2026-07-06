package bridge

import "video-merger/backend"

func (a *LProgram) LAssetPreviewStop(session string) string {
	backend.LAssetPreviewStop(session)
	return "Preview loading stopped."
}
