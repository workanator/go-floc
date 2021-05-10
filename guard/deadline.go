package guard

import (
	"time"

	"github.com/workanator/go-floc/v3"
)

// Deadline protects the job from doing the job after the deadline. The job
// is run in it's own goroutine while the current goroutine waits until
// the job finished or the deadline came or the flow is finished.
func Deadline(when WhenDeadlineFunc, id interface{}, job floc.Job) floc.Job {
	return OnDeadline(when, id, job, nil)
}

// OnDeadline protects the job from doing the job after deadline.
// In addition it takes TimeoutTrigger func which called if time is out.
// The job is run in it's own goroutine while the current goroutine waits
// until the job finished or the deadline came or the flow is finished.
func OnDeadline(when WhenDeadlineFunc, id interface{}, job floc.Job, timeoutTrigger TimeoutTrigger) floc.Job {
	// Construct the job with Timeout guard.
	whenTimeout := func(ctx floc.Context, id interface{}) time.Duration {
		return time.Until(when(ctx, id))
	}

	return OnTimeout(whenTimeout, id, job, timeoutTrigger)
}
