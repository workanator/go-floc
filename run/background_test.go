package run

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/state"
)

func TestBackgroundInactive(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()
	theFlow.Complete(nil)

	// Construct the state object which as data contains the counter.
	theState := state.New(new(int))

	// Counstruct the result job.
	theJob := Background(jobIncrement)

	// Run the job.
	floc.Run(theFlow, theState, updateCounter, theJob)

	if getCounter(theState) != 0 {
		t.Fatalf("%s expects counter to be zero", t.Name())
	}
}
