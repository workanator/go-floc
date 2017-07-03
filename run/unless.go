package run

import floc "github.com/workanator/go-floc"

// Unless runs job if the condition does not met.
func Unless(predicate floc.Predicate, job floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		if !predicate(flow, state) {
			job(flow, state, update)
		}
	}
}
