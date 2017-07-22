package run

import (
	"testing"
	"time"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/guard"
)

func TestDelay(t *testing.T) {
	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewState(new(int))
	defer state.Release()

	// Counstruct the result job.
	job := Delay(
		1*time.Nanosecond,
		jobIncrement,
		jobIncrement,
		jobIncrement,
		jobIncrement,
		jobIncrement,
	)

	// Run the job.
	floc.Run(flow, state, updateCounter, job)

	expect := 5
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}

func TestDelayInactive(t *testing.T) {
	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewState(new(int))
	defer state.Release()

	// Counstruct the result job.
	job := Delay(
		1*time.Nanosecond,
		jobIncrement,
		jobIncrement,
		guard.Cancel(nil),
		jobIncrement,
		jobIncrement,
	)

	// Run the job.
	floc.Run(flow, state, updateCounter, job)

	expect := 2
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}

func TestDelayInterrupt(t *testing.T) {
	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewState(new(int))
	defer state.Release()

	// Counstruct the result job.
	job := Parallel(
		Delay(50*time.Millisecond, jobIncrement),
		Delay(5*time.Millisecond, guard.Cancel(nil)),
	)

	// Run the job.
	floc.Run(flow, state, updateCounter, job)

	expect := 0
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}
