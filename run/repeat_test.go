package run

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/guard"
)

func TestRepeat(t *testing.T) {
	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	const times = 10

	job := Repeat(
		times,
		jobIncrement,
	)

	// Run the job.
	floc.Run(flow, state, updateCounter, job)

	expect := times
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}

func TestRepeatInterrupt(t *testing.T) {
	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	const times = 10

	job := Repeat(
		times,
		Sequence(
			jobIncrement,
			guard.Complete(nil),
		),
	)

	// Run the job.
	floc.Run(flow, state, updateCounter, job)

	expect := 1
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}
