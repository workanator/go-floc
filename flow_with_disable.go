package floc

import "sync/atomic"

const (
	controlEnabled  = 1 // Call to Complete or Cancel is permitted
	controlDisabled = 0 // Call to Complete or Cancel is prohibited
)

type disablableFlowControl struct {
	parent Flow
	state  int32
}

// DisableFunc when invoked disables calls to Complete and Cancel.
type DisableFunc func()

// NewFlowWithDisable creates a new instance of the flow, containing
// the parent flow, and a disable function which allows to disable calls
// to Complete and Cancel.
func NewFlowWithDisable(parent Flow) (Flow, DisableFunc) {
	flow := &disablableFlowControl{
		parent: parent,
		state:  controlEnabled,
	}

	disable := func() {
		atomic.StoreInt32(&flow.state, controlDisabled)
	}

	return flow, disable
}

// Release finishes the flow and releases all underlying resources.
func (flow *disablableFlowControl) Release() {
	// Propagate to the parent flow
	flow.parent.Cancel(nil)
}

// Done returns a channel that's closed when the flow done.
// Successive calls to Done return the same value.
func (flow *disablableFlowControl) Done() <-chan struct{} {
	// Propagate to the parent flow
	return flow.parent.Done()
}

// Complete finishes the flow with success status and stops
// execution of further jobs if any.
func (flow *disablableFlowControl) Complete(data interface{}) {
	// Propagate to the parent flow unless disabled
	if atomic.LoadInt32(&flow.state) == controlEnabled {
		flow.parent.Complete(data)
	}
}

// Cancel cancels execution of the flow.
func (flow *disablableFlowControl) Cancel(data interface{}) {
	// Propagate to the parent flow unless disabled
	if atomic.LoadInt32(&flow.state) == controlEnabled {
		flow.parent.Cancel(data)
	}
}

// Tests if execution of the flow is either completed or canceled.
func (flow *disablableFlowControl) IsFinished() bool {
	// Propagate to the parent flow
	return flow.parent.IsFinished()
}

// Result returns the result code and the result data of the flow.
func (flow *disablableFlowControl) Result() (result Result, data interface{}) {
	// Propagate to the parent flow
	return flow.parent.Result()
}
