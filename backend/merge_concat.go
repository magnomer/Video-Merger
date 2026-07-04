package backend

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func LMergerConcatRun(LRuntimeContext context.Context, group LBatch, outputPath string) LMergerResult {
	if LRuntimeContext.Err() != nil {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: "Merge canceled before FFmpeg started.",
		}
	}

	listFilePath, err := LConcatCreate(group)
	if err != nil {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: fmt.Sprintf("Could not create concat list file: %v", err),
		}
	}
	defer os.Remove(listFilePath)

	cmd := exec.CommandContext(
		LRuntimeContext,
		"ffmpeg",
		"-f", "concat",
		"-safe", "0",
		"-i", listFilePath,
		"-c", "copy",
		outputPath,
	)
	LCommandHide(cmd)

	output, err := cmd.CombinedOutput()
	if err != nil {
		os.Remove(outputPath)

		if LRuntimeContext.Err() != nil {
			return LMergerResult{
				LTaskSuccess: false,
				LTaskMessage: "Merge canceled.",
			}
		}

		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: fmt.Sprintf("FFmpeg merge failed: %v\n%s", err, string(output)),
		}
	}

	result := LMergerResult{
		LDestinationPath: outputPath,
		LTaskSuccess:     true,
		LTaskMessage:     "Merge completed successfully.",
	}
	LMetricSet(LRuntimeContext, &result)

	return result
}
