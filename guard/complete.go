package guard

import floc "gopkg.in/workanator/go-floc.v1"

/*
Complete completes execution of the flow with the data given.
*/
func Complete(data interface{}) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		flow.Complete(data)
	}
}
