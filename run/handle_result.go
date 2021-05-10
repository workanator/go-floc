package run

import (
	"github.com/workanator/go-floc/v3"
)

func handleResult(ctrl floc.Control, err error) error {
	if err != nil {
		ctrl.Fail(nil, err)
		return err
	}

	return nil
}
