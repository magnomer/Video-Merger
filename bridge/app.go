package bridge

import (
	"context"
	"sync"
)

type LProgram struct {
	LRuntimeContext context.Context

	LManifestWailsData []byte

	LTaskLock   sync.Mutex
	LTaskCancel context.CancelFunc
}

func LProgramCreate(LManifestWailsData []byte) *LProgram {
	return &LProgram{
		LManifestWailsData: LManifestWailsData,
	}
}

func (a *LProgram) LProgramStart(LRuntimeContext context.Context) {
	a.LRuntimeContext = LRuntimeContext
}
