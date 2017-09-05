package guard

import (
	floc "gopkg.in/workanator/go-floc.v1"
)

// Resume resumes execution of the flow possibly canceled or completed by
// the job. If filter is empty or nil execution will be resumed regardless
// the reason it was finished. Otherwise execution will be resumed if the
// reason it finished with is in the filter result set.
func Resume(filter floc.ResultSet, job floc.Job) floc.Job {
	if len(filter) == 0 {
		// Result filtering is omitted so make the job simple with resuming always
		// happen.
		return func(flow floc.Flow, state floc.State, update floc.Update) {
			resFlow, resume := floc.NewFlowWithResume(flow)
			defer resume()

			job(resFlow, state, update)
		}
	}

	// Make the job which is aware of the result.
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		resFlow, resume := floc.NewFlowWithResume(flow)
		defer func() {
			// Test if execution finished first
			if resFlow.IsFinished() {
				res, data := resFlow.Result()
				if filter.Contains(res) {
					// Resume execution
					resume()
				} else {
					// Propagate the result
					switch res {
					case floc.Canceled:
						flow.Cancel(data)

					case floc.Completed:
						flow.Complete(data)
					}
				}
			}
		}()

		job(resFlow, state, update)
	}
}
