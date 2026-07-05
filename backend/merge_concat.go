package backend

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func LMergerConcatRun(LRuntimeContext context.Context, options LPreference, group LBatch, outputPath string) LMergerResult {
	if LRuntimeContext.Err() != nil {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: "Merge canceled before FFmpeg started.",
		}
	}

	listFilePath, err := LConcatCreate(options, group)
	if err != nil {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: fmt.Sprintf("Could not create concat list file: %v", err),
		}
	}
	defer LTemporaryOwnedRemove(listFilePath)

	temporaryOutputPath, err := LDiskOutputTemporaryRead(outputPath)
	if err != nil {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: fmt.Sprintf("Could not reserve temporary output: %v", err),
		}
	}
	defer os.Remove(temporaryOutputPath)

	ffmpegPath, err := LCommandFFmpegRead(options)
	if err != nil {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: fmt.Sprintf("FFmpeg path is invalid: %v", err),
		}
	}

	cmd := exec.CommandContext(
		LRuntimeContext,
		ffmpegPath,
		"-y",
		"-f", "concat",
		"-safe", "0",
		"-i", listFilePath,
		"-c", "copy",
		temporaryOutputPath,
	)
	LCommandHide(cmd)

	output, err := cmd.CombinedOutput()
	if err != nil {
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

	if err := LDiskPublishMove(temporaryOutputPath, outputPath); err != nil {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: fmt.Sprintf("Could not publish merged output: %v", err),
		}
	}

	result := LMergerResult{
		LDestinationPath: outputPath,
		LTaskSuccess:     true,
		LTaskMessage:     "Merge completed successfully.",
	}
	LMetricSet(LRuntimeContext, options, &result)

	return result
}
