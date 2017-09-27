package run

import (
	"gopkg.in/workanator/go-floc.v2"
)

func handleResult(ctrl floc.Control, err error) error {
	if err != nil {
		ctrl.Fail(nil, err)
		return err
	}

	return nil
}
