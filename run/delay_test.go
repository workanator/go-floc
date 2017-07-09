package run

import (
	"testing"
	"time"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/guard"
	"github.com/workanator/go-floc/state"
)

func TestDelay(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	theState := state.New(new(int))

	// Counstruct the result job.
	theJob := Delay(
		1*time.Nanosecond,
		jobIncrement,
		jobIncrement,
		jobIncrement,
		jobIncrement,
		jobIncrement,
	)

	// Run the job.
	floc.Run(theFlow, theState, updateCounter, theJob)

	expect := 5
	v := getCounter(theState)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}

func TestDelayInterrupt(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	theState := state.New(new(int))

	// Counstruct the result job.
	theJob := Delay(
		1*time.Nanosecond,
		jobIncrement,
		jobIncrement,
		guard.Cancel(nil),
		jobIncrement,
		jobIncrement,
	)

	// Run the job.
	floc.Run(theFlow, theState, updateCounter, theJob)

	expect := 2
	v := getCounter(theState)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}
