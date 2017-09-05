package guard

import (
	"time"

	floc "gopkg.in/workanator/go-floc.v1"
)

// TimeoutTrigger triggers when the execution of the job timed out.
type TimeoutTrigger func(flow floc.Flow, state floc.State, id interface{})

// Timeout protects the job from taking too much time on execution.
// The job is run in it's own goroutine while the current goroutine waits
// until the job finished or time went out or the flow is finished.
func Timeout(whenTimeout WhenTimeoutFunc, id interface{}, job floc.Job) floc.Job {
	return TimeoutWithTrigger(whenTimeout, id, job, nil)
}

// TimeoutWithTrigger protects the job from taking too much time on execution.
// In addition it takes TimeoutTrigger func which called if time is out.
// The job is run in it's own goroutine while the current goroutine waits
// until the job finished or time went out or the flow is finished.
func TimeoutWithTrigger(whenTimeout WhenTimeoutFunc, id interface{}, job floc.Job, timeoutTrigger TimeoutTrigger) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		done := make(chan struct{})
		defer close(done)

		// Run the job
		go func() {
			defer func() { done <- struct{}{} }()
			job(flow, state, update)
		}()

		// Create timer
		timer := time.NewTimer(whenTimeout(state, id))
		defer timer.Stop()

		// Wait for one of possible events
		jobFinished := false
		select {
		case <-flow.Done():
			// The execution finished

		case <-done:
			// The job finished
			jobFinished = true

		case <-timer.C:
			// The execution is timed out
			if timeoutTrigger != nil {
				timeoutTrigger(flow, state, id)
			} else {
				flow.Cancel(ErrTimeout{ID: id, At: time.Now().UTC()})
			}
		}

		// Wait anyway for the job to finish before return because if we do not
		// that may result in unpredicted behavior. And we assume the job is
		// aware of the flow state, e.g. it tests periodically if it's finished.
		if !jobFinished {
			<-done
		}
	}
}
