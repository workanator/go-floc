package run

import floc "github.com/workanator/go-floc"

// Sequence runs jobs sequentially.
func Sequence(jobs ...floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
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
