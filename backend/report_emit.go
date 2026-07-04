package backend

type LReportCallback func(report LReport)

func LReportEmit(onReport LReportCallback, report LReport) {
	if onReport != nil {
		onReport(report)
	}
}
