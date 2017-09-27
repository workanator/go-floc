package floc

import (
	"context"
	"sync/atomic"
)

const (
	statusRunning  = 0
	statusFinished = 1
)

type flowControl struct {
	context.Context
	cancel context.CancelFunc
	status int32
	result int32 // the underlying type of floc.Result is int32
	data   interface{}
}

// NewFlow creates a new instance of the flow control. Once the flow is finished
// the instance should not be copied or reused for controlling other flows.
func NewFlow() Flow {
	ctx, cancel := context.WithCancel(context.TODO())

	return &flowControl{
		Context: ctx,
		cancel:  cancel,
		status:  statusRunning,
		result:  None.Int32(), // floc.None may be not 0 so do explicit assign
	}
}

// Done returns a channel that's closed when the flow done.
// Successive calls to Done return the same value.
func (flow *flowControl) Done() <-chan struct{} {
	return flow.Context.Done()
}

// Release finishes the flow and releases all underlying resources.
func (flow *flowControl) Release() {
	flow.Cancel(nil) // That has no effect if the flow is already finished
}

// Complete finishes the flow with success status and stops
// execution of further jobs if any.
func (flow *flowControl) Complete(data interface{}) {
	flow.finish(Completed, data)
}

// Cancel cancels execution of the flow.
func (flow *flowControl) Cancel(data interface{}) {
	flow.finish(Canceled, data)
}

// IsFinished tests if execution of the flow is either completed or canceled.
func (flow *flowControl) IsFinished() bool {
	r := atomic.LoadInt32(&flow.result)
	return Result(r).IsFinished()
}

// Result returns the result code and the result data of the flow.
func (flow *flowControl) Result() (result Result, data interface{}) {
	// Load the current result
	r := atomic.LoadInt32(&flow.result)
	result = Result(r)

	// Return data only if the flow is finished
	if result.IsFinished() {
		return result, flow.data
	}

	// Otherwise return nil because reading the data field while the flow is not
	// finished may lead to unpredicted behavior, fot example reading value
	// while other goroutine is writing it.
	return result, nil
}

func (flow *flowControl) finish(result Result, data interface{}) {
	// Try to change status to finished
	if atomic.CompareAndSwapInt32(&flow.status, statusRunning, statusFinished) {
		// Set data
		flow.data = data

		// Set the result and cancel the context
		atomic.StoreInt32(&flow.result, result.Int32())
		flow.cancel()
	}
}
