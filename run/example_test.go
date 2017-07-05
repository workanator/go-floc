package run

import (
	"fmt"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/state"
)

func Example() {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	theState := state.New(new(int))

	// The function updates the state with key-value given. In the example key is
	// useless because the state contains only the counter so the function just
	// increments the counter with the value given.
	theUpdate := func(flow floc.Flow, state floc.State, key string, value interface{}) {
		// Get data from the state with exclusive lock.
		data, lock := state.GetExclusive()
		counter := data.(*int)

		// Lock the data and update it.
		lock.Lock()
		defer lock.Unlock()

		*counter += value.(int)
	}

	// The predicate which tests if the counter value is even.
	isEven := func(flow floc.Flow, state floc.State) bool {
		// Get data from the state with non-exclusive lock.
		data, lock := state.Get()
		counter := data.(*int)

		// Lock the data and read it.
		lock.Lock()
		defer lock.Unlock()

		return *counter%2 == 0
	}

	// The job prints EVEN.
	printEven := func(flow floc.Flow, state floc.State, update floc.Update) {
		fmt.Print("EVEN ")
	}

	// The job prints the current value of the counter.
	printNumber := func(flow floc.Flow, state floc.State, update floc.Update) {
		// Get data from the state with non-exclusive lock.
		data, lock := state.Get()
		counter := data.(*int)

		// Lock the data and print it.
		lock.Lock()
		defer lock.Unlock()

		fmt.Println(*counter)
	}

	// The job increments the counter by 1.
	increment := func(flow floc.Flow, state floc.State, update floc.Update) {
		update(flow, state, "", 1)
	}

	// Counstruct the result job which repeats sequence of jobs 10 times.
	theJob := Repeat(10, Sequence(
		If(isEven, printEven), // Print EVEN if the value of the counter is even
		printNumber,           // Print the value of teh counter
		increment,             // Increment the counter
	))

	// Run the job.
	floc.Run(theFlow, theState, theUpdate, theJob)

	// Output:
	// EVEN 0
	// 1
	// EVEN 2
	// 3
	// EVEN 4
	// 5
	// EVEN 6
	// 7
	// EVEN 8
	// 9
}
