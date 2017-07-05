package run

import floc "github.com/workanator/go-floc"

/*
Loop repeats running jobs forever. Jobs are run sequentially.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Visual Representation:
    +-------------------------+
		|                         |
		V                         |
  ----->[JOB_1]-...->[JOB_N]--+
*/
func Loop(jobs ...floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		for {
			for _, job := range jobs {
				// Do not start the next job if the execution is finished
				if flow.IsFinished() {
					return
				}

				// Do the job
				job(flow, state, update)
			}
		}
	}
}
