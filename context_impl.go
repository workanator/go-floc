package floc

import (
	"context"
	"sync"
	"sync/atomic"
)

type flowContext struct {
	ctx atomic.Value
	mu  sync.Mutex
}

func NewContext() Context {
	ctx := &flowContext{
		ctx: atomic.Value{},
		mu:  sync.Mutex{},
	}

	ctx.ctx.Store(context.TODO())

	return ctx
}

// Release releases resources.
func (flowCtx flowContext) Release() {

}

// Ctx returns the underlying context.
func (flowCtx flowContext) Ctx() context.Context {
	return flowCtx.ctx.Load().(context.Context)
}

// UpdateCtx sets the new underlying context.
func (flowCtx flowContext) UpdateCtx(ctx context.Context) {
	flowCtx.mu.Lock()
	defer flowCtx.mu.Unlock()

	flowCtx.ctx.Store(ctx)
}

// Done returns a channel that's closed when the flow done.
// Successive calls to Done return the same value.
func (flowCtx flowContext) Done() <-chan struct{} {
	ctx := flowCtx.ctx.Load().(context.Context)
	return ctx.Done()
}

// Value returns the value associated with this context for key,
// or nil if no value is associated with key.
func (flowCtx flowContext) Value(key interface{}) (value interface{}) {
	ctx := flowCtx.ctx.Load().(context.Context)
	return ctx.Value(key)
}

// Create a new context with value and make it the current.
func (flowCtx flowContext) AddValue(key, value interface{}) {
	flowCtx.mu.Lock()
	defer flowCtx.mu.Unlock()

	oldCtx := flowCtx.ctx.Load().(context.Context)
	newCtx := context.WithValue(oldCtx, key, value)
	flowCtx.ctx.Store(newCtx)
}
