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
	status int32
	ctx    Context
	cancel context.CancelFunc
	result int32
	data   interface{}
	err    error
}

// NewControl constructs Control instance from context given.
// The function panics if the context given is nil.
func NewControl(ctx Context) Control {
	if ctx == nil {
		panic("context is nil")
	}

	oldCtx := ctx.Ctx()
	cancelCtx, cancelFunc := context.WithCancel(oldCtx)
	ctx.UpdateCtx(cancelCtx)

	return &flowControl{
		status: statusRunning,
		ctx:    ctx,
		cancel: cancelFunc,
		result: None.Int32(),
	}
}

// Release releases resources.
func (flowCtrl *flowControl) Release() {
	flowCtrl.Cancel(nil)
}

// Complete finishes the flow with success status.
func (flowCtrl *flowControl) Complete(data interface{}) {
	// Try to change status to finished
	if atomic.CompareAndSwapInt32(&flowCtrl.status, statusRunning, statusFinished) {
		flowCtrl.data = data

		// Set the result to Completed
		atomic.StoreInt32(&flowCtrl.result, Completed.Int32())
		flowCtrl.cancel()
	}
}

// Cancel cancels the execution of the flow.
func (flowCtrl *flowControl) Cancel(data interface{}) {
	// Try to change status to finished
	if atomic.CompareAndSwapInt32(&flowCtrl.status, statusRunning, statusFinished) {
		flowCtrl.data = data

		// Set the result to Canceled
		atomic.StoreInt32(&flowCtrl.result, Canceled.Int32())
		flowCtrl.cancel()
	}
}

// Fail cancels the execution of the flow with error.
func (flowCtrl *flowControl) Fail(data interface{}, err error) {
	// Try to change status to finished
	if atomic.CompareAndSwapInt32(&flowCtrl.status, statusRunning, statusFinished) {
		flowCtrl.data = data
		flowCtrl.err = err

		// Set the result to Failed
		atomic.StoreInt32(&flowCtrl.result, Failed.Int32())
		flowCtrl.cancel()
	}
}

// IsFinished tests if execution of the flow is either completed or canceled.
func (flowCtrl *flowControl) IsFinished() bool {
	return atomic.LoadInt32(&flowCtrl.status) == statusFinished
}

// Result returns the result code and the result data of the flow. The call
// to the function is effective only if the flow is finished.
func (flowCtrl *flowControl) Result() (result Result, data interface{}, err error) {
	// Load the current status and result
	s := atomic.LoadInt32(&flowCtrl.status)
	r := atomic.LoadInt32(&flowCtrl.result)
	result = Result(r)

	// Return data only if the flow is finished
	if s == statusFinished && result.IsFinished() {
		return result, flowCtrl.data, flowCtrl.err
	}

	// Otherwise return nil because reading the data field while the flow is not
	// finished may lead to unpredicted behavior, fot example reading value
	// while other goroutine is writing it.
	return result, nil, nil
}
