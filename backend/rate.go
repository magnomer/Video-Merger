package backend

import (
	"fmt"
	"math"
)

type LRateComparison struct {
	LTaskMessage      string
	LRateWarningState bool
	LRateSmallState   bool
}

func LRateCompare(expected string, actual string) LRateComparison {
	expectedFPS, expectedOK := LRateParse(expected)
	actualFPS, actualOK := LRateParse(actual)

	if !expectedOK || !actualOK {
		return LRateComparison{
			LTaskMessage:      fmt.Sprintf("frame rate metadata differs: %s vs %s", expected, actual),
			LRateWarningState: false,
		}
	}

	difference := math.Abs(expectedFPS - actualFPS)

	return LRateDifferenceRead(expected, expectedFPS, actual, actualFPS, difference)
}
