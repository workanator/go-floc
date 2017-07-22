package run

import (
	"testing"
	"time"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/guard"
)

func TestRace(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	theJob := Race(
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
	floc.Run(theFlow, state, updateCounter, theJob)

	// Because run.Race allows only one winner the counter must be incremented
	// only once.
	const expect = 1
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter value to be %d but get %d", t.Name(), expect, v)
	}
}

func TestRaceInactive(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()
	theFlow.Complete(nil)

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	theJob := Race(
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
	floc.Run(theFlow, state, updateCounter, theJob)

	if getCounter(state) != 0 {
		t.Fatalf("%s expects counter to be zero", t.Name())
	}
}

func TestRaceInterrupt(t *testing.T) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	theJob := Race(
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
