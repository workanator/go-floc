package run

import (
	"github.com/workanator/go-floc.v2"
)

const locRepeat = "Repeat"

/*
Repeat repeats running jobs for N times. Jobs start sequentially.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
                          NO
    +-----------[JOB]<---------+
    |                          |
    V                          | YES
  ----(ITERATED COUNT TIMES?)--+---->
*/
func Repeat(count int, jobs ...floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for n := 1; n <= count; n++ {
			for _, job := range jobs {
				// Do not start the next job if the execution is finished
				if ctrl.IsFinished() {
					return nil
				}

				// Do the job
				err := job(ctx, ctrl)
				if handledErr := handleResult(ctrl, err, locRepeat); handledErr != nil {
					return handledErr
				}
			}
		}

		return nil
	}
}
