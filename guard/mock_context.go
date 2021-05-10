package guard

import "github.com/workanator/go-floc/v3"

// Mock context which propagates all calls to the parent context
// but Done() returns mock channel.
type mockContext struct {
	floc.Context
	mock floc.Context
}

// Release releases the mock context.
func (ctx mockContext) Release() {
	ctx.mock.Release()
}

// Done returns the channel of the mock context.
func (ctx mockContext) Done() <-chan struct{} {
	return ctx.mock.Done()
}
