package floc

import "context"

/*
Context provides safe access to the underlying context.
For more details on the underlying context see https://golang.org/pkg/context/#Context.
*/
type Context interface {
	Releaser

	// Ctx returns the underlying context.
	Ctx() context.Context

	// UpdateCtx sets the new underlying context.
	UpdateCtx(ctx context.Context)

	// Done returns a channel that's closed when the flow done.
	// Successive calls to Done return the same value.
	Done() <-chan struct{}

	// Value returns the value associated with this context for key,
	// or nil if no value is associated with key.
	Value(key interface{}) (value interface{})

	// Create a new context with value and make it the current.
	// This is an equivalent to.
	//
	//    oldCtx := ctx.Ctx()
	//    newCtx := context.WithValue(oldCtx, key, value)
	//    ctx.UpdateCtx(newCtx)
	//
	AddValue(key, value interface{})
}
