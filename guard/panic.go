package guard

import floc "github.com/workanator/go-floc.v1"

// PanicTrigger is triggered when the coroutine state is recovered after
// panicing.
type PanicTrigger func(flow floc.Flow, state floc.State, v interface{})

// Panic protects the job from falling into panic. On panic the flow will
// be canceled with the ErrPanic result. Guarding the job from falling into
// panic is effective only if the job runs in the current goroutine.
func Panic(job floc.Job) floc.Job {
	return PanicWithTrigger(job, nil)
}

// IgnorePanic protects the job from falling into panic. On panic the panic
// will be ignored. Guarding the job from falling into
// panic is effective only if the job runs in the current goroutine.
func IgnorePanic(job floc.Job) floc.Job {
	return PanicWithTrigger(job, func(flow floc.Flow, state floc.State, v interface{}) {})
}

// PanicWithTrigger protects the job from falling into panic. In addition it
// takes PanicTrigger func which is called in case of panic. Guarding the job
// from falling into panic is effective only if the job runs in the current
// goroutine.
func PanicWithTrigger(job floc.Job, panicTrigger PanicTrigger) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		defer func() {
			if r := recover(); r != nil {
				if panicTrigger != nil {
					panicTrigger(flow, state, r)
				} else {
					flow.Cancel(ErrPanic{Data: r})
				}
			}
		}()

		// Do the job
		job(flow, state, update)
	}
}
