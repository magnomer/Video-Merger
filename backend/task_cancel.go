package backend

func LMergerCancelReportCreate(result LRouteResult, options LPreference, message string) LReport {
	if result.LPreferenceInput == nil {
		result.LPreferenceInput = options.LPreferenceInput
	}

	if result.LPreferenceOutput == "" {
		result.LPreferenceOutput = options.LPreferenceOutput
	}

	result.LTaskCancel = true
	result.LTaskMessage = message

	for index := range result.LTaskResult {
		if result.LTaskResult[index].LProgressStatus == "Processing" || result.LTaskResult[index].LProgressStatus == "" {
			result.LTaskResult[index].LProgressStatus = "Canceled"
		}
	}

	return LReportCreate(result, true)
}

func LTaskCancelSet(result LRouteResult, processedFiles int, startIndex int, message string) LReport {
	result.LTaskCancel = true
	result.LTaskMessage = message
	result.LProgressProcessed = processedFiles

	for cancelIndex := startIndex; cancelIndex < len(result.LTaskResult); cancelIndex++ {
		result.LTaskResult[cancelIndex].LProgressStatus = "Canceled"
	}

	return LReportCreate(result, true)
}
