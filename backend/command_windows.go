//go:build windows

package backend

import (
	"os/exec"
	"syscall"
)

func LCommandHide(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
}
