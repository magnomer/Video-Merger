package backend

import (
	"fmt"
	"path/filepath"
)

func LDestinationBatchRead(options LPreference, group LBatch) string {
	if options.LPreferenceMirror {
		if len(group.LBatchClip) == 0 {
			return "."
		}

		return filepath.Dir(group.LBatchClip[0].LClipPath)
	}

	if group.LBatchDirectory == "" {
		return options.LPreferenceOutput
	}

	return filepath.Join(options.LPreferenceOutput, group.LBatchDirectory)
}

func LDestinationFind(outputFolder string, baseName string, extension string) string {
	outputPath := filepath.Join(outputFolder, baseName+extension)

	if !LDiskCheck(outputPath) {
		return outputPath
	}

	counter := 2

	for {
		candidate := filepath.Join(
			outputFolder,
			fmt.Sprintf("%s (%d)%s", baseName, counter, extension),
		)

		if !LDiskCheck(candidate) {
			return candidate
		}

		counter++
	}
}

func LDestinationPlanResolve(group LBatchResult, suffix string) LBatchResult {
	if group.LBatchPlan == "" || !LDiskCheck(group.LBatchPlan) {
		return group
	}

	folder := filepath.Dir(group.LBatchPlan)
	group.LBatchPlan = LDestinationFind(folder, group.LBatchName+suffix, LBatchExtensionRead(group.LBatchClip))

	return group
}
