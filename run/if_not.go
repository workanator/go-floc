package run

import (
	"github.com/workanator/go-floc/v3"
)

/*
IfNot runs the first job if the condition is not met and runs
the second job, if it's passed, if the condition is met.
The function panics if no or more than two jobs are given.

For expressiveness Then() and Else() can be used.

  flow := run.IfNot(testSomething,
    run.Then(doSomething),
    run.Else(doSomethingElse),
  )

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SINGLE

Diagram:
                      +----->[JOB_1]---+
                      | NO             |
  --(CONDITION MET?)--+                +-->
                      | YES            |
                      +----->[JOB_2]---+
*/
func IfNot(predicate floc.Predicate, jobs ...floc.Job) floc.Job {
	count := len(jobs)
	if count == 1 {
		return func(ctx floc.Context, ctrl floc.Control) error {
			// Do not start the job if the execution is finished
			if ctrl.IsFinished() {
				return nil
			}

			// Test the predicate and run the job on success
			if !predicate(ctx) {
				err := jobs[idxTrue](ctx, ctrl)
				if handledErr := handleResult(ctrl, err); handledErr != nil {
					return handledErr
				}
			}

			return nil
		}
	} else if count == 2 {
		return func(ctx floc.Context, ctrl floc.Control) error {
			// Do not start the job if the execution is finished
			if ctrl.IsFinished() {
				return nil
			}

			// Test the predicate and run the appropriate job
			var err error
			if !predicate(ctx) {
				err = jobs[idxTrue](ctx, ctrl)
			} else {
				err = jobs[idxFalse](ctx, ctrl)
			}

			if handlerErr := handleResult(ctrl, err); handlerErr != nil {
				return handlerErr
			}

			return nil
		}
	}

	panic("IfNot requires one or two jobs")
}
