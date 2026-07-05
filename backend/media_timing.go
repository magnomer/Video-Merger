package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os/exec"
	"sort"
	"strconv"
)

type LProbePacketResult struct {
	LProbePacket []LProbePacket `json:"packets"`
}

type LProbePacket struct {
	LProbePtsTime      string `json:"pts_time"`
	LProbeDtsTime      string `json:"dts_time"`
	LProbeDurationTime string `json:"duration_time"`
}

type LMetricTiming struct {
	LMetricPacketCount   int
	LMetricDurationCount int
	LMetricMedian        float64
	LMetricMode          float64
	LMetricModeShare     float64
	LMetricDistinctCount int
	LMetricIrregular     bool
	LMetricMonotonic     bool
}

func LProbePacketRun(LRuntimeContext context.Context, preference LPreference, path string, videoIndex int) ([]LProbePacket, error) {
	ffprobePath, err := LCommandFFprobeRead(preference)
	if err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(
		LRuntimeContext,
		ffprobePath,
		"-v", "error",
		"-select_streams", fmt.Sprintf("v:%d", videoIndex),
		"-show_packets",
		"-show_entries", "packet=pts_time,dts_time,duration_time",
		"-of", "json",
		path,
	)
	LCommandHide(cmd)

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var result LProbePacketResult
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, err
	}

	return result.LProbePacket, nil
}

func LMetricTimingMeasure(packets []LProbePacket) LMetricTiming {
	result := LMetricTiming{
		LMetricPacketCount: len(packets),
		LMetricMonotonic:   true,
	}

	durations := []float64{}
	stamps := []float64{}

	for _, packet := range packets {
		if duration, ok := LMetricFloatParse(packet.LProbeDurationTime); ok && duration > 0 {
			durations = append(durations, duration)
		}

		if stamp, ok := LMetricTimestampRead(packet); ok {
			stamps = append(stamps, stamp)
		}
	}

	for i := 1; i < len(stamps); i++ {
		delta := stamps[i] - stamps[i-1]
		if delta < -0.000001 {
			result.LMetricMonotonic = false
		}
		if delta > 0 && len(durations) == 0 {
			durations = append(durations, delta)
		}
	}

	if len(durations) == 0 {
		return result
	}

	sort.Float64s(durations)
	result.LMetricDurationCount = len(durations)
	result.LMetricMedian = durations[len(durations)/2]
	result.LMetricMode, result.LMetricModeShare, result.LMetricDistinctCount = LMetricModeRead(durations)

	if result.LMetricModeShare < 0.95 || result.LMetricDistinctCount > 3 {
		result.LMetricIrregular = true
	}

	return result
}

func LMetricFloatParse(value string) (float64, bool) {
	if value == "" || value == "N/A" {
		return 0, false
	}

	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil || math.IsNaN(parsed) || math.IsInf(parsed, 0) {
		return 0, false
	}

	return parsed, true
}

func LMetricTimestampRead(packet LProbePacket) (float64, bool) {
	if stamp, ok := LMetricFloatParse(packet.LProbeDtsTime); ok {
		return stamp, true
	}

	return LMetricFloatParse(packet.LProbePtsTime)
}

func LMetricModeRead(values []float64) (float64, float64, int) {
	counts := map[int]int{}
	for _, value := range values {
		bucket := int(math.Round(value * 1000000))
		counts[bucket]++
	}

	modeBucket := 0
	modeCount := 0
	for bucket, count := range counts {
		if count > modeCount {
			modeBucket = bucket
			modeCount = count
		}
	}

	return float64(modeBucket) / 1000000, float64(modeCount) / float64(len(values)), len(counts)
}
