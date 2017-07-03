package run

import floc "github.com/workanator/go-floc"

// While runs job while the condition met.
func While(predicate floc.Predicate, job floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		for predicate(flow, state) {
			job(flow, state, update)
		}
	}
}
