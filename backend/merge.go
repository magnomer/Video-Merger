package backend

import (
	"context"
	"fmt"
)

func LMergerBatchRun(LRuntimeContext context.Context, options LPreference, group LBatch, outputFolder string, suffix string, ignoreCautions bool, forceMergeWarnings bool) LMergerResult {
	compatibility, err := LCompatibilityCheck(LRuntimeContext, options, group)
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

	return LMergerCancelRun(LRuntimeContext, options, group, compatibility, outputPath, ignoreCautions, forceMergeWarnings)
}

func LMergerCancelRun(
	LRuntimeContext context.Context,
	options LPreference,
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
		return LMergerCopyRun(LRuntimeContext, options, group, outputPath)
	}

	return LMergerConcatRun(LRuntimeContext, options, group, outputPath)
}
