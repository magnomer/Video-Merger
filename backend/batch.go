package backend

func LBatchSizeCalculate(files []LClip) int64 {
	var total int64

	for _, file := range files {
		total += file.LMetricSize
	}

	return total
}

func LBatchDurationCalculate(files []LClip) float64 {
	var total float64

	for _, file := range files {
		total += file.LMetricDuration
	}

	return total
}

func LBatchExtensionRead(files []LClip) string {
	if len(files) == 0 {
		return ".mp4"
	}

	return files[0].LClipExtension
}
