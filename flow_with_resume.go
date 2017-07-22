package floc

type resumeableFlowControl struct {
	Flow
	parent Flow
}

// ResumeFunc when invoked resumes the execution of the flow. Effective in
// case the flow was Canceled or Completed.
type ResumeFunc func() Flow

// NewFlowWithResume creates a new instance of the flow, containing
// the parent flow, and a resume function which allows to resume execution
// of the flow.
func NewFlowWithResume(parent Flow) (Flow, ResumeFunc) {
	flow := &resumeableFlowControl{
		Flow:   NewFlow(),
		parent: parent,
	}

	resume := func() Flow {
		return parent
	}

	return flow, resume
}

// Release finishes the flow and releases all underlying resources.
func (flow *resumeableFlowControl) Release() {
	flow.Flow.Release()
	flow.parent.Release()
}
