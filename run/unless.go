package run

import floc "github.com/workanator/go-floc.v1"

/*
Unless runs the job if the condition is not met.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
                      +-------------+
                      | YES         |
  --(CONDITION MET?)--+             +-->
                      | NO          |
                      +---->[JOB]---+
*/
func Unless(predicate floc.Predicate, job floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		if !predicate(state) {
			job(flow, state, update)
		}
	}
}
