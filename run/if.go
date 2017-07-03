package run

import floc "github.com/workanator/go-floc"

// If runs job if the condition met.
func If(predicate floc.Predicate, job floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		if predicate(flow, state) {
			job(flow, state, update)
		}
	}
}
