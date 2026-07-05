package backend

import (
	"context"
	"fmt"
)

func LMergerTaskRun(
	LRuntimeContext context.Context,
	options LPreference,
	result LRouteResult,
	onReport LReportCallback,
) LReport {
	result = LMergerTaskPrepare(result)
	totalFiles := result.LProgressTotal
	processedFiles := 0

	LReportEmit(onReport, LReportCreate(result, true))

	for index := range result.LTaskResult {
		if LRuntimeContext.Err() != nil {
			finalReport := LTaskCancelSet(result, processedFiles, index, "Processing canceled.")
			LReportEmit(onReport, finalReport)
			return finalReport
		}

		group := LMergerGroupCreate(result.LTaskResult[index])
		groupResult := LMergerTaskPlanSet(options, result.LTaskResult[index], group)

		if groupResult.LMergerResult != nil {
			result.LTaskResult[index] = groupResult
			result.LTaskMessage = groupResult.LMergerResult.LTaskMessage
			LReportEmit(onReport, LReportCreate(result, true))
			continue
		}

		groupResult.LProgressStatus = "Processing"
		result.LTaskResult[index] = groupResult
		result.LTaskMessage = "Processing " + group.LBatchName
		LReportEmit(onReport, LReportCreate(result, true))

		mergeResult := LMergerGroupRun(LRuntimeContext, options, groupResult, group)
		groupResult.LMergerResult = &mergeResult

		if LRuntimeContext.Err() != nil {
			groupResult.LProgressStatus = "Canceled"
			result.LTaskResult[index] = groupResult
			finalReport := LTaskCancelSet(result, processedFiles, index+1, mergeResult.LTaskMessage)
			LReportEmit(onReport, finalReport)
			return finalReport
		}

		groupResult = LMergerGroupStatusSet(groupResult, mergeResult)
		result.LTaskResult[index] = groupResult

		processedFiles += len(group.LBatchClip)
		result.LProgressProcessed = processedFiles
		result.LTaskMessage = mergeResult.LTaskMessage

		LReportEmit(onReport, LReportCreate(result, true))
	}

	result.LTaskMessage = "Processing completed."
	result.LProgressProcessed = totalFiles

	finalReport := LReportCreate(result, true)
	LReportEmit(onReport, finalReport)

	return finalReport
}

func LMergerTaskPrepare(result LRouteResult) LRouteResult {
	result.LProgressProcessed = 0
	result.LTaskCancel = false
	result.LTaskMessage = "Starting merge."

	for index := range result.LTaskResult {
		result.LTaskResult[index].LProgressStatus = ""
		result.LTaskResult[index].LMergerResult = nil
	}

	return result
}

func LMergerTaskPlanSet(options LPreference, groupResult LBatchResult, group LBatch) LBatchResult {
	outputPath, err := LDestinationPlanCreate(options, group)
	if err == nil {
		groupResult.LBatchPlan = outputPath
		return groupResult
	}

	mergeResult := LMergerResult{
		LTaskSuccess: false,
		LTaskMessage: fmt.Sprintf("Could not plan output: %v", err),
	}
	groupResult.LProgressStatus = "Failed"
	groupResult.LMergerResult = &mergeResult
	return groupResult
}
