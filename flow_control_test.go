package floc

import (
	"testing"
)

func TestFlowControlEmpty(t *testing.T) {
	flow := NewFlowControl()
	defer flow.Release()

	result, data := flow.Result()

	if !result.IsNone() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), None.String(), result)
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	}
}

func TestFlowControlComplete(t *testing.T) {
	const value = "complete"

	flow := NewFlowControl()
	defer flow.Release()

	flow.Complete(value)
	result, data := flow.Result()

	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), Completed.String(), result)
	} else if data == nil {
		t.Fatalf("%s expects data to be non-nil", t.Name())
	}

	s := data.(string)
	if s != value {
		t.Fatalf("%s expects data to be string %s but has %v", t.Name(), value, data)
	}
}

func TestFlowControlCancel(t *testing.T) {
	const value = "cancel"

	flow := NewFlowControl()
	defer flow.Release()

	flow.Cancel(value)
	result, data := flow.Result()

	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), Canceled.String(), result)
	} else if data == nil {
		t.Fatalf("%s expects data to be non-nil", t.Name())
	}

	s := data.(string)
	if s != value {
		t.Fatalf("%s expects data to be string %s but has %v", t.Name(), value, data)
	}
}

func TestFlowControlClose(t *testing.T) {
	flow := NewFlowControl()
	flow.Release()

	select {
	case <-flow.Done():
		result, _ := flow.Result()

		if !result.IsCanceled() {
			t.Fatalf("%s expects result to be %s but has %s", t.Name(), Canceled.String(), result)
		}

	default:
		t.Failed()
	}
}

func TestFlowControlIsFinished(t *testing.T) {
	flow := NewFlowControl()
	defer flow.Release()

	if flow.IsFinished() {
		t.Fatalf("%s must not be finished", t.Name())
	}

	flow.Complete(true)

	if !flow.IsFinished() {
		t.Fatalf("%s must be finished", t.Name())
	}
}
