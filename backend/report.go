package backend

func LReportCreate(result LRouteResult, includeMergeResult bool) LReport {
	report := LReport{
		LReportKind:        "analysis",
		LTaskMessage:       result.LTaskMessage,
		LTaskCancel:        result.LTaskCancel,
		LProgressTotal:     result.LProgressTotal,
		LProgressProcessed: result.LProgressProcessed,
		LProgressPercent:   LPercentCalculate(result.LProgressProcessed, result.LProgressTotal),
		LReportGroup:       []LReportGroup{},
	}

	if includeMergeResult {
		report.LReportKind = "merge"
	}

	for _, group := range result.LTaskResult {
		report.LReportGroup = append(report.LReportGroup, LReportGroupCreate(group, includeMergeResult))
	}

	return report
}

func LReportGroupCreate(group LBatchResult, includeMergeResult bool) LReportGroup {
	compatibility := LReportCompatibilityRead(group)
	task := LReportTaskRead(group, includeMergeResult)
	outputTitle := "Planned output"
	outputText := group.LBatchPlan

	if includeMergeResult {
		outputTitle = "Output"
		if group.LMergerResult == nil {
			outputText = "Pending."
		} else if group.LMergerResult.LTaskSuccess {
			outputText = group.LMergerResult.LDestinationPath
		} else {
			outputText = group.LMergerResult.LTaskMessage
		}
	}

	return LReportGroup{
		LReportKey:              LReportKeyRead(group),
		LReportName:             group.LBatchName,
		LReportDirectory:        group.LBatchDirectory,
		LReportSize:             LReportSizeRead(group, includeMergeResult),
		LReportDuration:         LReportDurationRead(group, includeMergeResult),
		LReportLoudness:         LReportLoudnessRead(group, includeMergeResult),
		LReportCompatibility:    compatibility.LReportLabel,
		LReportCompatibilityTag: compatibility.LReportTag,
		LReportTask:             task.LReportLabel,
		LReportTaskTag:          task.LReportTag,
		LReportOutputTitle:      outputTitle,
		LReportOutputText:       outputText,
		LReportFile:             LReportFileCreate(group.LBatchClip),
		LReportSection:          LReportSectionCreate(group, includeMergeResult),
	}
}
