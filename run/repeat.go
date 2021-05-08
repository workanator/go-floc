package run

import (
	"github.com/workanator/go-floc"
)

/*
Repeat repeats running the job for N times.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SINGLE

Diagram:
                          NO
    +-----------[JOB]<---------+
    |                          |
    V                          | YES
  ----(ITERATED COUNT TIMES?)--+---->
*/
func Repeat(times int, job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for n := 1; n <= times; n++ {
			// Do not start the job if the execution is finished
			if ctrl.IsFinished() {
				return nil
			}

			// Do the job
			err := job(ctx, ctrl)
			if handledErr := handleResult(ctrl, err); handledErr != nil {
				return handledErr
			}
		}

		return nil
	}
}
