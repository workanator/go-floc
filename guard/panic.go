package guard

import (
	"gopkg.in/workanator/go-floc.v2"
	"gopkg.in/workanator/go-floc.v2/errors"
)

const locPanic = "Panic"

// PanicTrigger is triggered when the goroutine state is recovered after
// panic.
type PanicTrigger func(ctx floc.Context, ctrl floc.Control, v interface{})

// Panic protects the job from falling into panic. On panic the flow will
// be canceled with the ErrPanic result. Guarding the job from falling into
// panic is effective only if the job runs in the current goroutine.
func Panic(job floc.Job) floc.Job {
	return OnPanic(job, nil)
}

// IgnorePanic protects the job from falling into panic. On panic the panic
// will be ignored. Guarding the job from falling into
// panic is effective only if the job runs in the current goroutine.
func IgnorePanic(job floc.Job) floc.Job {
	return OnPanic(job, func(ctx floc.Context, ctrl floc.Control, v interface{}) {})
}

// OnPanic protects the job from falling into panic. In addition it
// takes PanicTrigger func which is called in case of panic. Guarding the job
// from falling into panic is effective only if the job runs in the current
// goroutine.
func OnPanic(job floc.Job, panicTrigger PanicTrigger) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		defer func() {
			if r := recover(); r != nil {
				if panicTrigger != nil {
					panicTrigger(ctx, ctrl, r)
				} else {
					ctrl.Fail(r, errors.NewErrPanic(r))
				}
			}
		}()

		// Do the job
		return job(ctx, ctrl)
	}
}
