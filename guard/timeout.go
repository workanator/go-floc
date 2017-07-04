package guard

import (
	"fmt"
	"time"

	floc "github.com/workanator/go-floc"
)

// ErrTimeout is thrown with Cancel if not panic trigger is provided to Timeout.
type ErrTimeout struct {
	id interface{}
	at time.Time
}

// TimeoutTrigger triggers when the execution of the job timed out.
type TimeoutTrigger func(flow floc.Flow, state floc.State, id interface{})

// Timeout protects the job from taking too much time on execution.
func Timeout(timeout time.Duration, id interface{}, job floc.Job) floc.Job {
	return TimeoutWithTrigger(timeout, id, job, nil)
}

// TimeoutWithTrigger protects the job from taking too much time on execution.
// In addition it takes TimeoutTrigger func which called if time is out.
func TimeoutWithTrigger(timeout time.Duration, id interface{}, job floc.Job, timeoutTrigger TimeoutTrigger) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		done := make(chan struct{})
		defer close(done)

		// Run the job
		go func() {
			defer func() { done <- struct{}{} }()
			job(flow, state, update)
		}()

		// Create timer
		timer := time.NewTimer(timeout)
		defer timer.Stop()

		// Wait for one of possible events
		jobFinished := false
		select {
		case <-flow.Done():
			// The execution is finished

		case <-done:
			// The job finished
			jobFinished = true

		case <-timer.C:
			// The execution is timed out
			if timeoutTrigger != nil {
				timeoutTrigger(flow, state, id)
			} else {
				flow.Cancel(ErrTimeout{id: id, at: time.Now().UTC()})
			}
		}

		// Wait for the job to finish before return
		if !jobFinished {
			<-done
		}
	}
}

func (err ErrTimeout) Error() string {
	return fmt.Sprintf("%v timed out at %s", err.id, err.at)
}
