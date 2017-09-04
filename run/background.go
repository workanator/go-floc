package run

import (
	"github.com/workanator/go-floc.v2"
)

const locBackground = "Background"

/*
Background starts each job in it's own goroutine. The function does not
track the lifecycle of jobs started and does no synchronization with them
therefore all running in background jobs may remain active even if the flow
is finished. The function assumes all jobs are aware of the flow state and/or
synchronization and termination of them is implemented outside.

	floc.Run(run.Background(
		func(ctx floc.Context, ctrl floc.Control) error {
			for !ctrl.IsFinished() {
				fmt.Println(time.Now())
			}

			return nil
		}
	})

Summary:
	- Run jobs in goroutines : YES
	- Wait all jobs finish   : NO
	- Run order              : SEQUENCE

Diagram:
  --+------------>
    |
    +-->[JOB_1]
    |
    ...
    |
    +-->[JOB_N]
*/
func Background(jobs ...floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for _, job := range jobs {
			// Do not start the next job if the execution is finished
			if ctrl.IsFinished() {
				return nil
			}

			// Run the job in background
			go func(job floc.Job) {
				err := job(ctx, ctrl)
				handleResult(ctrl, err, locBackground)
			}(job)
		}

		return nil
	}
}
