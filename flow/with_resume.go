package flow

import floc "github.com/workanator/go-floc"

type resumeableFlow struct {
	parent floc.Flow
	fake   floc.Flow
}

// ResumeFunc when invoked resumes the execution of the flow. Effective in
// case the flow was Canceled ot Completed.
type ResumeFunc func() floc.Flow

// WithResume creates a new instance of the flow containing the parent flow
// and a resume function which allows to resume execution of the flow.
func WithResume(parent floc.Flow) (floc.Flow, ResumeFunc) {
	flow := &resumeableFlow{
		parent: parent,
		fake:   New(),
	}

	resume := func() floc.Flow {
		return parent
	}

	return flow, resume
}

// Done returns a channel that's closed when the flow done.
// Successive calls to Done return the same value.
func (f *resumeableFlow) Done() <-chan struct{} {
	return f.fake.Done()
}

// Close finishes the flow and releases all underlying resources.
func (f *resumeableFlow) Close() {
	f.fake.Cancel(nil)
}

// Complete finishes the flow with success status and stops
// execution of futher jobs if any.
func (f *resumeableFlow) Complete(data interface{}) {
	f.fake.Complete(data)
}

// Cancel cancels the execution of the flow.
func (f *resumeableFlow) Cancel(data interface{}) {
	f.fake.Cancel(data)
}

// Tests if the execution of the flow is either completed or canceled.
func (f *resumeableFlow) IsFinished() bool {
	return f.fake.IsFinished()
}

// Returns the result code and the result data of the flow.
func (f *resumeableFlow) Result() (result floc.Result, data interface{}) {
	return f.fake.Result()
}
