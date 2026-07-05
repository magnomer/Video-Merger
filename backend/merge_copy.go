package backend

import (
	"context"
	"fmt"
	"os"
)

func LMergerCopyRun(LRuntimeContext context.Context, options LPreference, group LBatch, outputPath string) LMergerResult {
	temporaryOutputPath, err := LDiskOutputTemporaryRead(outputPath)
	if err != nil {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: fmt.Sprintf("Could not reserve temporary output: %v", err),
		}
	}
	defer os.Remove(temporaryOutputPath)

	err = LDiskCancelCopy(LRuntimeContext, group.LBatchClip[0].LClipPath, temporaryOutputPath)
	if err != nil {
		if LRuntimeContext.Err() != nil {
			return LMergerResult{
				LTaskSuccess: false,
				LTaskMessage: "Copy canceled.",
			}
		}

		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: fmt.Sprintf("Could not copy single file: %v", err),
		}
	}

	if err := LDiskPublishMove(temporaryOutputPath, outputPath); err != nil {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: fmt.Sprintf("Could not publish copied output: %v", err),
		}
	}

	result := LMergerResult{
		LDestinationPath: outputPath,
		LTaskSuccess:     true,
		LTaskMessage:     "Single-file group copied successfully.",
	}
	LMetricSet(LRuntimeContext, options, &result)

	return result
}
