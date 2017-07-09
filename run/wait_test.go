package run

import (
	"fmt"
	"sync"
	"time"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/state"
)

// Event is some event protected with sync.Cond.
type Event struct {
	Cond *sync.Cond
	Done bool
}

func ExampleWait() {
	const max = 100000

	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	theState := state.New(&Event{
		Cond: sync.NewCond(&sync.Mutex{}),
	})

	// The predicate wait for the event
	waitEvent := func(flow floc.Flow, state floc.State) bool {
		// Get data from the state. Skip the locker because Cond has it's own
		// lock.
		data, _ := state.Get()
		event := data.(*Event)

		// Wait for the condition
		event.Cond.L.Lock()
		event.Cond.Wait()
		done := event.Done
		event.Cond.L.Unlock()

		return done
	}

	// Counstruct the result job.
	theJob := Sequence(
		// The background job counts to 100000 and sets the condition to true.
		Background(func(flow floc.Flow, state floc.State, update floc.Update) {
			// Get data from the state. Skip the locker because Cond has it's own
			// lock.
			data, _ := state.Get()
			event := data.(*Event)

			// Increase the counter
			counter := 0
			for counter < max && !flow.IsFinished() {
				counter++

				// Wake up the waiting job evey 1000th iteration.
				if counter%1000 == 0 {
					event.Cond.Signal()
				}
			}

			// Set the event to true
			event.Cond.L.Lock()
			event.Done = true
			event.Cond.L.Unlock()

			event.Cond.Signal()
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
