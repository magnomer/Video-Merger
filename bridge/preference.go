package bridge

import "video-merger/backend"

func (a *LProgram) LPreferenceLoad() (backend.LPreference, error) {
	return backend.LPreferenceLoad()
}

func (a *LProgram) LPreferenceSave(preference backend.LPreference) error {
	return backend.LPreferenceSave(preference)
}
