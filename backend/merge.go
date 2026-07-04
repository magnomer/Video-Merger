package backend

import (
	"context"
	"fmt"
)

func LMergerBatchRun(LRuntimeContext context.Context, group LBatch, outputFolder string, suffix string, ignoreCautions bool, forceMergeWarnings bool) LMergerResult {
	compatibility, err := LCompatibilityCheck(LRuntimeContext, group)
	if err != nil {
		if LRuntimeContext.Err() != nil {
			return LMergerResult{
				LTaskSuccess: false,
				LTaskMessage: "Merge canceled during compatibility check.",
			}
		}

		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: fmt.Sprintf("Compatibility check failed: %v", err),
		}
	}

	outputPath := LDestinationFind(outputFolder, group.LBatchName+suffix, group.LClipExtension)

	return LMergerCancelRun(LRuntimeContext, group, compatibility, outputPath, ignoreCautions, forceMergeWarnings)
}

func LMergerCancelRun(
	LRuntimeContext context.Context,
	group LBatch,
	compatibility LCompatibility,
	outputPath string,
	ignoreCautions bool,
	forceMergeWarnings bool,
) LMergerResult {
	if gateResult, ok := LMergerGateCheck(LRuntimeContext, group, compatibility, outputPath, ignoreCautions, forceMergeWarnings); !ok {
		return gateResult
	}

	if len(group.LBatchClip) == 1 {
		return LMergerCopyRun(LRuntimeContext, group, outputPath)
	}

	return LMergerConcatRun(LRuntimeContext, group, outputPath)
}
