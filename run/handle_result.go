package run

import (
	"gopkg.in/workanator/go-floc.v2"
)

func handleResult(ctrl floc.Control, err error, where string) error {
	if err != nil {
		ctrl.Fail(nil, err)
		return floc.NewErrLocation(err, where)
	}

	return nil
}
