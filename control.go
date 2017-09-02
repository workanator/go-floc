package floc

/*
Control allows to control execution of the flow.
*/
type Control interface {
	Releaser

	// Complete finishes the flow with success status and stops
	// execution of further jobs if any.
	Complete(data interface{})

	// Cancel cancels the execution of the flow.
	Cancel(data interface{})

	// Fail cancels the execution of the flow with error.
	Fail(data interface{}, err error)

	// IsFinished tests if execution of the flow is either completed or canceled.
	IsFinished() bool

	// Result returns the result code and the result data of the flow. The call
	// to the function is effective only if the flow is finished.
	Result() (result Result, data interface{}, err error)
}
