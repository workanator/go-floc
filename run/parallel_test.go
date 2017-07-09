package run

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/state"
)

func TestParallel(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	theState := state.New(new(int))

	// Counstruct the result job.
	theJob := Parallel(
		jobIncrement, // 1
		jobIncrement, // 2
		jobIncrement, // 3
		jobIncrement, // 4
		jobIncrement, // 5
		jobIncrement, // 6
		jobIncrement, // 7
		jobIncrement, // 8
		jobIncrement, // 9
		jobIncrement, // 10
	)

	// Run the job.
	floc.Run(theFlow, theState, updateCounter, theJob)

	const expect = 10
	v := getCounter(theState)
	if v != expect {
		t.Fatalf("%s expects counter value to be %d but get %d", t.Name(), expect, v)
	}
}

func TestParallelInactive(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()
	theFlow.Complete(nil)

	// Construct the state object which as data contains the counter.
	theState := state.New(new(int))

	// Counstruct the result job.
	theJob := Parallel(
		jobIncrement, // 1
		jobIncrement, // 2
		jobIncrement, // 3
		jobIncrement, // 4
		jobIncrement, // 5
		jobIncrement, // 6
		jobIncrement, // 7
		jobIncrement, // 8
		jobIncrement, // 9
		jobIncrement, // 10
	)

	// Run the job.
	floc.Run(theFlow, theState, updateCounter, theJob)

	if getCounter(theState) != 0 {
		t.Fatalf("%s expects counter to be zero", t.Name())
	}
}
