package run

import (
	"time"

	floc "github.com/workanator/go-floc"
)

/*
Delay does delay before starting each job. Jobs are run sequentially.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE
*/
func Delay(delay time.Duration, jobs ...floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		// Declare the timer and setup defered cleanup task for it.
		var timer *time.Timer
		defer func() {
			if timer != nil {
				timer.Stop()
			}
		}()

		for _, job := range jobs {
			// Do not start the next job if the flow is finished
			if flow.IsFinished() {
				return
			}

			// Create or reset the timer
			if timer == nil {
				timer = time.NewTimer(delay)
			} else {
				// Reseting timer means it was expired.
				timer.Reset(delay)
			}

			// Wait until delay passed or the execution of the flow is finished
			select {
			case <-flow.Done():
				// The execution finished, stop running jobs on timer.
				break

			case <-timer.C:
				// Do the job
				job(flow, state, update)
			}
		}
	}
}
