package backend

import "fmt"

func LReportKeyRead(group LBatchResult) string {
	return fmt.Sprintf("%s\x00%s", group.LBatchDirectory, group.LBatchName)
}

func LReportFileCreate(files []LClip) []LReportFile {
	items := []LReportFile{}

	for _, file := range files {
		items = append(items, LReportFile{
			LReportNumber:         file.LClipNumber,
			LReportName:           file.LClipName,
			LReportPath:           file.LClipPath,
			LReportDurationSecond: file.LMetricDuration,
		})
	}

	return items
}

func LReportSizeRead(group LBatchResult, includeMergeResult bool) string {
	if includeMergeResult {
		if group.LMergerResult == nil || group.LMergerResult.LMetricSize <= 0 {
			return "Pending"
		}

		return LByteFormat(group.LMergerResult.LMetricSize)
	}

	return LByteFormat(group.LMetricEstimateSize)
}

func LReportDurationRead(group LBatchResult, includeMergeResult bool) string {
	if includeMergeResult {
		if group.LMergerResult == nil || group.LMergerResult.LMetricDuration <= 0 {
			return "Pending"
		}

		return LDurationFormat(group.LMergerResult.LMetricDuration)
	}

	return LDurationFormat(group.LMetricEstimateDuration)
}

func LReportLoudnessRead(group LBatchResult, includeMergeResult bool) string {
	if !includeMergeResult || group.LMergerResult == nil {
		return "-"
	}

	return LLoudnessFormat(group.LMergerResult.LMetricLoudness)
}
