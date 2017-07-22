package floc

import (
	"context"
	"sync"
)

/*
FlowControl allows to control execution of the flow. Once the flow is finished
the instance of FlowControl should not be copied or reused for controling
other flows.
*/
type FlowControl struct {
	sync.RWMutex
	context.Context
	cancel context.CancelFunc
	result Result
	data   interface{}
}

// NewFlowControl creates a new instance of the flow control.
func NewFlowControl() Flow {
	ctx, cancel := context.WithCancel(context.TODO())

	return &FlowControl{
		Context: ctx,
		cancel:  cancel,
	}
}

// Done returns a channel that's closed when the flow done.
// Successive calls to Done return the same value.
func (flow *FlowControl) Done() <-chan struct{} {
	return flow.Context.Done()
}

// Release finishes the flow and releases all underlying resources.
func (flow *FlowControl) Release() {
	flow.Cancel(nil)
}

// Complete finishes the flow with success status and stops
// execution of further jobs if any.
func (flow *FlowControl) Complete(data interface{}) {
	flow.RWMutex.Lock()
	defer flow.RWMutex.Unlock()

	if flow.result.IsNone() {
		flow.result = Completed
		flow.data = data
		flow.cancel()
	}
}

// Cancel cancels execution of the flow.
func (flow *FlowControl) Cancel(data interface{}) {
	flow.RWMutex.Lock()
	defer flow.RWMutex.Unlock()

	if flow.result.IsNone() {
		flow.result = Canceled
		flow.data = data
		flow.cancel()
	}
}

// IsFinished tests if execution of the flow is either completed or canceled.
func (flow *FlowControl) IsFinished() bool {
	flow.RWMutex.RLock()
	defer flow.RWMutex.RUnlock()

	return flow.result.IsCompleted() || flow.result.IsCanceled()
}

// Result returns the result code and the result data of the flow.
func (flow *FlowControl) Result() (result Result, data interface{}) {
	flow.RWMutex.RLock()
	defer flow.RWMutex.RUnlock()

	return flow.result, flow.data
}
