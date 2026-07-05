package backend

import "fmt"

type LSignature struct {
	LProbeStream []LSignatureStream
}

type LSignatureStream struct {
	LSignatureCodecType     string
	LSignatureCodecName     string
	LSignatureWidth         int
	LSignatureHeight        int
	LSignaturePixelFormat   string
	LRateAverage            string
	LProbeTimeBase          string
	LSignatureRateSample    string
	LSignatureChannelCount  int
	LSignatureChannelLayout string
}

type LCompatibilityStream struct {
	LCompatibilityCaution []string
	LCompatibilityWarning []string
}

func LSignatureBuild(probe LProbe) LSignature {
	signature := LSignature{}

	for _, stream := range probe.LProbeStream {
		signature.LProbeStream = append(signature.LProbeStream, LSignatureStream{
			LSignatureCodecType:     stream.LProbeCodecType,
			LSignatureCodecName:     stream.LProbeCodecName,
			LSignatureWidth:         stream.LProbeWidth,
			LSignatureHeight:        stream.LProbeHeight,
			LSignaturePixelFormat:   stream.LProbePixelFormat,
			LRateAverage:            stream.LRateAverage,
			LProbeTimeBase:          stream.LProbeTimeBase,
			LSignatureRateSample:    stream.LProbeRateSample,
			LSignatureChannelCount:  stream.LProbeChannelCount,
			LSignatureChannelLayout: stream.LProbeChannelLayout,
		})
	}

	return signature
}

func LSignatureStreamCompare(expected LSignatureStream, actual LSignatureStream) LCompatibilityStream {
	result := LCompatibilityStream{}

	if expected.LSignatureCodecType != actual.LSignatureCodecType {
		result.LCompatibilityWarning = append(result.LCompatibilityWarning, fmt.Sprintf("codec type differs: %s vs %s", expected.LSignatureCodecType, actual.LSignatureCodecType))
	}

	if expected.LSignatureCodecName != actual.LSignatureCodecName {
		result.LCompatibilityWarning = append(result.LCompatibilityWarning, fmt.Sprintf("codec differs: %s vs %s", expected.LSignatureCodecName, actual.LSignatureCodecName))
	}

	if expected.LSignatureWidth != actual.LSignatureWidth {
		result.LCompatibilityWarning = append(result.LCompatibilityWarning, fmt.Sprintf("width differs: %d vs %d", expected.LSignatureWidth, actual.LSignatureWidth))
	}

	if expected.LSignatureHeight != actual.LSignatureHeight {
		result.LCompatibilityWarning = append(result.LCompatibilityWarning, fmt.Sprintf("height differs: %d vs %d", expected.LSignatureHeight, actual.LSignatureHeight))
	}

	if expected.LSignaturePixelFormat != actual.LSignaturePixelFormat {
		result.LCompatibilityWarning = append(result.LCompatibilityWarning, fmt.Sprintf("pixel format differs: %s vs %s", expected.LSignaturePixelFormat, actual.LSignaturePixelFormat))
	}

	if expected.LProbeTimeBase != actual.LProbeTimeBase {
		result.LCompatibilityWarning = append(result.LCompatibilityWarning, fmt.Sprintf("time base differs: %s vs %s", expected.LProbeTimeBase, actual.LProbeTimeBase))
	}

	if expected.LSignatureRateSample != actual.LSignatureRateSample {
		result.LCompatibilityWarning = append(result.LCompatibilityWarning, fmt.Sprintf("sample rate differs: %s vs %s", expected.LSignatureRateSample, actual.LSignatureRateSample))
	}

	if expected.LSignatureChannelCount != actual.LSignatureChannelCount {
		result.LCompatibilityWarning = append(result.LCompatibilityWarning, fmt.Sprintf("audio channels differ: %d vs %d", expected.LSignatureChannelCount, actual.LSignatureChannelCount))
	}

	if expected.LRateAverage != actual.LRateAverage {
		frameRateResult := LRateCompare(expected.LRateAverage, actual.LRateAverage)

		if frameRateResult.LRateWarningState {
			result.LCompatibilityWarning = append(result.LCompatibilityWarning, frameRateResult.LTaskMessage)
		} else if frameRateResult.LTaskMessage != "" && !frameRateResult.LRateSmallState {
			result.LCompatibilityCaution = append(result.LCompatibilityCaution, frameRateResult.LTaskMessage)
		}
	}

	if expected.LSignatureChannelLayout != actual.LSignatureChannelLayout {
		result.LCompatibilityCaution = append(result.LCompatibilityCaution, fmt.Sprintf("channel layout metadata differs: %s vs %s", expected.LSignatureChannelLayout, actual.LSignatureChannelLayout))
	}

	return result
}
