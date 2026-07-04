package backend

func LReportSectionCreate(group LBatchResult, includeMergeResult bool) []LReportSection {
	sections := []LReportSection{}
	sections = LReportMessageCreate(sections, "Notices", "notice", "Notice", group.LBatchCompatibility.LCompatibilityNotice)
	sections = LReportMessageCreate(sections, "Numbering notices", "notice", "Notice", group.LBatchNotice)
	sections = LReportMessageCreate(sections, "Compatibility cautions", "caution", "Caution", group.LBatchCompatibility.LCompatibilityCaution)
	sections = LReportMessageCreate(sections, "Compatibility warnings", "error", "Warning", group.LBatchCompatibility.LCompatibilityWarning)

	metrics := []LReportMetric{}
	if includeMergeResult {
		metrics = append(metrics,
			LReportMetric{LReportLabel: "Actual size", LReportValue: LReportSizeRead(group, includeMergeResult)},
			LReportMetric{LReportLabel: "Actual duration", LReportValue: LReportDurationRead(group, includeMergeResult)},
			LReportMetric{LReportLabel: "Actual LUFS", LReportValue: LReportLoudnessRead(group, includeMergeResult)},
		)
	} else {
		metrics = append(metrics,
			LReportMetric{LReportLabel: "Estimated size", LReportValue: LReportSizeRead(group, includeMergeResult)},
			LReportMetric{LReportLabel: "Estimated duration", LReportValue: LReportDurationRead(group, includeMergeResult)},
		)
	}

	sections = append(sections, LReportSection{
		LReportTitle:  "Metrics",
		LReportMetric: metrics,
	})

	return sections
}

func LReportMessageCreate(sections []LReportSection, title string, tag string, badge string, items []string) []LReportSection {
	if len(items) == 0 {
		return sections
	}

	return append(sections, LReportSection{
		LReportTitle: title,
		LReportTag:   tag,
		LReportBadge: badge,
		LReportItem:  items,
	})
}
