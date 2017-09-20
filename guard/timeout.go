package guard

import (
	"time"

	"gopkg.in/workanator/go-floc.v2"
	"gopkg.in/workanator/go-floc.v2/errors"
)

// TimeoutTrigger triggers when the execution of the job timed out.
type TimeoutTrigger func(ctx floc.Context, ctrl floc.Control, id interface{})

// Timeout protects the job from taking too much time on execution.
// The job is run in it's own goroutine while the current goroutine waits
// until the job finished or time went out or the flow is finished.
func Timeout(when WhenTimeoutFunc, id interface{}, job floc.Job) floc.Job {
	return OnTimeout(when, id, job, nil)
}

// OnTimeout protects the job from taking too much time on execution.
// In addition it takes TimeoutTrigger func which called if time is out.
// The job is run in it's own goroutine while the current goroutine waits
// until the job finished or time went out or the flow is finished.
func OnTimeout(when WhenTimeoutFunc, id interface{}, job floc.Job, timeoutTrigger TimeoutTrigger) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		// Create the channel to read the result error
		done := make(chan error)
		defer close(done)

		// Create timer
		timer := time.NewTimer(when(ctx, id))
		defer timer.Stop()

		// Run the job
		go func() {
			var err error
			defer func() { done <- err }()
			err = job(ctx, ctrl)
		}()

		// Wait for one of possible events
		select {
		case <-ctx.Done():
			// The execution finished

		case err := <-done:
			// The job finished. Return the result immediately.
			return err

		case <-timer.C:
			// The execution is timed out
			if timeoutTrigger != nil {
				timeoutTrigger(ctx, ctrl, id)
			} else {
				ctrl.Fail(id, errors.NewErrTimeout(id, time.Now().UTC()))
			}
		}

		// Wait until the job is finished and return the result.
		return <-done
	}
}
