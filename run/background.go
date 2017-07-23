package run

import floc "github.com/workanator/go-floc"

/*
Background starts each job in it's own goroutine. The function does not
track the lifecycle of jobs started and does no synchronization with them
therefore all running in background jobs may remain active even if the flow
is finished. The function assumes all jobs are aware of the flow state and/or
synchronization and termination of them is implemented outside.

	floc.Run(flow, state, update, run.Background {
		func(flow floc.FLow, state floc.State, update floc.Update) {
			for !flow.IsFinished() {
				fmt.Println(time.Now())
			}
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
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		for _, job := range jobs {
			// Do not start the next job if the execution is finished
			if flow.IsFinished() {
				return
			}

			// Run the job
			go job(flow, state, update)
		}
	}
}
