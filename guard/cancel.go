package guard

import (
	"gopkg.in/devishot/go-floc.v2"
)

/*
Cancel cancels execution of the flow with the data given.
*/
func Cancel(data interface{}) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		ctrl.Cancel(data)
		return nil
	}
}
