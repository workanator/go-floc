package guard

import (
	"time"

	floc "github.com/workanator/go-floc"
)

// Deadline protects the job from doing the job after the deadline. The job
// is run in it's own goroutine while the current goroutine waits until
// the job finished or the deadline came or the flow is finished.
func Deadline(whenDeadline WhenDeadlineFunc, id interface{}, job floc.Job) floc.Job {
	return DeadlineWithTrigger(whenDeadline, id, job, nil)
}

// DeadlineWithTrigger protects the job from doing the job after deadling.
// In addition it takes TimeoutTrigger func which called if time is out.
// The job is run in it's own goroutine while the current goroutine waits
// until the job finished or the deadline came or the flow is finished.
func DeadlineWithTrigger(whenDeadline WhenDeadlineFunc, id interface{}, job floc.Job, timeoutTrigger TimeoutTrigger) floc.Job {
	// Construct the job with Timeout guard.
	when := func(state floc.State, id interface{}) time.Duration {
		return time.Until(whenDeadline(state, id))
	}

	return TimeoutWithTrigger(when, id, job, timeoutTrigger)
}
