//go:build !windows

package backend

import (
	"os/exec"
	"runtime"
)

func LDirectoryOpen(path string) error {
	if runtime.GOOS == "darwin" {
		return exec.Command("open", path).Start()
	}

	return exec.Command("xdg-open", path).Start()
}
