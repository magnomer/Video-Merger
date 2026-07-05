package bridge

import "video-merger/backend"

func (a *LProgram) LTemporaryClean(preference backend.LPreference) (backend.LTemporaryResult, error) {
	return backend.LTemporaryClean(preference)
}
