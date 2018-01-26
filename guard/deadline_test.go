package guard

import (
	"testing"
	"time"

	"gopkg.in/devishot/go-floc.v2"
	"gopkg.in/devishot/go-floc.v2/run"
)

func TestDeadline(t *testing.T) {
	const ID int = 1

	flow := Deadline(DeadlineIn(100*time.Millisecond), ID, Complete(nil))

	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestOnDeadline(t *testing.T) {
	const ID int = 2

	flow := OnDeadline(
		DeadlineIn(50*time.Millisecond),
		ID,
		run.Delay(200*time.Millisecond, Complete(nil)),
		func(ctx floc.Context, ctrl floc.Control, id interface{}) {
			if deadlineID, ok := id.(int); !ok {
				t.Fatalf("%s expects data to be of type int but has %T", t.Name(), deadlineID)
			} else if id != ID {
				t.Fatalf("%s expects ID to be %d but has %d", t.Name(), ID, id)
			}

			ctrl.Cancel(nil)
		},
	)

	result, data, err := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result)
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}
