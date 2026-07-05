package backend

import (
	"context"
	"fmt"
	"math"
	"os/exec"
	"strings"
)

type LCompatibilityRateCheck struct {
	LClipName      string
	LClipPath      string
	LStreamIndex   int
	LVideoIndex    int
	LExpectedRate  string
	LActualRate    string
	LExpectedLabel string
	LActualLabel   string
}

func LCompatibilitySmallRateCheck(expected LSignatureStream, actual LSignatureStream) bool {
	if expected.LSignatureCodecType != "video" || actual.LSignatureCodecType != "video" {
		return false
	}

	if expected.LRateAverage == actual.LRateAverage {
		return false
	}

	return LRateCompare(expected.LRateAverage, actual.LRateAverage).LRateSmallState
}

func LCompatibilityRateGuardRun(
	LRuntimeContext context.Context,
	options LPreference,
	group LBatch,
	checks []LCompatibilityRateCheck,
) ([]string, []string, error) {
	if len(checks) == 0 {
		return nil, nil, nil
	}

	cautions := []string{}
	warnings := []string{}

	for _, check := range checks {
		if LRuntimeContext.Err() != nil {
			return nil, nil, LRuntimeContext.Err()
		}

		expected, err := LCompatibilityTimingRead(LRuntimeContext, options, group.LBatchClip[0].LClipPath, check.LVideoIndex)
		if err != nil {
			if LRuntimeContext.Err() != nil {
				return nil, nil, LRuntimeContext.Err()
			}
			cautions = append(cautions, fmt.Sprintf("%s stream %d: small frame-rate difference could not be verified by packet timing: %v", check.LClipName, check.LStreamIndex, err))
			continue
		}

		actual, err := LCompatibilityTimingRead(LRuntimeContext, options, check.LClipPath, check.LVideoIndex)
		if err != nil {
			if LRuntimeContext.Err() != nil {
				return nil, nil, LRuntimeContext.Err()
			}
			cautions = append(cautions, fmt.Sprintf("%s stream %d: small frame-rate difference could not be verified by packet timing: %v", check.LClipName, check.LStreamIndex, err))
			continue
		}

		cautions = append(cautions, LCompatibilityTimingCompare(check, expected, actual)...)
	}

	dryRunWarnings, err := LCompatibilityDryRunCheck(LRuntimeContext, options, group)
	if err != nil {
		if LRuntimeContext.Err() != nil {
			return nil, nil, LRuntimeContext.Err()
		}
		warnings = append(warnings, fmt.Sprintf("Concat dry run failed while verifying a small frame-rate difference: %v", err))
	} else {
		warnings = append(warnings, dryRunWarnings...)
	}

	return cautions, warnings, nil
}

func LCompatibilityTimingRead(LRuntimeContext context.Context, options LPreference, path string, videoIndex int) (LMetricTiming, error) {
	packets, err := LProbePacketRun(LRuntimeContext, options, path, videoIndex)
	if err != nil {
		return LMetricTiming{}, err
	}

	result := LMetricTimingMeasure(packets)
	if result.LMetricDurationCount == 0 {
		return LMetricTiming{}, fmt.Errorf("no packet durations or timestamp deltas were available")
	}

	return result, nil
}

func LCompatibilityTimingCompare(check LCompatibilityRateCheck, expected LMetricTiming, actual LMetricTiming) []string {
	messages := []string{}
	prefix := fmt.Sprintf("%s stream %d", check.LClipName, check.LStreamIndex)

	if !expected.LMetricMonotonic || !actual.LMetricMonotonic {
		messages = append(messages, prefix+": packet timestamps are not monotonic while verifying a small frame-rate difference")
	}

	if expected.LMetricIrregular != actual.LMetricIrregular {
		messages = append(messages, fmt.Sprintf("%s: frame timing pattern differs while frame-rate metadata differs slightly (%s vs %s)", prefix, LCompatibilityTimingKindRead(expected), LCompatibilityTimingKindRead(actual)))
	}

	if math.Abs(expected.LMetricMode-actual.LMetricMode) > 0.001 {
		messages = append(messages, fmt.Sprintf("%s: dominant packet duration differs %.6fs vs %.6fs", prefix, expected.LMetricMode, actual.LMetricMode))
	}

	if math.Abs(expected.LMetricMedian-actual.LMetricMedian) > 0.001 {
		messages = append(messages, fmt.Sprintf("%s: median packet duration differs %.6fs vs %.6fs", prefix, expected.LMetricMedian, actual.LMetricMedian))
	}

	return messages
}

func LCompatibilityTimingKindRead(timing LMetricTiming) string {
	if timing.LMetricIrregular {
		return "variable packet timing"
	}

	return "constant packet timing"
}

func LCompatibilityDryRunCheck(LRuntimeContext context.Context, options LPreference, group LBatch) ([]string, error) {
	listFilePath, err := LConcatCreate(options, group)
	if err != nil {
		return nil, err
	}
	defer LTemporaryOwnedRemove(listFilePath)

	ffmpegPath, err := LCommandFFmpegRead(options)
	if err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(
		LRuntimeContext,
		ffmpegPath,
		"-nostdin",
		"-v", "warning",
		"-f", "concat",
		"-safe", "0",
		"-i", listFilePath,
		"-map", "0",
		"-c", "copy",
		"-f", "null",
		"-",
	)
	LCommandHide(cmd)

	output, err := cmd.CombinedOutput()
	text := strings.TrimSpace(string(output))
	if err != nil {
		return nil, fmt.Errorf("%v %s", err, LCompatibilityTextShorten(text))
	}

	if LCompatibilityDryRunProblemCheck(text) {
		return []string{"Concat dry run reported timestamp or muxing problems: " + LCompatibilityTextShorten(text)}, nil
	}

	return nil, nil
}

func LCompatibilityDryRunProblemCheck(text string) bool {
	lowerText := strings.ToLower(text)
	patterns := []string{
		"non-monotonous dts",
		"non monotonous dts",
		"dts out of order",
		"invalid data found",
		"could not write header",
		"could not find codec parameters",
		"dimensions not set",
		"sample rate mismatch",
		"channel layout mismatch",
	}

	for _, pattern := range patterns {
		if strings.Contains(lowerText, pattern) {
			return true
		}
	}

	return false
}

func LCompatibilityTextShorten(text string) string {
	if text == "" {
		return ""
	}

	text = strings.Join(strings.Fields(text), " ")
	if len(text) <= 600 {
		return text
	}

	return text[:600] + "..."
}

func LCompatibilityVideoIndexRead(signature LSignature, streamIndex int) int {
	videoIndex := -1
	for index := 0; index <= streamIndex && index < len(signature.LProbeStream); index++ {
		if signature.LProbeStream[index].LSignatureCodecType == "video" {
			videoIndex++
		}
	}

	if videoIndex < 0 {
		return 0
	}

	return videoIndex
}
