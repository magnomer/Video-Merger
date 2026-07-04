package backend

import "context"

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

		groupResult := LDestinationPlanResolve(result.LTaskResult[index], options.LPreferenceSuffix)
		group := LMergerGroupCreate(groupResult)

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
	}

	return result
}
