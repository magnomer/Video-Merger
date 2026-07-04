package bridge

import (
	"video-merger/backend"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *LProgram) LInspectionStart(options backend.LPreference) (backend.LReport, error) {
	LRuntimeContext, cancel, err := a.lTaskStart()
	if err != nil {
		return backend.LReport{}, err
	}
	defer a.lTaskReset()
	defer cancel()

	return backend.LInspectionPathRun(
		LRuntimeContext,
		options,
		func(report backend.LReport) {
			runtime.EventsEmit(a.LRuntimeContext, "LReportEvent", report)
		},
	)
}

func (a *LProgram) LReportOutputSet(report backend.LReport, options backend.LPreference) (backend.LReport, error) {
	return backend.LReportOutputSet(report, options)
}

func (a *LProgram) LMergerRun(options backend.LPreference) (backend.LReport, error) {
	LRuntimeContext, cancel, err := a.lTaskStart()
	if err != nil {
		return backend.LReport{}, err
	}
	defer a.lTaskReset()
	defer cancel()

	return backend.LMergerPathRun(
		LRuntimeContext,
		options,
		func(report backend.LReport) {
			runtime.EventsEmit(a.LRuntimeContext, "LReportEvent", report)
		},
	)
}
