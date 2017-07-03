package run

import (
	"time"

	floc "github.com/workanator/go-floc"
)

// Delay runs jobs sequentially with delay.
func Delay(delay time.Duration, jobs ...floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		// Create timer
		timer := time.NewTimer(delay)
		defer timer.Stop()

		for _, job := range jobs {
			// Do not start the next job if the execution is finished
			if flow.IsFinished() {
				return
			}

			// Reset the timer
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(delay)

			// Wait until delay passed or the execution of the flow is finished
			select {
			case <-flow.Done():
				// The execution is finished
				return

			case <-timer.C:
				// Do the job
				job(flow, state, update)
			}
		}
	}
}
