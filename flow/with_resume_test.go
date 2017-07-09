package flow

import (
	"testing"

	floc "github.com/workanator/go-floc"
)

func TestWithResume(t *testing.T) {
	resumeFlow, resumeFunc := WithResume(New())

	// Complete resumeFlow
	resumeFlow.Complete(nil)
	result, _ := resumeFlow.Result()

	if result != floc.Completed {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed, result)
	}

	// Resume the flow
	flow := resumeFunc()
	result, _ = flow.Result()

	if result != floc.None {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.None, result)
	}
}
