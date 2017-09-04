package run

import (
	"fmt"

	floc "github.com/workanator/go-floc.v1"
	"github.com/workanator/go-floc.v1/guard"
	"github.com/workanator/go-floc.v1/pred"
)

func ExampleWhile() {
	const max = 100

	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewState(new(int))
	defer state.Release()

	// The function updates the state with key-value given. In the example key is
	// useless because the state contains only the counter so the function just
	// sets the counter to the value given.
	update := func(flow floc.Flow, state floc.State, key string, value interface{}) {
		// Get data from the state with exclusive lock.
		data, locker := state.DataWithWriteLocker()

		// Lock the data and update it.
		locker.Lock()
		defer locker.Unlock()

		counter := data.(*int)
		*counter = value.(int)
	}

	// The job prints the current value of the counter.
	printResult := func(flow floc.Flow, state floc.State, update floc.Update) {
		// Get data from the state with non-exclusive lock.
		data, locker := state.DataWithReadLocker()

		// Lock the data and print it.
		locker.Lock()
		defer locker.Unlock()

		counter := data.(*int)
		fmt.Println(*counter)
	}

	// The job does nothing.
	nop := func(flow floc.Flow, state floc.State, update floc.Update) {}

	// The predicate tests if the counter reached the limit
	testDone := func(state floc.State) bool {
		// Get the current value of the counter
		data, locker := state.DataWithReadLocker()

		locker.Lock()
		defer locker.Unlock()

		counter := data.(*int)
		current := *counter

		// Test if the limit is reached
		return current == max
	}

	// Counstruct the result job which repeats sequence of jobs 10 times.
	job := Sequence(
		// Increment the counter to max in background and exit
		Background(func(flow floc.Flow, state floc.State, update floc.Update) {
			data, locker := state.DataWithReadLocker()

			for !flow.IsFinished() {
				// Get the current value of the counter
				locker.Lock()
				counter := data.(*int)
				next := *counter + 1
				locker.Unlock()

				// Update the counter and test if it reached the limit
				update(flow, state, "", next)
				if next == max {
					break
				}
			}
		}),
		// Wait until the counter reaches the limit
		While(pred.Not(testDone), nop),
		// Print the result
		printResult,
		// Complete the flow
		guard.Complete(nil),
	)

	// Run the job.
	floc.Run(flow, state, update, job)

	// Output: 100
}
