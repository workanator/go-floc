package run

import (
	"github.com/workanator/go-floc.v2"
)

const locLoop = "Loop"

/*
Loop repeats running jobs forever. Jobs are run sequentially.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
    +-------------------------+
    |                         |
    V                         |
  ----->[JOB_1]-...->[JOB_N]--+
*/
func Loop(jobs ...floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for {
			for _, job := range jobs {
				// Do not start the next job if the execution is finished
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
}
