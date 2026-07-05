package backend

import (
	"context"
	"fmt"
)

func LMergerCopyRun(LRuntimeContext context.Context, options LPreference, group LBatch, outputPath string) LMergerResult {
	err := LDiskCancelCopy(LRuntimeContext, group.LBatchClip[0].LClipPath, outputPath)
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

	result := LMergerResult{
		LDestinationPath: outputPath,
		LTaskSuccess:     true,
		LTaskMessage:     "Single-file group copied successfully.",
	}
	LMetricSet(LRuntimeContext, options, &result)

	return result
}
