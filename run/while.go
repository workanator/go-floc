package run

import (
	"gopkg.in/workanator/go-floc.v2"
)

const locWhile = "While"

/*
While repeats running the job while the condition is met.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
                    YES
    +-------[JOB]<------+
    |                   |
    V                   | NO
  ----(CONDITION MET?)--+---->
*/
func While(predicate floc.Predicate, job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for !ctrl.IsFinished() && predicate(ctx) {
			err := job(ctx, ctrl)
			if handledErr := handleResult(ctrl, err, locWhile); handledErr != nil {
				return handledErr
			}
		}

		return nil
	}
}
