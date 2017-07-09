package run

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/guard"
	"github.com/workanator/go-floc/state"
)

func TestLoop(t *testing.T) {
	const max = 10

	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	theState := state.New(new(int))

	// Counstruct the result job.
	theJob := Loop(
		jobIncrement,
		If(predCounterEquals(max), guard.Complete(nil)),
	)

	// Run the job.
	floc.Run(theFlow, theState, updateCounter, theJob)

	expect := max
	v := getCounter(theState)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}
