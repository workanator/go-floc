package run

import (
	"testing"

	floc "gopkg.in/workanator/go-floc.v1"
	"gopkg.in/workanator/go-floc.v1/guard"
)

func TestLoop(t *testing.T) {
	const max = 10

	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewState(new(int))
	defer state.Release()

	// Counstruct the result job.
	job := Loop(
		jobIncrement,
		If(predCounterEquals(max), guard.Complete(nil)),
	)

	// Run the job.
	floc.Run(flow, state, updateCounter, job)

	expect := max
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}
