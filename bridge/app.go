package bridge

import (
	"context"
	"sync"
	"video-merger/backend"
)

type LProgram struct {
	LRuntimeContext context.Context

	LManifestWailsData []byte

	LTaskLock   sync.Mutex
	LTaskCancel context.CancelFunc

	LInspectionLock       sync.Mutex
	LInspectionPreference backend.LPreference
	LInspectionResult     backend.LRouteResult
	LInspectionReady      bool
}

func LProgramCreate(LManifestWailsData []byte) *LProgram {
	return &LProgram{
		LManifestWailsData: LManifestWailsData,
	}
}

func (a *LProgram) LProgramStart(LRuntimeContext context.Context) {
	a.LRuntimeContext = LRuntimeContext
}
