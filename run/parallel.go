package run

import (
	"github.com/workanator/go-floc"
)

/*
Parallel runs jobs in their own goroutines and waits until all of them finish.

Summary:
	- Run jobs in goroutines : YES
	- Wait all jobs finish   : YES
	- Run order              : PARALLEL

Diagram:
    +-->[JOB_1]--+
    |            |
  --+-->  ..   --+-->
    |            |
    +-->[JOB_N]--+
*/
func Parallel(jobs ...floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		// Do not start parallel jobs if the execution is finished
		if ctrl.IsFinished() {
			return nil
		}

		// Create channel which is used for back counting of finished jobs
		done := make(chan error, len(jobs))
		defer close(done)

		// Run jobs in parallel
		jobsRunning := 0
		for _, job := range jobs {
			jobsRunning++

			// Run the job in it's own goroutine
			go func(job floc.Job) {
				var err error
				defer func() { done <- err }()
				err = job(ctx, ctrl)
			}(job)
		}

		// Wait until all jobs done
		errs := make([]error, 0, len(jobs))

		for jobsRunning > 0 {
			select {
			case <-ctx.Done():
				// The execution finished but we should wait until all jobs finished
				// and we assume all jobs are aware of the flow state. If we do
				// not wait that may lead to unpredicted behavior.

			case err := <-done:
				// One of the jobs finished
				if handledErr := handleResult(ctrl, err); handledErr != nil {
					errs = append(errs, handledErr)
				}

				jobsRunning--
			}
		}

		if len(errs) > 0 {
			return floc.NewErrMultiple(errs[0], errs[1:]...)
		}

		return nil
	}
}
