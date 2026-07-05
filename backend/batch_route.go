package backend

type LBatchResult struct {
	LBatchName              string         `json:"LBatchName"`
	LBatchDirectory         string         `json:"LBatchDirectory"`
	LBatchClip              []LClip        `json:"LBatchClip"`
	LBatchNotice            []string       `json:"LBatchNotice"`
	LBatchCompatibility     LCompatibility `json:"LBatchCompatibility"`
	LBatchPlan              string         `json:"LBatchPlan"`
	LBatchAction            string         `json:"LBatchAction"`
	LProgressStatus         string         `json:"LProgressStatus"`
	LMetricEstimateSize     int64          `json:"LMetricEstimateSize"`
	LMetricEstimateDuration float64        `json:"LMetricEstimateDuration"`
	LMergerResult           *LMergerResult `json:"LMergerResult,omitempty"`
}

type LRouteResult struct {
	LPreferenceInput   []string       `json:"LPreferenceInput"`
	LPreferenceOutput  string         `json:"LPreferenceOutput"`
	LTaskResult        []LBatchResult `json:"LTaskResult"`
	LTaskCancel        bool           `json:"LTaskCancel"`
	LTaskMessage       string         `json:"LTaskMessage"`
	LProgressTotal     int            `json:"LProgressTotal"`
	LProgressProcessed int            `json:"LProgressProcessed"`
}

func LRouteResultCopy(result LRouteResult) LRouteResult {
	copyResult := result
	copyResult.LPreferenceInput = append([]string{}, result.LPreferenceInput...)
	copyResult.LTaskResult = append([]LBatchResult{}, result.LTaskResult...)

	for index := range copyResult.LTaskResult {
		copyResult.LTaskResult[index].LBatchClip = append([]LClip{}, result.LTaskResult[index].LBatchClip...)
		copyResult.LTaskResult[index].LBatchNotice = append([]string{}, result.LTaskResult[index].LBatchNotice...)
		copyResult.LTaskResult[index].LBatchCompatibility.LCompatibilityNotice = append([]string{}, result.LTaskResult[index].LBatchCompatibility.LCompatibilityNotice...)
		copyResult.LTaskResult[index].LBatchCompatibility.LCompatibilityCaution = append([]string{}, result.LTaskResult[index].LBatchCompatibility.LCompatibilityCaution...)
		copyResult.LTaskResult[index].LBatchCompatibility.LCompatibilityWarning = append([]string{}, result.LTaskResult[index].LBatchCompatibility.LCompatibilityWarning...)

		if result.LTaskResult[index].LMergerResult != nil {
			mergeResult := *result.LTaskResult[index].LMergerResult
			copyResult.LTaskResult[index].LMergerResult = &mergeResult
		}
	}

	return copyResult
}
