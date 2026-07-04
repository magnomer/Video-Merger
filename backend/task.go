package backend

import (
	"context"
	"errors"
)

func LMergerPathRun(
	LRuntimeContext context.Context,
	options LPreference,
	onReport LReportCallback,
) (LReport, error) {
	result, err := LInspectionCoreRun(
		LRuntimeContext,
		options,
		func(report LReport) {
			report.LReportKind = "merge"
			report.LTaskMessage = LMergerAnalysisMessageRead(report.LTaskMessage)
			LReportEmit(onReport, report)
		},
	)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			finalReport := LMergerCancelReportCreate(result, options, "Merge canceled during analysis.")
			LReportEmit(onReport, finalReport)
			return finalReport, nil
		}

		return LReport{}, err
	}

	return LMergerTaskRun(LRuntimeContext, options, result, onReport), nil
}

func LMergerAnalysisMessageRead(message string) string {
	if message == "" {
		return "Analyzing files before merge."
	}

	return message
}
