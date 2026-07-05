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
