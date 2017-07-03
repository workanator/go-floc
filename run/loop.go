package run

import floc "github.com/workanator/go-floc"

// Loop repeats running jobs until the execution of the flow is finished.
func Loop(jobs ...floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		for {
			for _, job := range jobs {
				// Do not start the next job if the execution is finished
				if flow.IsFinished() {
					return
				}

				// Do the job
				job(flow, state, update)
			}
		}
	}
}
