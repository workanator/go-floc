package guard

import (
	"github.com/workanator/go-floc"
)

/*
Complete completes execution of the flow with the data given.
*/
func Complete(data interface{}) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		ctrl.Complete(data)
		return nil
	}
}
