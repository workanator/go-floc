package run

import floc "github.com/workanator/go-floc"

// IfOrElse runs jobTrue if the condition met or jobFalse otherwise.
func IfOrElse(predicate floc.Predicate, jobTrue, jobFalse floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		if predicate(flow, state) {
			jobTrue(flow, state, update)
		} else {
			jobFalse(flow, state, update)
		}
	}
}
