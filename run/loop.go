package run

import (
	"gopkg.in/workanator/go-floc.v2"
)

const locLoop = "Loop"

/*
Loop repeats running the job forever.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SINGLE

Diagram:
    +----------+
    |          |
    V          |
  ----->[JOB]--+
*/
func Loop(job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for {
			// Do not start the job if the execution is finished
			if ctrl.IsFinished() {
				return nil
			}

			// Do the job
			err := job(ctx, ctrl)
			if handledErr := handleResult(ctrl, err, locLoop); handledErr != nil {
				return handledErr
			}
		}
	}
}
