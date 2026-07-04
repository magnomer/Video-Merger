package bridge

import "video-merger/backend"

func (a *LProgram) LDirectorySourceOpen(sourceText string) error {
	sourceFolder, err := backend.LDirectorySourceResolve(sourceText)
	if err != nil {
		return err
	}

	return backend.LDirectoryOpen(sourceFolder)
}

func (a *LProgram) LDirectoryOutputOpen(outputText string, sourceText string, sameAsInput bool) error {
	outputFolder, err := backend.LDirectoryOutputResolve(outputText, sourceText, sameAsInput)
	if err != nil {
		return err
	}

	return backend.LDirectoryOpen(outputFolder)
}
