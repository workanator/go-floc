package run

import floc "github.com/workanator/go-floc"

/*
IfOrElse runs jobTrue if the condition is met or runs jobFalse otherwise.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
                      +----->[JOB_TRUE]---+
                      | YES               |
  --(CONDITION MET?)--+                   +-->
                      | NO                |
                      +----->[JOB_FALSE]--+
*/
func IfOrElse(predicate floc.Predicate, jobTrue, jobFalse floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		if predicate(state) {
			jobTrue(flow, state, update)
		} else {
			jobFalse(flow, state, update)
		}
	}
}
