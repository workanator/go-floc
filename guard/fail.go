package guard

import (
	"gopkg.in/workanator/go-floc.v2"
)

/*
Fail cancels execution of the flow with the data and error given.
*/
func Fail(data interface{}, err error) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		ctrl.Fail(data, err)
		return nil
	}
}
