package run

import "github.com/workanator/go-floc.v2"

const locIf = "If"

/*
If runs the job if the condition is met.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
                      +----->[JOB]---+
                      | YES          |
  --(CONDITION MET?)--+              +-->
                      | NO           |
                      +--------------+
*/
func If(predicate floc.Predicate, job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		if predicate(ctx) {
			err := job(ctx, ctrl)
			if handledErr := handleResult(ctrl, err, locIf); handledErr != nil {
				return handledErr
			}
		}

		return nil
	}
}
