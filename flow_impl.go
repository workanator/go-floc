package floc

import (
	"context"
	"sync/atomic"
	"unsafe"
)

type flowControl struct {
	context.Context
	cancel context.CancelFunc
	result Result
	data   interface{}
}

// NewFlow creates a new instance of the flow control. Once the flow is finished
// the instance should not be copied or reused for controlling other flows.
func NewFlow() Flow {
	ctx, cancel := context.WithCancel(context.TODO())

	return &flowControl{
		Context: ctx,
		cancel:  cancel,
	}
}

// Done returns a channel that's closed when the flow done.
// Successive calls to Done return the same value.
func (flow *flowControl) Done() <-chan struct{} {
	return flow.Context.Done()
}

// Release finishes the flow and releases all underlying resources.
func (flow *flowControl) Release() {
	flow.Cancel(nil)
}

// Complete finishes the flow with success status and stops
// execution of further jobs if any.
func (flow *flowControl) Complete(data interface{}) {
	// Try to chnage the result from None to Completed and if it's successful
	// finish the flow.
	if atomic.CompareAndSwapInt32(flow.resultAsInt32Ptr(), None.Int32(), Completed.Int32()) {
		flow.data = data
		flow.cancel()
	}
}

// Cancel cancels execution of the flow.
func (flow *flowControl) Cancel(data interface{}) {
	// Try to chnage the result from None to Canceled and if it's successful
	// finish the flow.
	if atomic.CompareAndSwapInt32(flow.resultAsInt32Ptr(), None.Int32(), Canceled.Int32()) {
		flow.data = data
		flow.cancel()
	}
}

// IsFinished tests if execution of the flow is either completed or canceled.
func (flow *flowControl) IsFinished() bool {
	return Result(atomic.LoadInt32((*int32)(unsafe.Pointer(&flow.result)))).IsFinished()
}

// Result returns the result code and the result data of the flow.
func (flow *flowControl) Result() (result Result, data interface{}) {
	// Return data only if the flow is finished
	if result = Result(atomic.LoadInt32(flow.resultAsInt32Ptr())); result.IsFinished() {
		return result, flow.data
	}

	// Otherwise return nil because reading the data field while the flow is not
	// finished may lead to unpredicted behavior, fot example reading value
	// while other goroutine is writing it.
	return result, nil
}

// resultAsInt32Ptr returns the pointer to result field as *int32. Required for
// compatibility with atomic operations.
func (flow *flowControl) resultAsInt32Ptr() *int32 {
	return (*int32)(unsafe.Pointer(&flow.result))
}
