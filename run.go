package floc

import "github.com/workanator/go-floc/errors"

// Run creates a new Context and Control and runs the flow.
func Run(job Job) (result Result, data interface{}, err error) {
	// Create context and control
	ctx := NewContext()
	defer ctx.Release()

	ctrl := NewControl(ctx)
	defer ctrl.Release()

	// Run the flow
	return RunWith(ctx, ctrl, job)
}

// RunWith runs the flow with the Context and Control given.
func RunWith(ctx Context, ctrl Control, job Job) (result Result, data interface{}, err error) {
	// Return invalid job error if the job is nil
	if job == nil {
		return None, nil, errors.ErrInvalidJob{}
	}

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
