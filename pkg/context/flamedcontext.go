package context

import (
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/iface"
	"sync"
	"time"
)

type FlamedContext struct {
	mMutex                   sync.Mutex
	mFlamed                  *flamed.Flamed
	mTransactionProcessorMap map[string]iface.ITransactionProcessor
	mGlobalRequestTimeout    time.Duration
}

func (f *FlamedContext) Flamed() *flamed.Flamed {
	f.mMutex.Lock()
	defer f.mMutex.Unlock()
	return f.mFlamed
}

func (f *FlamedContext) TransactionProcessorMap() map[string]iface.ITransactionProcessor {
	f.mMutex.Lock()
	defer f.mMutex.Unlock()
	return f.mTransactionProcessorMap
}

func (f *FlamedContext) GlobalRequestTimeout() time.Duration {
	f.mMutex.Lock()
	defer f.mMutex.Unlock()
	return f.mGlobalRequestTimeout
}

func (f *FlamedContext) AddTP(tp iface.ITransactionProcessor) {
	f.mMutex.Lock()
	defer f.mMutex.Unlock()
	f.mTransactionProcessorMap[tp.FamilyName()+"::"+tp.FamilyVersion()] = tp
}

func (f *FlamedContext) SetFlamed(fd *flamed.Flamed) {
	f.mMutex.Lock()
	defer f.mMutex.Unlock()
	f.mFlamed = fd
}

func (f *FlamedContext) SetGlobalTimeout(t time.Duration) {
	f.mMutex.Lock()
	defer f.mMutex.Unlock()
	f.mGlobalRequestTimeout = t
}

func NewFlamedContext() *FlamedContext {
	return &FlamedContext{mTransactionProcessorMap: map[string]iface.ITransactionProcessor{}}
}
