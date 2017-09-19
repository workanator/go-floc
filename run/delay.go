package run

import (
	"time"

	"gopkg.in/workanator/go-floc.v2"
)

const locDelay = "Delay"

/*
Delay does delay before starting the job.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SINGLE

Diagram:
  --(DELAY)-->[JOB]-->
*/
func Delay(delay time.Duration, job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		// Do not start the job if the flow is finished
		if ctrl.IsFinished() {
			return nil
		}

		// Declare the timer and setup deferred cleanup task for it.
		timer := time.NewTimer(delay)
		defer timer.Stop()

		// Wait until delay passed or the execution of the flow is finished
		select {
		case <-ctx.Done():
			// The execution finished, stop running jobs on timer.

		case <-timer.C:
			// Do the job
			err := job(ctx, ctrl)
			if handlerErr := handleResult(ctrl, err, locDelay); handlerErr != nil {
				return handlerErr
			}
		}

		return nil
	}
}
