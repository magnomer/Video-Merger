package backend

import (
	"context"
	"encoding/json"
	"os/exec"
)

type LProbe struct {
	LProbeStream []LProbeStream `json:"streams"`
	LProbeFormat LProbeFormat   `json:"format"`
}

type LProbeStream struct {
	LProbeIndex         int    `json:"index"`
	LProbeCodecType     string `json:"codec_type"`
	LProbeCodecName     string `json:"codec_name"`
	LProbeWidth         int    `json:"width,omitempty"`
	LProbeHeight        int    `json:"height,omitempty"`
	LProbePixelFormat   string `json:"pix_fmt,omitempty"`
	LRateAverage        string `json:"avg_frame_rate,omitempty"`
	LProbeTimeBase      string `json:"time_base,omitempty"`
	LProbeRateSample    string `json:"sample_rate,omitempty"`
	LProbeChannelCount  int    `json:"channels,omitempty"`
	LProbeChannelLayout string `json:"channel_layout,omitempty"`
}

type LProbeFormat struct {
	LProbeFilename   string `json:"filename"`
	LMetricDuration  string `json:"duration"`
	LProbeFormatName string `json:"format_name"`
}

func LProbeRun(LRuntimeContext context.Context, preference LPreference, path string) (LProbe, error) {
	ffprobePath, err := LCommandFFprobeRead(preference)
	if err != nil {
		return LProbe{}, err
	}

	cmd := exec.CommandContext(
		LRuntimeContext,
		ffprobePath,
		"-v", "quiet",
		"-print_format", "json",
		"-show_streams",
		"-show_format",
		path,
	)
	LCommandHide(cmd)

	output, err := cmd.Output()
	if err != nil {
		return LProbe{}, err
	}

	var result LProbe
	err = json.Unmarshal(output, &result)
	if err != nil {
		return LProbe{}, err
	}

	return result, nil
}
