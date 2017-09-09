package run

import (
	"gopkg.in/workanator/go-floc.v2"
	"gopkg.in/workanator/go-floc.v2/errors"
)

func handleResult(ctrl floc.Control, err error, where string) error {
	if err != nil {
		locationErr := errors.NewErrLocation(err, where)
		ctrl.Fail(nil, locationErr)

		return locationErr
	}

	return nil
}
