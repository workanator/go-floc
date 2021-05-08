/*
Package floc allows to orchestrate goroutines with ease. The goal of the project
is to make the process of running goroutines in parallel and synchronizing them
easy.

Floc follows for objectives:

-- Easy to use functional interface.

-- Better control over execution with one entry point and one exit point.

-- Simple parallelism and synchronization of jobs.

-- As little overhead, in comparison to direct use of goroutines and sync
primitives, as possible.

The package categorizes middleware used for architect flows in subpackages.

-- `guard` contains middleware which help protect flow from falling into panic
or unpredicted behavior.

-- `pred` contains some basic predicates for AND, OR, NOT logic.

-- `run` provides middleware for designing flow, i.e. for running job
sequentially, in parallel, in background and so on.

Here is a quick example of what the package capable of.

  // The flow computes something complex and does writing results in
  // background.
  flow := run.Sequence(
    run.Background(WriteToDisk),
    run.While(pred.Not(TestComputed), run.Sequence(
      run.Parallel(
        ComputeSomething,
        ComputeSomethingElse,
        guard.Panic(ComputeDangerousThing),
      ),
      run.Parallel(
        PrepareForWrite,
        UpdateComputedFlag,
      ),
    )),
    CompleteWithSuccess,
  )

  // The entry point: produce the result.
  result, data, err := floc.Run(flow)

  // The exit point: consume the result.
  if err != nil {
    fmt.Println(err)
  } else if result.IsCompleted() {
    fmt.Println(data)
  } else {
    fmt.Printf("Finished with result %s and data %v", result.String(), data)
  }
*/
package floc
