package backend

import (
	"fmt"
	"path/filepath"
	"strings"
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

func LDestinationPlanCreate(options LPreference, group LBatch) (string, error) {
	options.LPreferenceOutput = strings.TrimSpace(options.LPreferenceOutput)
	options.LPreferenceSuffix = strings.TrimSpace(options.LPreferenceSuffix)

	if err := LSuffixCheck(options.LPreferenceSuffix); err != nil {
		return "", err
	}
	if !options.LPreferenceMirror && options.LPreferenceOutput == "" {
		return "", fmt.Errorf("output folder is required unless Same as input is checked")
	}

	outputFolder := LDestinationBatchRead(options, group)
	if strings.TrimSpace(outputFolder) == "" {
		return "", fmt.Errorf("output folder is empty")
	}

	return LDestinationFind(outputFolder, group.LBatchName+options.LPreferenceSuffix, group.LClipExtension), nil
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
