/*
Package floc makes parallel programming easy. The goal of the project is to
make programs more expressive through functional paradigm.

Floc follows for objectives:

-- Split the overall work into the number of small jobs. Floc cannot force you
to do that but doing that grants many advantages starting from simpler testing
and up to better control on execution.

-- Make end algorithms more clear and simpler by expressing them through
the combination of jobs. In short terms floc allows to express job through jobs.

-- Provide better control over execution with one entry point and one exit
point. That is achieved by allowing any job finish execution with Cancel or
Complete.

-- Simple parallelism and synchronization of jobs.

-- As little overhead, in comparison to direct use of goroutines and sync
primitives, as possible.

  job := run.Sequence(
    run.Background(writeToDisk),
    run.While(pred.Not(testComputed), run.Sequence(
      run.Parallel(
        computeSomething,
        computeSomethindElse,
        guard.Panic(computeDangerousThing),
      ),
      run.Parallel(
        prepareForWrite,
        updateComputedFlag,
      ),
    )
    )),
    completeWithSuccess,
  )

  // The entry point: produce the result.
  floc.Run(flow, state, update, job)

  // The exit point: consume the result.
  result, data := flow.Result()
*/
package floc
