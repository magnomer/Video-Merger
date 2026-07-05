package backend

import (
	"context"
	"errors"
)

func LInspectionPathRun(LRuntimeContext context.Context, options LPreference, onReport LReportCallback) (LReport, error) {
	result, err := LInspectionCoreRun(LRuntimeContext, options, onReport)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			finalReport := LInspectionCancelReportCreate(result, options, "Analysis canceled.")
			LReportEmit(onReport, finalReport)
			return finalReport, nil
		}

		return LReport{}, err
	}

	return LReportCreate(result, false), nil
}

func LInspectionCoreRun(LRuntimeContext context.Context, options LPreference, onReport LReportCallback) (LRouteResult, error) {
	result, options, err := LInspectionPrepare(options)
	if err != nil {
		return LRouteResult{}, err
	}

	marker, err := LMarkerResolve(options)
	if err != nil {
		return LRouteResult{}, err
	}

	mediaFiles, err := LClipResolve(LRuntimeContext, options, options.LPreferenceInput, options.LPreferenceTree, marker)
	if err != nil {
		return LInspectionResolveErrorRead(result, err)
	}

	groups := LBatchClipCreate(mediaFiles)
	result.LProgressTotal = len(mediaFiles)
	result.LTaskMessage = "Analysis started."
	LReportEmit(onReport, LReportCreate(result, false))

	return LInspectionGroupRun(LRuntimeContext, options, result, groups, onReport)
}
