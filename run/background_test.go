package run

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
)

func TestBackgroundInactive(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()
	theFlow.Complete(nil)

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	theJob := Background(jobIncrement)

	// Run the job.
	floc.Run(theFlow, state, updateCounter, theJob)

	if getCounter(state) != 0 {
		t.Fatalf("%s expects counter to be zero", t.Name())
	}
}
