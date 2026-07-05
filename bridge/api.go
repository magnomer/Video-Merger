package bridge

import (
	"context"
	"errors"
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

	result, err := backend.LInspectionCoreRun(
		LRuntimeContext,
		options,
		func(report backend.LReport) {
			runtime.EventsEmit(a.LRuntimeContext, "LReportEvent", report)
		},
	)
	if err != nil {
		a.lInspectionClear()
		if errors.Is(err, context.Canceled) {
			finalReport := backend.LInspectionCancelReportCreate(result, options, "Analysis canceled.")
			runtime.EventsEmit(a.LRuntimeContext, "LReportEvent", finalReport)
			return finalReport, nil
		}

		return backend.LReport{}, err
	}

	a.lInspectionSet(options, result)
	return backend.LReportCreate(result, false), nil
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

	if result, ok := a.lInspectionRead(options); ok {
		return backend.LMergerTaskRun(
			LRuntimeContext,
			options,
			result,
			func(report backend.LReport) {
				runtime.EventsEmit(a.LRuntimeContext, "LReportEvent", report)
			},
		), nil
	}

	return backend.LMergerPathRun(
		LRuntimeContext,
		options,
		func(report backend.LReport) {
			runtime.EventsEmit(a.LRuntimeContext, "LReportEvent", report)
		},
	)
}
