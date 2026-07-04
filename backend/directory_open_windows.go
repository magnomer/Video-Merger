//go:build windows

package backend

import (
	"os/exec"
	"path/filepath"
)

func LDirectoryOpen(path string) error {
	cleanPath := filepath.Clean(path)
	cmd := exec.Command("cmd", "/C", "start", "", cleanPath)
	LCommandHide(cmd)
	return cmd.Start()
}
