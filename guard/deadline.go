package guard

import (
	"time"

	floc "github.com/workanator/go-floc"
)

// Deadline protects the job from doing the job after the deadline. The job
// is run in it's own goroutine while the current goroutine waits until
// the job finished or the deadline came or the flow is finished.
func Deadline(deadline time.Time, id interface{}, job floc.Job) floc.Job {
	return DeadlineWithTrigger(deadline, id, job, nil)
}

// DeadlineWithTrigger protects the job from doing the job after deadling.
// In addition it takes TimeoutTrigger func which called if time is out.
// The job is run in it's own goroutine while the current goroutine waits
// until the job finished or the deadline came or the flow is finished.
func DeadlineWithTrigger(deadline time.Time, id interface{}, job floc.Job, timeoutTrigger TimeoutTrigger) floc.Job {
	// If the deadline passed return a job which just triggers timeout trigger.
	if time.Now().After(deadline) {
		if timeoutTrigger != nil {
			return func(flow floc.Flow, state floc.State, update floc.Update) {
				timeoutTrigger(flow, state, id)
			}
		}

		return func(flow floc.Flow, state floc.State, update floc.Update) {
			flow.Cancel(ErrTimeout{ID: id, At: time.Now().UTC()})
		}
	}

	// Construct the job with Timeout guard.
	return TimeoutWithTrigger(time.Until(deadline), id, job, timeoutTrigger)
}
