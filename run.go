package floc

import "gopkg.in/workanator/go-floc.v2/errors"

func Run(job Job) (result Result, data interface{}, err error) {
	// Return invalid job error if the job is nil
	if job == nil {
		return None, nil, errors.ErrInvalidJob{}
	}

	// Create context and control
	ctx := NewContext()
	defer ctx.Release()

	ctrl := NewControl(ctx)
	defer ctrl.Release()

	// Run the flow and return the result
	unhandledErr := job(ctx, ctrl)

	result, data, err = ctrl.Result()
	if result != None {
		return result, data, err
	}

	// Return Failed if unhandled error left after the execution.
	if unhandledErr != nil {
		return Failed, nil, unhandledErr
	}

	return None, nil, nil
}
