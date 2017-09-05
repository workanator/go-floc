package run

import floc "gopkg.in/workanator/go-floc.v1"

/*
Sequence runs jobs sequentially - one by one.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
  -->[JOB_1]-...->[JOB_N]-->
*/
func Sequence(jobs ...floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
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
