package floc

import (
	"context"
	"sync/atomic"
)

type flowControl struct {
	ctx    Context
	cancel context.CancelFunc
	result int32
	data   interface{}
}

// NewControl constructs Control instance from context given.
func NewControl(ctx Context) Control {
	oldCtx := ctx.Ctx()
	newCtx, cancel := context.WithCancel(oldCtx)
	ctx.UpdateCtx(newCtx)

	return &flowControl{
		ctx:    ctx,
		cancel: cancel,
		result: None.Int32(),
	}
}

// Release releases resources.
func (flowCtrl flowControl) Release() {
	flowCtrl.Cancel(nil)
	flowCtrl.ctx.Release()
}

// Complete finishes the flow with success status.
func (flowCtrl flowControl) Complete(data interface{}) {
	// Try to change the result from None to Completed and if it's successful
	// finish the flow.
	if atomic.CompareAndSwapInt32(&flowCtrl.result, None.Int32(), Completed.Int32()) {
		flowCtrl.data = data
		flowCtrl.cancel()
	}
}

// Cancel cancels the execution of the flow.
func (flowCtrl flowControl) Cancel(data interface{}) {
	// Try to change the result from None to Canceled and if it's successful
	// finish the flow.
	if atomic.CompareAndSwapInt32(&flowCtrl.result, None.Int32(), Canceled.Int32()) {
		flowCtrl.data = data
		flowCtrl.cancel()
	}
}

// Fail cancels the execution of the flow with error.
func (flowCtrl flowControl) Fail(data interface{}, err error) {
	// Try to change the result from None to Failed and if it's successful
	// finish the flow.
	if atomic.CompareAndSwapInt32(&flowCtrl.result, None.Int32(), Failed.Int32()) {
		flowCtrl.data = data
		// err ???
		flowCtrl.cancel()
	}
}

// IsFinished tests if execution of the flow is either completed or canceled.
func (flowCtrl flowControl) IsFinished() bool {
	r := atomic.LoadInt32(&flowCtrl.result)
	return Result(r).IsFinished()
}

// Result returns the result code and the result data of the flow. The call
// to the function is effective only if the flow is finished.
func (flowCtrl flowControl) Result() (result Result, data interface{}, err error) {
	// Load the current result
	r := atomic.LoadInt32(&flowCtrl.result)
	result = Result(r)

	// Return data only if the flow is finished
	if result.IsFinished() {
		return result, flowCtrl.data, nil
	}

	// Otherwise return nil because reading the data field while the flow is not
	// finished may lead to unpredicted behavior, fot example reading value
	// while other goroutine is writing it.
	return result, nil, nil
}
