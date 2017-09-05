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
	job(ctx, ctrl)

	return ctrl.Result()
}
