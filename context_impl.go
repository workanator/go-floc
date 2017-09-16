package floc

import (
	"context"
	"sync"
)

type flowContext struct {
	ctx context.Context
	mu  sync.RWMutex
}

func NewContext() Context {
	ctx := &flowContext{
		ctx: context.TODO(),
		mu:  sync.RWMutex{},
	}

	return ctx
}

// Release releases resources.
func (flowCtx *flowContext) Release() {

}

// Ctx returns the underlying context.
func (flowCtx *flowContext) Ctx() context.Context {
	flowCtx.mu.RLock()
	defer flowCtx.mu.RUnlock()

	return flowCtx.ctx
}

// UpdateCtx sets the new underlying context.
func (flowCtx *flowContext) UpdateCtx(ctx context.Context) {
	flowCtx.mu.Lock()
	defer flowCtx.mu.Unlock()

	flowCtx.ctx = ctx
}

// Done returns a channel that's closed when the flow done.
// Successive calls to Done return the same value.
func (flowCtx *flowContext) Done() <-chan struct{} {
	flowCtx.mu.RLock()
	defer flowCtx.mu.RUnlock()

	return flowCtx.ctx.Done()
}

// Value returns the value associated with this context for key,
// or nil if no value is associated with key.
func (flowCtx *flowContext) Value(key interface{}) (value interface{}) {
	flowCtx.mu.RLock()
	ctx := flowCtx.ctx
	flowCtx.mu.RUnlock()

	return ctx.Value(key)
}

// Create a new context with value and make it the current.
func (flowCtx *flowContext) AddValue(key, value interface{}) {
	flowCtx.mu.Lock()
	defer flowCtx.mu.Unlock()

	flowCtx.ctx = context.WithValue(flowCtx.ctx, key, value)
}
