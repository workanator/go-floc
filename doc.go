/*
Package floc allows to orchestrate goroutines with ease. The goal of the project
is to make the process of running goroutines in parallel and synchronizing them
easy.

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

The package categorizes middleware used for flow building in subpackages.

-- `guard` contains middleware which help protect flow from falling into panic
or unpredicted behavior.

-- `pred` contains some basic predicates for AND, OR, NOT logics.

-- `run` provides middleware for designing flow, i.e. for running job
sequentially, in parallel, in background and so on.

Here is a quick example of what the package capable of.

  // The job computes something complex and does writing of results in
  // background.
  job := run.Sequence(
    run.Background(WriteToDisk),
    run.While(pred.Not(TestComputed), run.Sequence(
      run.Parallel(
        ComputeSomething,
        ComputeSomethindElse,
        guard.Panic(ComputeDangerousThing),
      ),
      run.Parallel(
        PrepareForWrite,
        UpdateComputedFlag,
      ),
    )
    )),
    CompleteWithSuccess,
  )

  // The entry point: produce the result.
  floc.Run(flow, state, update, job)

  // The exit point: consume the result.
  result, data := flow.Result()
*/
package floc
