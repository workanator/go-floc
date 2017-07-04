package guard

import (
	"time"

	floc "github.com/workanator/go-floc"
)

// Deadline protects the job from doing the job after deadling.
func Deadline(deadline time.Time, id interface{}, job floc.Job) floc.Job {
	return DeadlineWithTrigger(deadline, id, job, nil)
}

// DeadlineWithTrigger protects the job from doing the job after deadling.
// In addition it takes TimeoutTrigger func which called if time is out.
func DeadlineWithTrigger(deadline time.Time, id interface{}, job floc.Job, timeoutTrigger TimeoutTrigger) floc.Job {
	// If the deadline passed return a job which just triggers timeout trigger.
	if time.Now().After(deadline) {
		return func(flow floc.Flow, state floc.State, update floc.Update) {
			if timeoutTrigger != nil {
				timeoutTrigger(flow, state, id)
			} else {
				flow.Cancel(ErrTimeout{id: id, at: time.Now().UTC()})
			}
		}
	}

	// Constructr the job with Timeout guard.
	return TimeoutWithTrigger(time.Until(deadline), id, job, timeoutTrigger)
}
