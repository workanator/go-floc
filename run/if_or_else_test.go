package run

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/guard"
)

func TestIfOrElseTrue(t *testing.T) {
	// Construct the flow control object.
	flow := floc.NewFlowControl()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	job := IfOrElse(predCounterEquals(0), jobIncrement, guard.Cancel(nil))

	// Run the job.
	floc.Run(flow, state, updateCounter, job)

	expect := 1
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}

func TestIfOrElseFalse(t *testing.T) {
	// Construct the flow control object.
	flow := floc.NewFlowControl()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	job := IfOrElse(predCounterEquals(1), jobIncrement, guard.Cancel(nil))

	// Run the job.
	floc.Run(flow, state, updateCounter, job)

	expect := 0
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}
