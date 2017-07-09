package flow

import floc "github.com/workanator/go-floc"

type resumeableFlow struct {
	floc.Flow
	parent floc.Flow
}

// ResumeFunc when invoked resumes the execution of the flow. Effective in
// case the flow was Canceled or Completed.
type ResumeFunc func() floc.Flow

// WithResume creates a new instance of the flow, containing the parent flow,
// and a resume function which allows to resume execution of the flow.
func WithResume(parent floc.Flow) (floc.Flow, ResumeFunc) {
	flow := &resumeableFlow{
		Flow:   New(),
		parent: parent,
	}

	resume := func() floc.Flow {
		return parent
	}

	return flow, resume
}
