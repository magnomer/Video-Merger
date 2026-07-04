package bridge

import (
	"context"
	"errors"
)

func (a *LProgram) lTaskStart() (context.Context, context.CancelFunc, error) {
	a.LTaskLock.Lock()
	defer a.LTaskLock.Unlock()

	if a.LTaskCancel != nil {
		return nil, nil, errors.New("processing is already running")
	}

	LRuntimeContext, cancel := context.WithCancel(context.Background())
	a.LTaskCancel = cancel

	return LRuntimeContext, cancel, nil
}

func (a *LProgram) lTaskReset() {
	a.LTaskLock.Lock()
	defer a.LTaskLock.Unlock()

	a.LTaskCancel = nil
}

func (a *LProgram) LTaskStop() string {
	a.LTaskLock.Lock()
	defer a.LTaskLock.Unlock()

	if a.LTaskCancel == nil {
		return "No processing job is currently running."
	}

	a.LTaskCancel()

	return "Stop requested."
}
