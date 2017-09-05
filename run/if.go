package run

import floc "gopkg.in/workanator/go-floc.v1"

/*
If runs the job if the condition is met.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
                      +----->[JOB]---+
                      | YES          |
  --(CONDITION MET?)--+              +-->
                      | NO           |
                      +--------------+
*/
func If(predicate floc.Predicate, job floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		if predicate(state) {
			job(flow, state, update)
		}
	}
}
