package floc

// Flow provides the control over execution of the flow.
type Flow interface {
	Releaser

	// Done returns a channel that's closed when the flow done.
	// Successive calls to Done return the same value.
	Done() <-chan struct{}

	// Complete finishes the flow with success status and stops
	// execution of further jobs if any.
	Complete(data interface{})

	// Cancel cancels the execution of the flow.
	Cancel(data interface{})

	// IsFinished tests if execution of the flow is either completed or canceled.
	IsFinished() bool

	// Result returns the result code and the result data of the flow. The call
	// to the function is effective only if the flow is finished.
	Result() (result Result, data interface{})
}
