package flow

import (
	"context"
	"sync"

	floc "github.com/workanator/go-floc"
)

type defaultFlow struct {
	sync.RWMutex
	context.Context
	cancel context.CancelFunc
	result floc.Result
	data   interface{}
}

// New creates a new instance of the default flow.
func New() floc.Flow {
	ctx, cancel := context.WithCancel(context.TODO())

	return &defaultFlow{
		Context: ctx,
		cancel:  cancel,
	}
}

// Done returns a channel that's closed when the flow done.
// Successive calls to Done return the same value.
func (f *defaultFlow) Done() <-chan struct{} {
	return f.Context.Done()
}

// Close finishes the flow and releases all underlying resources.
func (f *defaultFlow) Close() {
	f.Cancel(nil)
}

// Complete finishes the flow with success status and stops
// execution of futher jobs if any.
func (f *defaultFlow) Complete(data interface{}) {
	f.RWMutex.Lock()
	defer f.RWMutex.Unlock()

	if f.result == floc.None {
		f.result = floc.Completed
		f.data = data
		f.cancel()
	}
}

// Cancel cancels the execution of the flow.
func (f *defaultFlow) Cancel(data interface{}) {
	f.RWMutex.Lock()
	defer f.RWMutex.Unlock()

	if f.result == floc.None {
		f.result = floc.Canceled
		f.data = data
		f.cancel()
	}
}

// Tests if the execution of the flow is either completed or canceled.
func (f *defaultFlow) IsFinished() bool {
	f.RWMutex.RLock()
	defer f.RWMutex.RUnlock()

	return f.result == floc.Completed || f.result == floc.Canceled
}

// Returns the result code and the result data of the flow.
func (f *defaultFlow) Result() (result floc.Result, data interface{}) {
	f.RWMutex.RLock()
	defer f.RWMutex.RUnlock()

	return f.result, f.data
}
