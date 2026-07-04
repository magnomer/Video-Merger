package backend

type LClip struct {
	LClipPath       string  `json:"LClipPath"`
	LClipName       string  `json:"LClipName"`
	LBatchName      string  `json:"LBatchName"`
	LClipNumber     int     `json:"LClipNumber"`
	LClipExtension  string  `json:"LClipExtension"`
	LBatchDirectory string  `json:"LBatchDirectory"`
	LMetricSize     int64   `json:"LMetricSize"`
	LMetricDuration float64 `json:"LMetricDuration"`
}

type LBatch struct {
	LBatchName      string   `json:"LBatchName"`
	LClipExtension  string   `json:"LClipExtension"`
	LBatchDirectory string   `json:"LBatchDirectory"`
	LBatchClip      []LClip  `json:"LBatchClip"`
	LBatchNotice    []string `json:"LBatchNotice"`
}
