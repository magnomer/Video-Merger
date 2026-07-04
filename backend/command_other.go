//go:build !windows

package backend

import "os/exec"

func LCommandHide(cmd *exec.Cmd) {
}
