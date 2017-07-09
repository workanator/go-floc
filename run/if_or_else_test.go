package run

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/guard"
	"github.com/workanator/go-floc/state"
)

func TestIfOrElseTrue(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	theState := state.New(new(int))

	// Counstruct the result job.
	theJob := IfOrElse(predCounterEquals(0), jobIncrement, guard.Cancel(nil))

	// Run the job.
	floc.Run(theFlow, theState, updateCounter, theJob)

	expect := 1
	v := getCounter(theState)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}

func TestIfOrElseFalse(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	theState := state.New(new(int))

	// Counstruct the result job.
	theJob := IfOrElse(predCounterEquals(1), jobIncrement, guard.Cancel(nil))

	// Run the job.
	floc.Run(theFlow, theState, updateCounter, theJob)

	expect := 0
	v := getCounter(theState)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}
