package run

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/guard"
)

func TestRepeat(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	const times = 10
	theJob := Repeat(
		times,
		jobIncrement,
	)

	// Run the job.
	floc.Run(theFlow, state, updateCounter, theJob)

	expect := times
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}

func TestRepeatInterrupt(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	const times = 10
	theJob := Repeat(
		times,
		Sequence(
			jobIncrement,
			guard.Complete(nil),
		),
	)

	// Run the job.
	floc.Run(theFlow, state, updateCounter, theJob)

	expect := 1
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}
