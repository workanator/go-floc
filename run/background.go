package run

import floc "github.com/workanator/go-floc"

// Background runs jobs in background and forget about them. All running
// in background jobs will remain active even if the flow is closed.
func Background(jobs ...floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		for _, job := range jobs {
			// Do not start the next job if the execution is finished
			if flow.IsFinished() {
				return
			}

			// Run the job
			go job(flow, state, update)
		}
	}
}
