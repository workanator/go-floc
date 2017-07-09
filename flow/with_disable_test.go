package flow

import (
	"testing"

	floc "github.com/workanator/go-floc"
)

func TestWithDisableComplete(t *testing.T) {
	const value = "complete"

	flow, _ := WithDisable(New())
	flow.Complete(value)
	result, data := flow.Result()

	if result != floc.Completed {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed, result)
	} else if data == nil {
		t.Fatalf("%s expects data to be non-nil", t.Name())
	}

	s := data.(string)
	if s != value {
		t.Fatalf("%s expects data to be string %s but has %v", t.Name(), value, data)
	}
}

func TestWithDisableCancel(t *testing.T) {
	const value = "cancel"

	flow, _ := WithDisable(New())
	flow.Cancel(value)
	result, data := flow.Result()

	if result != floc.Canceled {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled, result)
	} else if data == nil {
		t.Fatalf("%s expects data to be non-nil", t.Name())
	}

	s := data.(string)
	if s != value {
		t.Fatalf("%s expects data to be string %s but has %v", t.Name(), value, data)
	}
}

func TestWithDisableClose(t *testing.T) {
	flow, _ := WithDisable(New())
	flow.Close()

	select {
	case <-flow.Done():
		result, _ := flow.Result()

		if result != floc.Canceled {
			t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled, result)
		}

	default:
		t.Failed()
	}
}

func TestWithDisableIsFinished(t *testing.T) {
	flow, _ := WithDisable(New())

	if flow.IsFinished() {
		t.Fatalf("%s must not be finished", t.Name())
	}

	flow.Complete(true)

	if !flow.IsFinished() {
		t.Fatalf("%s must be finished", t.Name())
	}
}

func TestDisableFuncComplete(t *testing.T) {
	flow, disableFunc := WithDisable(New())
	disableFunc()

	// Complete must not trigger because it's disabled
	flow.Complete(nil)
	result, data := flow.Result()

	if result != floc.None {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.None, result)
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	}
}

func TestDisableFuncCancel(t *testing.T) {
	flow, disableFunc := WithDisable(New())
	disableFunc()

	// Cancel must not trigger because it's disabled
	flow.Cancel(nil)
	result, data := flow.Result()

	if result != floc.None {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.None, result)
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	}
}
