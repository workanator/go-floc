package run

import (
	"gopkg.in/workanator/go-floc.v2"
)

const locUnless = "Unless"

/*
Unless runs the job if the condition is not met.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
                      +-------------+
                      | YES         |
  --(CONDITION MET?)--+             +-->
                      | NO          |
                      +---->[JOB]---+
*/
func Unless(predicate floc.Predicate, job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		// Do not start the job if the execution is finished
		if ctrl.IsFinished() {
			return nil
		}

		// Test the predicate and run the job
		if !predicate(ctx) {
			err := job(ctx, ctrl)
			if handledErr := handleResult(ctrl, err, locUnless); handledErr != nil {
				return handledErr
			}
		}

		return nil
	}
}
