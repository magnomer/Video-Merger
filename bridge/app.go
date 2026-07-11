package bridge

import (
	"context"
	"sync"
	"video-merger/backend"
)

type LProgram struct {
	LRuntimeContext context.Context

	LVersionData []byte

	LTaskLock   sync.Mutex
	LTaskCancel context.CancelFunc

	LInspectionLock       sync.Mutex
	LInspectionPreference backend.LPreference
	LInspectionResult     backend.LRouteResult
	LInspectionReady      bool
}

func LProgramCreate(LVersionData []byte) *LProgram {
	return &LProgram{
		LVersionData: LVersionData,
	}
}

func (a *LProgram) LProgramStart(LRuntimeContext context.Context) {
	a.LRuntimeContext = LRuntimeContext
}

func (a *LProgram) LProgramStop(LRuntimeContext context.Context) bool {
	backend.LAssetPreviewStop("")
	a.LTaskStop()
	return false
}

func (a *LProgram) LProgramShutdown(LRuntimeContext context.Context) {
	backend.LAssetPreviewStop("")
	a.LTaskStop()
}
