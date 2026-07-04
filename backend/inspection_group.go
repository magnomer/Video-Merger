package backend

import (
	"context"
	"errors"
)

func LInspectionGroupRun(
	LRuntimeContext context.Context,
	options LPreference,
	result LRouteResult,
	groups []LBatch,
	onReport LReportCallback,
) (LRouteResult, error) {
	for _, group := range groups {
		if LRuntimeContext.Err() != nil {
			result.LTaskCancel = true
			result.LTaskMessage = "Analysis canceled."
			return result, LRuntimeContext.Err()
		}

		groupResult := LInspectionGroupCreate(options, group)
		result.LTaskResult = append(result.LTaskResult, groupResult)
		result.LTaskMessage = "Analyzing " + group.LBatchName
		LReportEmit(onReport, LReportCreate(result, false))

		compatibility, err := LCompatibilityCheck(LRuntimeContext, group)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				result.LTaskCancel = true
				result.LTaskMessage = "Analysis canceled."
				result.LTaskResult[len(result.LTaskResult)-1].LProgressStatus = "Canceled"
				return result, err
			}

			return LRouteResult{}, err
		}

		result.LProgressProcessed += len(group.LBatchClip)
		result.LTaskResult[len(result.LTaskResult)-1].LBatchCompatibility = compatibility
		result.LTaskResult[len(result.LTaskResult)-1].LProgressStatus = "Finished"
		result.LTaskMessage = "Analyzed " + group.LBatchName

		LReportEmit(onReport, LReportCreate(result, false))
	}

	result.LTaskMessage = "Analysis completed."
	result.LProgressProcessed = result.LProgressTotal

	return result, nil
}

func LInspectionGroupCreate(options LPreference, group LBatch) LBatchResult {
	groupOutputFolder := LDestinationBatchRead(options, group)
	plannedOutputPath := LDestinationFind(groupOutputFolder, group.LBatchName+options.LPreferenceSuffix, group.LClipExtension)

	action := "merge"
	if len(group.LBatchClip) == 1 {
		action = "copy"
	}

	return LBatchResult{
		LBatchName:      group.LBatchName,
		LBatchDirectory: group.LBatchDirectory,
		LBatchClip:      group.LBatchClip,
		LBatchNotice:    group.LBatchNotice,
		LBatchCompatibility: LCompatibility{
			LBatchName:          group.LBatchName,
			LCompatibilityState: true,
		},
		LBatchPlan:              plannedOutputPath,
		LBatchAction:            action,
		LProgressStatus:         "Processing",
		LMetricEstimateSize:     LBatchSizeCalculate(group.LBatchClip),
		LMetricEstimateDuration: LBatchDurationCalculate(group.LBatchClip),
	}
}
