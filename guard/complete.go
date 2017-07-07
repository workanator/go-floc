package guard

import floc "github.com/workanator/go-floc"

// Complete completes execution of the flow with the data given.
func Complete(data interface{}) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		flow.Complete(data)
	}
}
