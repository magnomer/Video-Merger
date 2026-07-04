package backend

import (
	"context"
	"errors"
)

func LInspectionResolveErrorRead(result LRouteResult, err error) (LRouteResult, error) {
	if errors.Is(err, context.Canceled) {
		result.LTaskCancel = true
		result.LTaskMessage = "Analysis canceled."
		return result, err
	}

	return LRouteResult{}, err
}

func LInspectionCancelReportCreate(result LRouteResult, options LPreference, message string) LReport {
	if result.LPreferenceInput == nil {
		result.LPreferenceInput = options.LPreferenceInput
	}

	if result.LPreferenceOutput == "" {
		result.LPreferenceOutput = options.LPreferenceOutput
	}

	result.LTaskCancel = true
	result.LTaskMessage = message

	return LReportCreate(result, false)
}
