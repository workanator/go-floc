/*
Package pred provides predicates for basic logics.

Predicates with conditional jobs like run.If allows to make non-linear
algorithms. In terms of floc predicate should return true or false depending
on state.

  testReady := func(flow floc.FLow, state floc.State) bool {
    data, lock := state.Get()
    env := data.(*MyEnv)

    lock.Lock()
    defer lock.Unlock()

    return env.SomethingIsReady
  }

  job := run.Sequence(
    ..., // Some job done here
    job.If(testReady, job.Background(writeToDisk)), // Write some data ready to disk in background
    ..., // Some job more
  )
*/
package pred
