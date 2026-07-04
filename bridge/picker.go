package bridge

import "github.com/wailsapp/wails/v2/pkg/runtime"

func (a *LProgram) LPickerFileOpen() ([]string, error) {
	return runtime.OpenMultipleFilesDialog(a.LRuntimeContext, runtime.OpenDialogOptions{
		Title: "Select video files",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Video files",
				Pattern:     "*.mp4;*.mov;*.mkv;*.m4v",
			},
		},
	})
}

func (a *LProgram) LPickerFolderOpen() (string, error) {
	return runtime.OpenDirectoryDialog(a.LRuntimeContext, runtime.OpenDialogOptions{
		Title: "Select folder",
	})
}
