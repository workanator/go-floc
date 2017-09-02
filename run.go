package floc

func Run(job Job) (result Result, data interface{}, err error) {
	// Return invalid job error if the job is nil
	if job == nil {
		return None, nil, ErrInvalidJob{}
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
