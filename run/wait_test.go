package run

import (
	"fmt"
	"time"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/state"
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
	theFlow := flow.New()

	// Construct the state object which as data contains the event.
	theState := state.New(&Event{
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
	theJob := Sequence(
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
	floc.Run(theFlow, theState, nil, theJob)

	// Output: Done
}
