package guard

import (
	"testing"
	"time"

	"github.com/workanator/go-floc/v3"
	"github.com/workanator/go-floc/v3/run"
)

func TestTimeout(t *testing.T) {
	const ID int = 1

	flow := run.Sequence(
		Timeout(ConstTimeout(time.Second), ID, func(floc.Context, floc.Control) error { return nil }),
		Complete(nil),
	)

	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result)
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestTimeout2(t *testing.T) {
	const ID int = 2

	flow := Timeout(ConstTimeout(50*time.Millisecond), ID,
		run.Delay(200*time.Millisecond, Complete(nil)),
	)

	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result)
	} else if data == nil {
		t.Fatalf("%s expects data to be not nil", t.Name())
	} else if d, ok := data.(int); !ok {
		t.Fatalf("%s expects data to be of type int but has %T", t.Name(), data)
	} else if d != ID {
		t.Fatalf("%s expects data to be %d but has %d", t.Name(), ID, d)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	} else if e, ok := err.(floc.ErrTimeout); !ok {
		t.Fatalf("%s expects error to be of type ErrTimeout but has %T", t.Name(), err)
	} else if e.ID() == nil {
		t.Fatalf("%s expects error ID to be not nil", t.Name())
	} else if id, ok := e.ID().(int); !ok {
		t.Fatalf("%s expects error ID to be of type int but has %T", t.Name(), e.ID())
	} else if id != ID {
		t.Fatalf("%s expects error ID to be %d but has %d", t.Name(), ID, id)
	}
}

func TestOnTimeout(t *testing.T) {
	const ID int = 3

	flow := OnTimeout(
		ConstTimeout(50*time.Millisecond),
		ID,
		run.Delay(200*time.Millisecond, Complete(nil)),
		func(ctx floc.Context, ctrl floc.Control, id interface{}) {
			if timeoutID, ok := id.(int); !ok {
				t.Fatalf("%s expects data to be of type int but has %T", t.Name(), timeoutID)
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
