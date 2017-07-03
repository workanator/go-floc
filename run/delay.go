package run

import (
	"time"

	floc "github.com/workanator/go-floc"
)

// Delay runs jobs sequentially with delay.
func Delay(delay time.Duration, jobs ...floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		// Create a channel which help to catch the moment when delay exceeded
		done := make(chan int)
		defer close(done)

		for _, job := range jobs {
			// Do not start the next job if the execution is finished
			if flow.IsFinished() {
				return
			}

			// Run a job which will sleep for the given amount of time
			go func() {
				defer func() { done <- 0 }()
				time.Sleep(delay)
			}()

			// Wait until delay passed of the execution of the flow is finished
			select {
			case <-flow.Done():
				// The execution is finished
				return

			case <-done:
				// Do the job
				job(flow, state, update)
			}
		}
	}
}
