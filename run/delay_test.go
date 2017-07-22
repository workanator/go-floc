package run

import (
	"testing"
	"time"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/guard"
)

func TestDelay(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

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
	floc.Run(theFlow, state, updateCounter, theJob)

	expect := 5
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}

func TestDelayInactive(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

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
	floc.Run(theFlow, state, updateCounter, theJob)

	expect := 2
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}

func TestDelayInterrupt(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	theJob := Parallel(
		Delay(50*time.Millisecond, jobIncrement),
		Delay(5*time.Millisecond, guard.Cancel(nil)),
	)

	// Run the job.
	floc.Run(theFlow, state, updateCounter, theJob)

	expect := 0
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}
