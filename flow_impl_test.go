package floc

import (
	"testing"
)

func TestFlowEmpty(t *testing.T) {
	flow := NewFlow()
	defer flow.Release()

	result, data := flow.Result()

	if !result.IsNone() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), None.String(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	}
}

func TestFlowComplete(t *testing.T) {
	const value = "complete"

	flow := NewFlow()
	defer flow.Release()

	flow.Complete(value)
	result, data := flow.Result()

	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), Completed.String(), result.String())
	} else if data == nil {
		t.Fatalf("%s expects data to be non-nil", t.Name())
	}

	s := data.(string)
	if s != value {
		t.Fatalf("%s expects data to be string %s but has %v", t.Name(), value, data)
	}
}

func TestFlowCancel(t *testing.T) {
	const value = "cancel"

	flow := NewFlow()
	defer flow.Release()

	flow.Cancel(value)
	result, data := flow.Result()

	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), Canceled.String(), result.String())
	} else if data == nil {
		t.Fatalf("%s expects data to be non-nil", t.Name())
	}

	s := data.(string)
	if s != value {
		t.Fatalf("%s expects data to be string %s but has %v", t.Name(), value, data)
	}
}

func TestFlowClose(t *testing.T) {
	flow := NewFlow()
	flow.Release()

	select {
	case <-flow.Done():
		result, _ := flow.Result()

		if !result.IsCanceled() {
			t.Fatalf("%s expects result to be %s but has %s", t.Name(), Canceled.String(), result.String())
		}

	default:
		t.Failed()
	}
}

func TestFlowIsFinished(t *testing.T) {
	flow := NewFlow()
	defer flow.Release()

	if flow.IsFinished() {
		t.Fatalf("%s must not be finished", t.Name())
	}

	flow.Complete(true)

	if !flow.IsFinished() {
		t.Fatalf("%s must be finished", t.Name())
	}
}
