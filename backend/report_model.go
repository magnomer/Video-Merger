package backend

type LReport struct {
	LReportKind        string         `json:"LReportKind"`
	LTaskMessage       string         `json:"LTaskMessage"`
	LTaskCancel        bool           `json:"LTaskCancel"`
	LProgressTotal     int            `json:"LProgressTotal"`
	LProgressProcessed int            `json:"LProgressProcessed"`
	LProgressPercent   int            `json:"LProgressPercent"`
	LReportGroup       []LReportGroup `json:"LReportGroup"`
}

type LReportGroup struct {
	LReportKey              string           `json:"LReportKey"`
	LReportName             string           `json:"LReportName"`
	LReportDirectory        string           `json:"LReportDirectory"`
	LReportSize             string           `json:"LReportSize"`
	LReportDuration         string           `json:"LReportDuration"`
	LReportLoudness         string           `json:"LReportLoudness"`
	LReportCompatibility    string           `json:"LReportCompatibility"`
	LReportCompatibilityTag string           `json:"LReportCompatibilityTag"`
	LReportTask             string           `json:"LReportTask"`
	LReportTaskTag          string           `json:"LReportTaskTag"`
	LReportOutputTitle      string           `json:"LReportOutputTitle"`
	LReportOutputText       string           `json:"LReportOutputText"`
	LReportFile             []LReportFile    `json:"LReportFile"`
	LReportSection          []LReportSection `json:"LReportSection"`
}

type LReportFile struct {
	LReportNumber         int     `json:"LReportNumber"`
	LReportName           string  `json:"LReportName"`
	LReportPath           string  `json:"LReportPath"`
	LReportAsset          string  `json:"LReportAsset"`
	LReportDurationSecond float64 `json:"LReportDurationSecond"`
}

type LReportSection struct {
	LReportTitle  string          `json:"LReportTitle"`
	LReportTag    string          `json:"LReportTag"`
	LReportBadge  string          `json:"LReportBadge"`
	LReportItem   []string        `json:"LReportItem"`
	LReportMetric []LReportMetric `json:"LReportMetric"`
}

type LReportMetric struct {
	LReportLabel string `json:"LReportLabel"`
	LReportValue string `json:"LReportValue"`
}

type LReportBadge struct {
	LReportLabel string
	LReportTag   string
}
