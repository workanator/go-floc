/*
Package run is the collection of jobs which make the architecture of the flow.

Each function in the package is middleware which always takes at leats one
floc.Job to run and constructs and returns another floc.Job. That allows to
organize jobs in any combination and in result is only one floc.Job which can
be run with floc.Run().

  func ExampleJob() {
    // Construct the flow control object.
    theFlow := flow.New()

    // Construct the state object which as data contains the counter.
    theState := state.New(new(int))

    // The function updates the state with key-value given. In the example key is
    // useless because the state contains only the counter so the function just
    // increments the counter with the value given.
    theUpdate := func(flow floc.Flow, state floc.State, key string, value interface{}) {
      data, lock := state.GetExclusive()
      counter := data.(*int)

      lock.Lock()
      defer lock.Unlock()

      *counter += value.(int)
    }

    // The predicate which tests if the counter value is even.
    isEven := func(flow floc.Flow, state floc.State) bool {
      data, lock := state.Get()
      counter := data.(*int)

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
      data, lock := state.Get()
      counter := data.(*int)

      lock.Lock()
      defer lock.Unlock()

      fmt.Println(*counter)
    }

    // The job increments the counter by 1.
    increment := func(flow floc.Flow, state floc.State, update floc.Update) {
      update(flow, state, "", 1)
    }

    // Counstruct the result job which repeats sequence of jobs 10 times.
    theJob := run.Repeat(10, run.Sequence(
      run.If(isEven, printEven),
      printNumber,
      increment,
    ))

    // Run the job.
    floc.Run(theFlow, theState, theUpdate, theJob)
  }
*/
package run
