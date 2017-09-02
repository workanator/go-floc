package floc

import (
	"context"
	"sync"
)

type flowContext struct {
	context.Context
	sync.RWMutex
}

func NewContext() Context {
	return &flowContext{
		Context: context.TODO(),
		RWMutex: sync.RWMutex{},
	}
}

// Release releases resources.
func (flowCtx flowContext) Release() {

}

// Ctx returns the underlying context.
func (flowCtx flowContext) Ctx() context.Context {
	flowCtx.RLock()
	defer flowCtx.RUnlock()

	return flowCtx.Context
}

// UpdateCtx sets the new underlying context.
func (flowCtx flowContext) UpdateCtx(ctx context.Context) {
	flowCtx.Lock()
	defer flowCtx.Unlock()

	flowCtx.Context = ctx
}

// Value returns the value associated with this context for key,
// or nil if no value is associated with key.
func (flowCtx flowContext) Value(key interface{}) (value interface{}) {
	flowCtx.RLock()
	defer flowCtx.RUnlock()

	return flowCtx.Context.Value(key)
}

// Create a new context with value and make it the current.
func (flowCtx flowContext) AddValue(key, value interface{}) {
	flowCtx.Lock()
	defer flowCtx.Unlock()

	newCtx := context.WithValue(flowCtx.Context, key, value)
	flowCtx.Context = newCtx
}
