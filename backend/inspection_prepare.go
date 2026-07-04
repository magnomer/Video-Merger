package backend

import (
	"errors"
	"strings"
)

func LInspectionPrepare(options LPreference) (LRouteResult, LPreference, error) {
	options.LPreferenceSuffix = strings.TrimSpace(options.LPreferenceSuffix)

	if err := LSuffixCheck(options.LPreferenceSuffix); err != nil {
		return LRouteResult{}, options, err
	}

	if len(options.LPreferenceInput) == 0 {
		return LRouteResult{}, options, errors.New("at least one input path is required")
	}

	if !options.LPreferenceMirror && strings.TrimSpace(options.LPreferenceOutput) == "" {
		return LRouteResult{}, options, errors.New("output folder is required unless Same as input is checked")
	}

	return LRouteResult{
		LPreferenceInput:   options.LPreferenceInput,
		LPreferenceOutput:  options.LPreferenceOutput,
		LTaskResult:        []LBatchResult{},
		LProgressTotal:     0,
		LProgressProcessed: 0,
		LTaskMessage:       "Preparing analysis.",
	}, options, nil
}
