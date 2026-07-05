package backend

import (
	"bufio"
	"context"
	"io"
	"os/exec"
	"regexp"
	"strconv"
)

var LLoudnessPattern = regexp.MustCompile(`(?m)^\s*I:\s*(-?\d+(?:\.\d+)?)\s+LUFS`)

func LLoudnessMeasure(LRuntimeContext context.Context, preference LPreference, path string) *float64 {
	cmd := exec.CommandContext(
		LRuntimeContext,
		LCommandFFmpegRead(preference),
		"-hide_banner",
		"-nostats",
		"-i", path,
		"-vn",
		"-filter:a", "ebur128",
		"-f", "null",
		"-",
	)
	LCommandHide(cmd)

	pipe, err := cmd.StderrPipe()
	if err != nil {
		return nil
	}

	if err := cmd.Start(); err != nil {
		return nil
	}

	lastValue := LLoudnessScan(pipe)
	if err := cmd.Wait(); err != nil {
		return nil
	}

	return lastValue
}

func LLoudnessScan(reader io.Reader) *float64 {
	scanner := bufio.NewScanner(reader)
	var lastValue *float64

	for scanner.Scan() {
		matches := LLoudnessPattern.FindStringSubmatch(scanner.Text())
		if len(matches) == 0 {
			continue
		}

		value, err := strconv.ParseFloat(matches[1], 64)
		if err != nil {
			continue
		}

		captured := value
		lastValue = &captured
	}

	return lastValue
}
