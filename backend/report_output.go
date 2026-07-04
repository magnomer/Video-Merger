package backend

import (
	"path/filepath"
	"strings"
)

func LReportOutputSet(report LReport, options LPreference) (LReport, error) {
	if report.LReportKind != "analysis" {
		return report, nil
	}

	options.LPreferenceOutput = strings.TrimSpace(options.LPreferenceOutput)
	options.LPreferenceSuffix = strings.TrimSpace(options.LPreferenceSuffix)

	if err := LSuffixCheck(options.LPreferenceSuffix); err != nil {
		return report, err
	}

	for index := range report.LReportGroup {
		report.LReportGroup[index].LReportOutputTitle = "Planned output"
		report.LReportGroup[index].LReportOutputText = LReportOutputResolve(report.LReportGroup[index], options)
	}

	return report, nil
}

func LReportOutputResolve(group LReportGroup, options LPreference) string {
	if !options.LPreferenceMirror && options.LPreferenceOutput == "" {
		return "Output folder is required unless Same as input is checked."
	}

	outputFolder := options.LPreferenceOutput
	firstFile := LReportFileFirstRead(group)

	if options.LPreferenceMirror {
		if firstFile.LReportPath == "" {
			outputFolder = "."
		} else {
			outputFolder = filepath.Dir(firstFile.LReportPath)
		}
	} else if group.LReportDirectory != "" {
		outputFolder = filepath.Join(outputFolder, group.LReportDirectory)
	}

	extension := filepath.Ext(firstFile.LReportPath)
	if extension == "" {
		extension = filepath.Ext(firstFile.LReportName)
	}
	if extension == "" {
		extension = ".mp4"
	}

	return LDestinationFind(outputFolder, group.LReportName+options.LPreferenceSuffix, extension)
}

func LReportFileFirstRead(group LReportGroup) LReportFile {
	if len(group.LReportFile) == 0 {
		return LReportFile{}
	}

	firstFile := group.LReportFile[0]
	for _, file := range group.LReportFile[1:] {
		if file.LReportNumber < firstFile.LReportNumber {
			firstFile = file
		}
	}

	return firstFile
}
