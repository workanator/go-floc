package guard

import (
	"gopkg.in/workanator/go-floc.v2"
)

// Resume resumes execution of the flow possibly canceled or completed by
// the job. If filter is empty or nil execution will be resumed regardless
// the reason it was finished. Otherwise execution will be resumed if the
// reason it finished with is in the filter result set.
func Resume(filter floc.ResultSet, job floc.Job) floc.Job {
	if len(filter) == 0 {
		// Result filtering is omitted so make the job simple with resuming always
		// happen.
		return func(ctx floc.Context, ctrl floc.Control) error {
			mockCtx := floc.NewContext()
			mockCtrl := floc.NewControl(mockCtx)

			return job(ctx, mockCtrl)
		}
	}

	// Make the job which is aware of the result.
	return func(ctx floc.Context, ctrl floc.Control) error {
		mockCtx := floc.NewContext()
		mockCtrl := floc.NewControl(mockCtx)

		defer func() {
			// Test if execution finished first
			if mockCtrl.IsFinished() {
				res, data, err := mockCtrl.Result()
				if !filter.Contains(res) {
					// Propagate the result
					switch res {
					case floc.Canceled:
						mockCtrl.Cancel(data)
					case floc.Completed:
						mockCtrl.Complete(data)
					case floc.Failed:
						mockCtrl.Fail(data, err)
					}
				}
			}
		}()

		return job(ctx, mockCtrl)
	}
}
