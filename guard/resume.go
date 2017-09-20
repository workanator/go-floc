package guard

import (
	"gopkg.in/workanator/go-floc.v2"
)

// Mock context which propagates all calls to the parent context
// but Done() returns mock channel.
type mockContext struct {
	floc.Context
	mock floc.Context
}

// Release releases the mock context.
func (ctx mockContext) Release() {
	ctx.mock.Release()
}

// Done returns the channel of the mock context.
func (ctx mockContext) Done() <-chan struct{} {
	return ctx.mock.Done()
}

// Resume resumes execution of the flow possibly canceled or completed by
// the job. If filter is empty or nil execution will be resumed regardless
// the reason it was finished. Otherwise execution will be resumed if the
// reason it finished with is in the filter result set.
func Resume(filter floc.ResultSet, job floc.Job) floc.Job {
	// If result filtering is omitted make the job simple with resuming always
	// happen.
	if len(filter) == 0 {
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
				if !filter.Contains(res) {
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
