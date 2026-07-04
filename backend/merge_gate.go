package backend

import (
	"context"
	"fmt"
	"os"
)

func LMergerGateCheck(
	LRuntimeContext context.Context,
	group LBatch,
	compatibility LCompatibility,
	outputPath string,
	ignoreCautions bool,
	forceMergeWarnings bool,
) (LMergerResult, bool) {
	if LRuntimeContext.Err() != nil {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: "Merge canceled before starting.",
		}, false
	}

	if len(group.LBatchClip) == 0 {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: "Group has no files.",
		}, false
	}

	err := os.MkdirAll(LDirectoryPathRead(outputPath), 0755)
	if err != nil {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: fmt.Sprintf("Could not create output folder: %v", err),
		}, false
	}

	if len(compatibility.LCompatibilityCaution) > 0 && !ignoreCautions {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: "Group has cautions. Check Ignore cautions to continue.",
		}, false
	}

	if len(compatibility.LCompatibilityWarning) > 0 && !forceMergeWarnings {
		return LMergerResult{
			LTaskSuccess: false,
			LTaskMessage: "Group has warnings. Check Force merge despite warnings to continue.",
		}, false
	}

	return LMergerResult{}, true
}
