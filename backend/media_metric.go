package backend

import (
	"context"
	"os"
	"strconv"
)

type LMetric struct {
	LMetricSize     int64
	LMetricDuration float64
	LMetricLoudness *float64
}

func LMetricMeasure(LRuntimeContext context.Context, path string) LMetric {
	metrics := LMetric{}

	info, err := os.Stat(path)
	if err == nil {
		metrics.LMetricSize = info.Size()
	}

	probe, err := LProbeRun(LRuntimeContext, path)
	if err == nil {
		duration, parseErr := strconv.ParseFloat(probe.LProbeFormat.LMetricDuration, 64)
		if parseErr == nil {
			metrics.LMetricDuration = duration
		}
	}

	if LRuntimeContext.Err() == nil {
		metrics.LMetricLoudness = LLoudnessMeasure(LRuntimeContext, path)
	}

	return metrics
}

func LMetricSet(LRuntimeContext context.Context, result *LMergerResult) {
	metrics := LMetricMeasure(LRuntimeContext, result.LDestinationPath)
	result.LMetricSize = metrics.LMetricSize
	result.LMetricDuration = metrics.LMetricDuration
	result.LMetricLoudness = metrics.LMetricLoudness
}
