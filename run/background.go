package run

import (
	"gopkg.in/workanator/go-floc.v2"
)

const locBackground = "Background"

/*
Background starts the job in it's own goroutine. The function does not
track the lifecycle of the job started and does no synchronization with it
therefore the job running in background may remain active even if the flow
is finished. The function assumes the job is aware of the flow state and/or
synchronization and termination of it is implemented outside.

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
	- Run order              : SINGLE

Diagram:
  --+----------->
    |
    +-->[JOB]
*/
func Background(job floc.Job) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		// Do not start the job if the flow is finished
		if ctrl.IsFinished() {
			return nil
		}

		// Run the job in background
		go func(job floc.Job) {
			err := job(ctx, ctrl)
			handleResult(ctrl, err, locBackground)
		}(job)

		return nil
	}
}
