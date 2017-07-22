package run

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
)

func TestIfTrue(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	theJob := If(predCounterEquals(0), jobIncrement)

	// Run the job.
	floc.Run(theFlow, state, updateCounter, theJob)

	expect := 1
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}

func TestIfFalse(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	theJob := If(predCounterEquals(1), jobIncrement)

	// Run the job.
	floc.Run(theFlow, state, updateCounter, theJob)

	expect := 0
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}
