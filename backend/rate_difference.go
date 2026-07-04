package backend

import "fmt"

func LRateDifferenceRead(expected string, expectedFPS float64, actual string, actualFPS float64, difference float64) LRateComparison {
	if difference <= 0.0001 {
		return LRateComparison{}
	}

	if difference <= 0.05 {
		return LRateComparison{
			LTaskMessage: fmt.Sprintf(
				"frame rate metadata differs slightly: %s %.4f fps vs %s %.4f fps. This is probably timestamp/metadata noise from cutting.",
				expected,
				expectedFPS,
				actual,
				actualFPS,
			),
			LRateWarningState: false,
		}
	}

	if difference <= 1.0 {
		return LRateComparison{
			LTaskMessage: fmt.Sprintf(
				"frame rate metadata differs meaningfully: %s %.4f fps vs %s %.4f fps. The files may still merge, but review the output.",
				expected,
				expectedFPS,
				actual,
				actualFPS,
			),
			LRateWarningState: false,
		}
	}

	return LRateComparison{
		LTaskMessage: fmt.Sprintf(
			"frame rate metadata differs significantly: %s %.4f fps vs %s %.4f fps. This may indicate a real mismatch.",
			expected,
			expectedFPS,
			actual,
			actualFPS,
		),
		LRateWarningState: true,
	}
}
