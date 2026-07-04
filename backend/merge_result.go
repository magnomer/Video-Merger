package backend

type LMergerResult struct {
	LDestinationPath string   `json:"LDestinationPath"`
	LTaskSuccess     bool     `json:"LTaskSuccess"`
	LTaskMessage     string   `json:"LTaskMessage"`
	LMetricSize      int64    `json:"LMetricSize"`
	LMetricDuration  float64  `json:"LMetricDuration"`
	LMetricLoudness  *float64 `json:"LMetricLoudness,omitempty"`
}
