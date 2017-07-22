package floc

import (
	"testing"
)

func TestFlowWithResume(t *testing.T) {
	resumeFlow, resumeFunc := NewFlowWithResume(NewFlow())
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
