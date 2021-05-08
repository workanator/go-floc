package guard

import (
	"github.com/workanator/go-floc"
)

// Resume resumes execution of the flow possibly finished by the job.
// If the mask is empty execution will be resumed regardless the reason
// it was finished. Otherwise execution will be resumed if the reason
// it finished with is masked.
func Resume(mask floc.ResultMask, job floc.Job) floc.Job {
	// If the mask is empty make the job simple with resuming always
	// happen.
	if mask.IsEmpty() {
		return func(ctx floc.Context, ctrl floc.Control) error {
			mockCtx := floc.NewContext()
			defer mockCtx.Release()

			mockCtrl := floc.NewControl(mockCtx)
			defer mockCtrl.Release()

			return job(mockContext{ctx, mockCtx}, mockCtrl)
		}
	}

	// Make the job which is aware of the result.
	return func(ctx floc.Context, ctrl floc.Control) error {
		mockCtx := mockContext{
			Context: ctx,
			mock:    floc.NewContext(),
		}

		mockCtrl := floc.NewControl(mockCtx)

		defer func() {
			mockCtrl.Release()
			mockCtx.Release()

			// Test if execution finished first
			if mockCtrl.IsFinished() {
				res, data, err := mockCtrl.Result()
				if !mask.IsMasked(res) {
					// Propagate the result
					switch res {
					case floc.Canceled:
						ctrl.Cancel(data)
					case floc.Completed:
						ctrl.Complete(data)
					case floc.Failed:
						ctrl.Fail(data, err)
					}
				}
			}
		}()

		return job(mockCtx, mockCtrl)
	}
}
