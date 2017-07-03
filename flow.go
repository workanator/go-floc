package floc

// Flow provides function which allow to control flow
type Flow interface {
	// Done returns a channel that's closed when the flow done.
	// Successive calls to Done return the same value.
	Done() <-chan struct{}

	// Close finishes the flow and releases all underlying resources.
	Close()

	// Complete finishes the flow with success status and stops
	// execution of futher nodes if any.
	Complete(data interface{})

	// Cancel cancels the execution of the flow.
	Cancel(data interface{})

	// Tests if the execution of the flow is either completed or canceled.
	IsFinished() bool

	// Returns the result code and the result data of the flow.
	Result() (result Result, data interface{})
}
