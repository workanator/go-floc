package run

import (
	"testing"
	"time"

	floc "gopkg.in/workanator/go-floc.v1"
	"gopkg.in/workanator/go-floc.v1/guard"
)

func TestRace(t *testing.T) {
	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewState(new(int))
	defer state.Release()

	// Counstruct the result job.
	job := Race(
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
	floc.Run(flow, state, updateCounter, job)

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
	flow := floc.NewFlow()
	defer flow.Release()

	flow.Complete(nil)

	// Construct the state object which as data contains the counter.
	state := floc.NewState(new(int))
	defer state.Release()

	// Counstruct the result job.
	job := Race(
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
	floc.Run(flow, state, updateCounter, job)

	if getCounter(state) != 0 {
		t.Fatalf("%s expects counter to be zero", t.Name())
	}
}

func TestRaceInterrupt(t *testing.T) {
	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewState(new(int))
	defer state.Release()

	// Counstruct the job.
	job := Race(
		Delay(50*time.Millisecond, jobIncrement),
		Delay(5*time.Millisecond, guard.Cancel(nil)),
	)

	// Run the job.
	floc.Run(flow, state, updateCounter, job)

	expect := 0
	v := getCounter(state)
	if v != expect {
		t.Fatalf("%s expects counter to be %d but has %d", t.Name(), expect, v)
	}
}
