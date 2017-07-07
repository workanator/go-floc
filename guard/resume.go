package guard

import (
	"fmt"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
)

// Resume resumes execution of the flow possibly canceled or completed by
// the job. If filter is empty or nil execution will be resumed regardless
// the reason it was finished. Otherwise execution will be resumed if the
// reason it finished with is in the filter result set.
func Resume(filter floc.ResultSet, job floc.Job) floc.Job {
	if len(filter) == 0 {
		// Result filtering is omitted so make the job simple with resuming always
		// happen.
		return func(theFlow floc.Flow, theState floc.State, theUpdate floc.Update) {
			resFlow, resume := flow.WithResume(theFlow)
			defer resume()

			job(resFlow, theState, theUpdate)
		}
	}

	// Make the job which is aware of the result.
	return func(theFlow floc.Flow, theState floc.State, theUpdate floc.Update) {
		resFlow, resume := flow.WithResume(theFlow)
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
						theFlow.Cancel(data)

					case floc.Completed:
						theFlow.Complete(data)

					default:
						// Something is wrong
						panic(fmt.Errorf("flow finished with unknown result %v and data %v", res, data))
					}
				}
			}
		}()

		job(resFlow, theState, theUpdate)
	}
}
