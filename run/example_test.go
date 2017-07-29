package run

import (
	"fmt"
	"sync/atomic"

	floc "github.com/workanator/go-floc"
)

func Example_withLocking() {
	// The example shows how to read/write state with locking.

	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewState(new(int))
	defer state.Release()

	// The function updates the state with key-value given. In the example key is
	// useless because the state contains only the counter so the function just
	// increments the counter with the value given.
	update := func(flow floc.Flow, state floc.State, key string, value interface{}) {
		// Get data from the state with exclusive lock.
		data, locker := state.DataWithWriteLocker()
		counter := data.(*int)

		// Lock the data and update it.
		locker.Lock()
		defer locker.Unlock()

		*counter += value.(int)
	}

	// The predicate which tests if the counter value is even.
	isEven := func(state floc.State) bool {
		// Get data from the state with non-exclusive lock.
		data, locker := state.DataWithReadLocker()
		counter := data.(*int)

		// Lock the data and read it.
		locker.Lock()
		defer locker.Unlock()

		return *counter%2 == 0
	}

	// The job prints EVEN.
	printEven := func(flow floc.Flow, state floc.State, update floc.Update) {
		fmt.Print("EVEN ")
	}

	// The job prints the current value of the counter.
	printNumber := func(flow floc.Flow, state floc.State, update floc.Update) {
		// Get data from the state with non-exclusive lock.
		data, locker := state.DataWithReadLocker()
		counter := data.(*int)

		// Lock the data and print it.
		locker.Lock()
		defer locker.Unlock()

		fmt.Println(*counter)
	}

	// The job increments the counter by 1.
	increment := func(flow floc.Flow, state floc.State, update floc.Update) {
		update(flow, state, "", 1)
	}

	// Counstruct the result job which repeats sequence of jobs 10 times.
	job := Repeat(10, Sequence(
		If(isEven, printEven), // Print EVEN if the value of the counter is even
		printNumber,           // Print the value of the counter
		increment,             // Increment the counter
	))

	// Run the job.
	floc.Run(flow, state, update, job)

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

func Example_withAtomic() {
	// The example demonstrates how to read/write state with atomic operations,
	// i.e. without locking.

	// Construct the flow control object.
	flow := floc.NewFlow()
	defer flow.Release()

	// Construct the state object which as data contains the counter.
	state := floc.NewState(new(int32))
	defer state.Release()

	// The function updates the state with key-value given. In the example key is
	// useless because the state contains only the counter so the function just
	// increments the counter with the value given.
	update := func(flow floc.Flow, state floc.State, key string, value interface{}) {
		// Get data from the state.
		counter := state.Data().(*int32)

		atomic.AddInt32(counter, int32(value.(int)))
	}

	// The predicate which tests if the counter value is even.
	isEven := func(state floc.State) bool {
		// Get data from the state.
		counter := state.Data().(*int32)

		return atomic.LoadInt32(counter)%2 == 0
	}

	// The job prints EVEN.
	printEven := func(flow floc.Flow, state floc.State, update floc.Update) {
		fmt.Print("EVEN ")
	}

	// The job prints the current value of the counter.
	printNumber := func(flow floc.Flow, state floc.State, update floc.Update) {
		// Get data from the state.
		counter := state.Data().(*int32)

		fmt.Println(atomic.LoadInt32(counter))
	}

	// The job increments the counter by 1.
	increment := func(flow floc.Flow, state floc.State, update floc.Update) {
		update(flow, state, "", 1)
	}

	// Counstruct the result job which repeats sequence of jobs 10 times.
	job := Repeat(10, Sequence(
		If(isEven, printEven), // Print EVEN if the value of the counter is even
		printNumber,           // Print the value of the counter
		increment,             // Increment the counter
	))

	// Run the job.
	floc.Run(flow, state, update, job)

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
