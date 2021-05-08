package run

import (
	"github.com/workanator/go-floc"
)

/*
Sequence runs jobs sequentially, one by one.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
  -->[JOB_1]-...->[JOB_N]-->
*/
func Sequence(jobs ...floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for _, job := range jobs {
			// Do not start the next job if the execution is finished
			if ctrl.IsFinished() {
				return nil
			}

			// Run the job
			err := job(ctx, ctrl)
			if handledErr := handleResult(ctrl, err); handledErr != nil {
				return handledErr
			}
		}

		return nil
	}
}
