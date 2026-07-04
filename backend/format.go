package backend

import "fmt"

func LByteFormat(bytes int64) string {
	if bytes <= 0 {
		return "Unknown"
	}

	units := []string{"B", "KB", "MB", "GB", "TB"}
	size := float64(bytes)
	unitIndex := 0

	for size >= 1024 && unitIndex < len(units)-1 {
		size /= 1024
		unitIndex++
	}

	if size >= 10 || unitIndex == 0 {
		return fmt.Sprintf("%.0f %s", size, units[unitIndex])
	}

	return fmt.Sprintf("%.1f %s", size, units[unitIndex])
}

func LDurationFormat(seconds float64) string {
	if seconds <= 0 {
		return "Unknown"
	}

	rounded := int(seconds + 0.5)
	hours := rounded / 3600
	minutes := (rounded % 3600) / 60
	remainingSeconds := rounded % 60

	if hours > 0 {
		return fmt.Sprintf("%d:%s:%s", hours, LNumberPad(minutes), LNumberPad(remainingSeconds))
	}

	return fmt.Sprintf("%d:%s", minutes, LNumberPad(remainingSeconds))
}

func LLoudnessFormat(loudness *float64) string {
	if loudness == nil {
		return "Unknown"
	}

	return fmt.Sprintf("%.1f LUFS", *loudness)
}

func LNumberPad(value int) string {
	return fmt.Sprintf("%02d", value)
}
