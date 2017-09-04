package run

import (
	"time"

	"github.com/workanator/go-floc.v2"
)

const locDelay = "Delay"

/*
Delay does delay before starting each job. Jobs are run sequentially.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
  --(DELAY)-->[JOB_1]-...-(DELAY)-->[JOB_N]-->
*/
func Delay(delay time.Duration, jobs ...floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		// Declare the timer and setup deferred cleanup task for it.
		var timer *time.Timer
		defer func() {
			if timer != nil {
				timer.Stop()
			}
		}()

		for _, job := range jobs {
			// Do not start the next job if the flow is finished
			if ctrl.IsFinished() {
				return nil
			}

			// Create or reset the timer
			if timer == nil {
				timer = time.NewTimer(delay)
			} else {
				// Reset timer means it was expired.
				timer.Reset(delay)
			}

			// Wait until delay passed or the execution of the flow is finished
			select {
			case <-ctrl.Done():
				// The execution finished, stop running jobs on timer.
				break

			case <-timer.C:
				// Do the job
				err := job(ctx, ctrl)
				if handlerErr := handleResult(ctrl, err, locDelay); handlerErr != nil {
					return handlerErr
				}
			}
		}

		return nil
	}
}
