package backend

import (
	"strconv"
	"strings"
)

func LRateParse(value string) (float64, bool) {
	value = strings.TrimSpace(value)

	if value == "" || value == "0/0" {
		return 0, false
	}

	if strings.Contains(value, "/") {
		return LRateRatioRead(value)
	}

	fps, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false
	}

	return fps, true
}

func LRateRatioRead(value string) (float64, bool) {
	parts := strings.Split(value, "/")
	if len(parts) != 2 {
		return 0, false
	}

	numerator, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, false
	}

	denominator, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, false
	}

	if denominator == 0 {
		return 0, false
	}

	return numerator / denominator, true
}
