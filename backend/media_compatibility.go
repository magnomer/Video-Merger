package backend

import (
	"context"
	"fmt"
)

type LCompatibility struct {
	LBatchName            string   `json:"LBatchName"`
	LCompatibilityState   bool     `json:"LCompatibilityState"`
	LCompatibilityNotice  []string `json:"LCompatibilityNotice"`
	LCompatibilityCaution []string `json:"LCompatibilityCaution"`
	LCompatibilityWarning []string `json:"LCompatibilityWarning"`
}

func LCompatibilityCheck(LRuntimeContext context.Context, preference LPreference, group LBatch) (LCompatibility, error) {
	report := LCompatibility{
		LBatchName:          group.LBatchName,
		LCompatibilityState: true,
	}

	if len(group.LBatchClip) < 2 {
		report.LCompatibilityNotice = append(report.LCompatibilityNotice, "Single-file group. It will be copied instead of merged.")
		return report, nil
	}

	firstProbe, err := LProbeRun(LRuntimeContext, preference, group.LBatchClip[0].LClipPath)
	if err != nil {
		if LRuntimeContext.Err() != nil {
			return LCompatibility{}, LRuntimeContext.Err()
		}

		report.LCompatibilityState = false
		report.LCompatibilityWarning = append(report.LCompatibilityWarning, fmt.Sprintf("Could not inspect %s: %v", group.LBatchClip[0].LClipName, err))
		return report, nil
	}

	firstSignature := LSignatureBuild(firstProbe)

	for _, file := range group.LBatchClip[1:] {
		if LRuntimeContext.Err() != nil {
			return LCompatibility{}, LRuntimeContext.Err()
		}

		currentProbe, err := LProbeRun(LRuntimeContext, preference, file.LClipPath)
		if err != nil {
			if LRuntimeContext.Err() != nil {
				return LCompatibility{}, LRuntimeContext.Err()
			}

			report.LCompatibilityState = false
			report.LCompatibilityWarning = append(report.LCompatibilityWarning, fmt.Sprintf("Could not inspect %s: %v", file.LClipName, err))
			continue
		}

		currentSignature := LSignatureBuild(currentProbe)

		if len(currentSignature.LProbeStream) != len(firstSignature.LProbeStream) {
			report.LCompatibilityState = false
			report.LCompatibilityWarning = append(
				report.LCompatibilityWarning,
				fmt.Sprintf("%s has a different number of streams.", file.LClipName),
			)
			continue
		}

		for i := range firstSignature.LProbeStream {
			expected := firstSignature.LProbeStream[i]
			actual := currentSignature.LProbeStream[i]

			result := LSignatureStreamCompare(expected, actual)

			for _, caution := range result.LCompatibilityCaution {
				report.LCompatibilityCaution = append(
					report.LCompatibilityCaution,
					fmt.Sprintf("%s stream %d: %s", file.LClipName, i, caution),
				)
			}

			for _, warning := range result.LCompatibilityWarning {
				report.LCompatibilityState = false
				report.LCompatibilityWarning = append(
					report.LCompatibilityWarning,
					fmt.Sprintf("%s stream %d: %s", file.LClipName, i, warning),
				)
			}
		}
	}

	return report, nil
}
