package guard

import (
	"testing"
	"time"

	floc "github.com/workanator/go-floc.v1"
	"github.com/workanator/go-floc.v1/run"
)

func TestDeadline(t *testing.T) {
	const ID int = 1

	f := floc.NewFlow()
	s := floc.NewState(nil)

	// Make deadline 100 milliseconds in the future and with the job which
	// should finish prioir the dealine
	job := Deadline(DeadlineIn(100*time.Millisecond), ID, Complete(nil))

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed.String(), result)
	}
}

func TestDeadlineWithTrigger(t *testing.T) {
	const ID int = 2

	f := floc.NewFlow()
	s := floc.NewState(nil)

	// Make deadline 50 milliseconds in the future and with the job which should
	// run with the delay in 200 milliseconds so the trigger should be invoked
	job := DeadlineWithTrigger(
		DeadlineIn(50*time.Millisecond),
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
