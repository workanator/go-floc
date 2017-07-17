package guard

import floc "github.com/workanator/go-floc"

/*
Cancel cancels execution of the flow with the data given.
*/
func Cancel(data interface{}) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		flow.Cancel(data)
	}
}
