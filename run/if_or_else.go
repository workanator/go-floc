package run

import (
	"gopkg.in/workanator/go-floc.v2"
)

const locIfOrElse = "IfOrElse"

/*
IfOrElse runs jobTrue if the condition is met or runs jobFalse otherwise.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
                      +----->[JOB_TRUE]---+
                      | YES               |
  --(CONDITION MET?)--+                   +-->
                      | NO                |
                      +----->[JOB_FALSE]--+
*/
func IfOrElse(predicate floc.Predicate, jobTrue, jobFalse floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		// Do not start the job if the execution is finished
		if ctrl.IsFinished() {
			return nil
		}

		// Test the predicate and run the appropriate job
		var err error
		if predicate(ctx) {
			err = jobTrue(ctx, ctrl)
		} else {
			err = jobFalse(ctx, ctrl)
		}

		if handlerErr := handleResult(ctrl, err, locIfOrElse); handlerErr != nil {
			return handlerErr
		}

		return nil
	}
}
