package run

import "github.com/workanator/go-floc.v2"

func handleResult(ctrl floc.Control, err error) (stop bool) {
	if err != nil {
		ctrl.Fail(nil, err)
		return true
	}

	return false
}
