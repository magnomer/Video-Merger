package backend

func LReportCompatibilityRead(group LBatchResult) LReportBadge {
	compatibility := group.LBatchCompatibility
	hasWarning := len(compatibility.LCompatibilityWarning) > 0
	hasCaution := len(compatibility.LCompatibilityCaution) > 0
	hasNotice := len(compatibility.LCompatibilityNotice) > 0 || len(group.LBatchNotice) > 0

	if !compatibility.LCompatibilityState || hasWarning {
		return LReportBadge{LReportLabel: "Incompatible", LReportTag: "error"}
	}

	if hasCaution {
		return LReportBadge{LReportLabel: "Caution", LReportTag: "caution"}
	}

	if hasNotice {
		return LReportBadge{LReportLabel: "Notice", LReportTag: "notice"}
	}

	return LReportBadge{LReportLabel: "Compatible", LReportTag: "ok"}
}

func LReportTaskRead(group LBatchResult, includeMergeResult bool) LReportBadge {
	if group.LProgressStatus == "Processing" {
		return LReportBadge{LReportLabel: "Processing", LReportTag: "notice"}
	}

	if group.LProgressStatus == "Canceled" {
		return LReportBadge{LReportLabel: "Canceled", LReportTag: "error"}
	}

	if !includeMergeResult || group.LMergerResult == nil {
		return LReportBadge{LReportLabel: "Not merged", LReportTag: "neutral"}
	}

	if group.LMergerResult.LTaskSuccess {
		if group.LBatchAction == "copy" {
			return LReportBadge{LReportLabel: "Copied", LReportTag: "ok"}
		}

		return LReportBadge{LReportLabel: "Merged", LReportTag: "ok"}
	}

	return LReportBadge{LReportLabel: "Failed", LReportTag: "error"}
}
