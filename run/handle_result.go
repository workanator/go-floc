package run

import (
	"gopkg.in/workanator/go-floc.v2"
	"gopkg.in/workanator/go-floc.v2/errors"
)

func handleResult(ctrl floc.Control, err error, where string) error {
	if err != nil {
		ctrl.Fail(nil, err)
		return errors.NewErrLocation(err, where)
	}

	return nil
}
