package bridge

import (
	"strings"
	"video-merger/backend"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *LProgram) LPickerFileOpen() ([]string, error) {
	return runtime.OpenMultipleFilesDialog(a.LRuntimeContext, runtime.OpenDialogOptions{
		Title: "Select video files",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Video files",
				Pattern:     LPickerPatternRead(),
			},
		},
	})
}

func (a *LProgram) LPickerFolderOpen() (string, error) {
	return runtime.OpenDirectoryDialog(a.LRuntimeContext, runtime.OpenDialogOptions{
		Title: "Select folder",
	})
}

func (a *LProgram) LPickerFFmpegOpen() (string, error) {
	return runtime.OpenFileDialog(a.LRuntimeContext, runtime.OpenDialogOptions{
		Title: "Select ffmpeg executable",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "FFmpeg executable",
				Pattern:     "ffmpeg.exe;ffmpeg",
			},
			{
				DisplayName: "Executable files",
				Pattern:     "*.exe",
			},
		},
	})
}

func LPickerPatternRead() string {
	patterns := make([]string, 0, len(backend.LClipSupportedExtensionList))
	for _, extension := range backend.LClipSupportedExtensionList {
		patterns = append(patterns, "*"+extension)
	}

	return strings.Join(patterns, ";")
}
