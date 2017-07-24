package guard

import (
	"testing"
	"time"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/run"
)

func TestTimeout(t *testing.T) {
	const ID int = 1

	f := floc.NewFlow()
	s := floc.NewState(nil)

	// Make timeout in 1 seconds with the job which should finish prior
	// the timeout
	job := run.Sequence(
		Timeout(ConstTimeout(time.Second), ID, func(floc.Flow, floc.State, floc.Update) {}),
		Complete(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed.String(), result)
	}
}

func TestTimeoutWithDefaultBehavior(t *testing.T) {
	const ID int = 2

	f := floc.NewFlow()
	s := floc.NewState(nil)

	// Make timeout in 50 milliseconds while job start is delayed by
	// 200 milliseconds so the timeout should fire first
	job := Timeout(ConstTimeout(50*time.Millisecond), ID,
		run.Delay(200*time.Millisecond, Complete(nil)),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled.String(), result)
	}
}

func TestTimeoutWithTrigger(t *testing.T) {
	const ID int = 3

	f := floc.NewFlow()
	s := floc.NewState(nil)

	// Make deadline 50 milliseconds in the future and with the job which should
	// run with the delay in 200 milliseconds so the trigger should be invoked
	job := TimeoutWithTrigger(
		ConstTimeout(50*time.Millisecond),
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
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled.String(), result)
	}
}
