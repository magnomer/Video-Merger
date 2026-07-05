package bridge

import (
	"reflect"
	"strings"
	"video-merger/backend"
)

func (a *LProgram) lInspectionSet(options backend.LPreference, result backend.LRouteResult) {
	a.LInspectionLock.Lock()
	defer a.LInspectionLock.Unlock()

	a.LInspectionPreference = lInspectionPreferenceCreate(options)
	a.LInspectionResult = backend.LRouteResultCopy(result)
	a.LInspectionReady = true
}

func (a *LProgram) lInspectionClear() {
	a.LInspectionLock.Lock()
	defer a.LInspectionLock.Unlock()

	a.LInspectionPreference = backend.LPreference{}
	a.LInspectionResult = backend.LRouteResult{}
	a.LInspectionReady = false
}

func (a *LProgram) lInspectionRead(options backend.LPreference) (backend.LRouteResult, bool) {
	a.LInspectionLock.Lock()
	defer a.LInspectionLock.Unlock()

	if !a.LInspectionReady || !lInspectionPreferenceCompare(a.LInspectionPreference, lInspectionPreferenceCreate(options)) {
		return backend.LRouteResult{}, false
	}

	return backend.LRouteResultCopy(a.LInspectionResult), true
}

func lInspectionPreferenceCreate(options backend.LPreference) backend.LPreference {
	return backend.LPreference{
		LPreferenceInput:      lInspectionInputCreate(options.LPreferenceInput),
		LPreferenceTree:       options.LPreferenceTree,
		LPreferenceMarker:     strings.TrimSpace(options.LPreferenceMarker),
		LPreferencePattern:    strings.TrimSpace(options.LPreferencePattern),
		LPreferenceCustom:     options.LPreferenceCustom,
		LPreferenceUnnumbered: options.LPreferenceUnnumbered,
		LPreferenceFFmpeg:     strings.TrimSpace(options.LPreferenceFFmpeg),
	}
}

func lInspectionInputCreate(input []string) []string {
	cleanInput := []string{}

	for _, value := range input {
		cleanValue := strings.TrimSpace(value)
		if cleanValue != "" {
			cleanInput = append(cleanInput, cleanValue)
		}
	}

	return cleanInput
}

func lInspectionPreferenceCompare(first backend.LPreference, second backend.LPreference) bool {
	return first.LPreferenceTree == second.LPreferenceTree &&
		first.LPreferenceMarker == second.LPreferenceMarker &&
		first.LPreferencePattern == second.LPreferencePattern &&
		first.LPreferenceCustom == second.LPreferenceCustom &&
		first.LPreferenceUnnumbered == second.LPreferenceUnnumbered &&
		first.LPreferenceFFmpeg == second.LPreferenceFFmpeg &&
		reflect.DeepEqual(first.LPreferenceInput, second.LPreferenceInput)
}
