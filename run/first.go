package run

import (
  "gopkg.in/workanator/go-floc.v2"
  "gopkg.in/workanator/go-floc.v2/guard"
)

/*
First runs jobs in their own goroutines and waits until first of them finish.

Summary:
	- Run jobs in goroutines : YES
	- Wait all jobs finish   : NO
	- Run order              : PARALLEL

Diagram:
    +-->[JOB_1]--+
    |            |
  --+-->  ..     +-->
    |            |
    +-->[JOB_N]--+
*/
func First(jobs ...floc.Job) floc.Job {
  return func(ctx floc.Context, ctrl floc.Control) error {
    // Do not start parallel jobs if the execution is finished
    if ctrl.IsFinished() {
      return nil
    }

    mockCtx := guard.MockContext{
      Context: ctx,
      Mock:    floc.NewContext(),
    }
    defer mockCtx.Release()

    mockCtrl := floc.NewControl(mockCtx)
    defer mockCtrl.Release()

    // Run jobs in parallel
    for _, job := range jobs {
      // Run the job in it's own goroutine
      go func(job floc.Job) {
        var err error
        err = job(mockCtx, mockCtrl)
        handleResult(mockCtrl, err)
      }(job)
    }

    <-ctx.Done()

    // Wait until first jobs done
    res, data, err := mockCtrl.Result()
    switch res {
    case floc.Canceled:
      ctrl.Cancel(data)
    case floc.Completed:
      // Continue current flow
    case floc.Failed:
      ctrl.Fail(data, err)
    }

    return nil
  }
}
