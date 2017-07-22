package run

import (
	"fmt"
	"time"

	floc "github.com/workanator/go-floc"
)

// Event is some event protected with sync.Cond.
type Event struct {
	Done chan struct{}
}

func (e *Event) Release() {
	close(e.Done)
}

func ExampleWait() {
	const max = 100000

	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the event.
	state := floc.NewStateContainer(&Event{
		Done: make(chan struct{}),
	})

	// The predicate wait for the event
	waitEvent := func(state floc.State) bool {
		// Get data from the state. Skip the locker because Cond has it's own
		// lock.
		data, _ := state.Get()
		event := data.(*Event)

		// Wait for the condition
		select {
		case <-event.Done:
			return true

		default:
			return false
		}
	}

	// Counstruct the result job.
	job := Sequence(
		// The background job counts to 100000 and sets the condition to true.
		Background(func(flow floc.Flow, state floc.State, update floc.Update) {
			// Get data from the state.
			data, _ := state.Get()
			event := data.(*Event)

			// Notify the waiter the job is done
			defer func() { event.Done <- struct{}{} }()

			// Increase the counter
			counter := 0
			for counter < max && !flow.IsFinished() {
				counter++
			}
		}),
		// Wait until the event
		Wait(waitEvent, 1*time.Microsecond),
		// Print the result
		func(flow floc.Flow, state floc.State, update floc.Update) {
			fmt.Println("Done")
		},
	)

	// Run the job.
	floc.Run(flow, state, nil, job)

	// Output: Done
}
