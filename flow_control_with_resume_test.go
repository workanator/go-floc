package floc

import (
	"testing"
)

func TestFlowControlWithResume(t *testing.T) {
	resumeFlow, resumeFunc := NewFlowControlWithResume(NewFlowControl())
	defer resumeFlow.Release()

	// Complete resumeFlow
	resumeFlow.Complete(nil)
	result, _ := resumeFlow.Result()

	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), Completed.String(), result)
	}

	// Resume the flow
	flow := resumeFunc()
	result, _ = flow.Result()

	if !result.IsNone() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), None.String(), result)
	}
}
