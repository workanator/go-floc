package guard

import (
	"testing"
	"time"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/run"
	"github.com/workanator/go-floc/state"
)

func TestTimeoutZero(t *testing.T) {
	const ID = 1

	f := flow.New()
	s := state.New(nil)

	// Make zero timeout
	job := Timeout(0, ID, Complete(nil))

	floc.Run(f, s, nil, job)

	result, data := f.Result()
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled, result)
	}

	e, ok := data.(ErrTimeout)
	if !ok {
		t.Fatalf("%s expects data to be ErrTimeout but has %T", t.Name(), data)
	}

	if e.ID != ID {
		t.Fatalf("%s expects ID to be %d but has %d", t.Name(), ID, e.ID)
	}
}

func TestTimeoutNegativeWithTrigger(t *testing.T) {
	const ID int = 2

	f := flow.New()
	s := state.New(nil)

	// Make negative timeout with trigger which must be invoked
	job := TimeoutWithTrigger(
		-1*time.Second,
		ID,
		Complete(nil),
		func(flow floc.Flow, state floc.State, id interface{}) {
			ident, ok := id.(int)
			if !ok {
				t.Fatalf("%s expects data to be int but has %T", t.Name(), ident)
			}

			if id != ID {
				t.Fatalf("%s expects ID to be %d but has %d", t.Name(), ID, id)
			}

			flow.Cancel(nil)
		},
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled, result)
	}
}

func TestTimeout(t *testing.T) {
	const ID int = 3

	f := flow.New()
	s := state.New(nil)

	// Make timeout in 1 seconds with the job which should finish prior
	// the timeout
	job := run.Sequence(
		Timeout(time.Second, ID, func(floc.Flow, floc.State, floc.Update) {}),
		Complete(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed, result)
	}
}

func TestTimeoutWithDefaultBehavior(t *testing.T) {
	const ID int = 4

	f := flow.New()
	s := state.New(nil)

	// Make timeout in 50 milliseconds while job start is delayed by
	// 200 milliseconds so the timeout should fire first
	job := Timeout(50*time.Millisecond, ID,
		run.Delay(200*time.Millisecond, Complete(nil)),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled, result)
	}
}

func TestTimeoutWithTrigger(t *testing.T) {
	const ID int = 5

	f := flow.New()
	s := state.New(nil)

	// Make deadline 50 milliseconds in the future and with the job which should
	// run with the delay in 200 milliseconds so the trigger should be invoked
	job := TimeoutWithTrigger(
		50*time.Millisecond,
		ID,
		run.Delay(200*time.Millisecond, Complete(nil)),
		func(flow floc.Flow, state floc.State, id interface{}) {
			ident, ok := id.(int)
			if !ok {
				t.Fatalf("%s expects data to be int but has %T", t.Name(), ident)
			}

			if id != ID {
				t.Fatalf("%s expects ID to be %d but has %d", t.Name(), ID, id)
			}

			flow.Cancel(nil)
		},
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled, result)
	}
}
