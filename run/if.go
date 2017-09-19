package run

import (
	"gopkg.in/workanator/go-floc.v2"
)

const (
	locIf    = "If"
	idxTrue  = 0
	idxFalse = 1
)

/*
If runs the first job if the condition is met and runs
the second job, if it's passed, if the condition is not met.
The function panics if no or more than two jobs are given.

For expressiveness Then() and Else() can be used.

  flow := run.If(testSomething,
    run.Then(doSomething),
    run.Else(doSomethingElse),
  )

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SINGLE

Diagram:
                      +----->[JOB_1]---+
                      | YES            |
  --(CONDITION MET?)--+                +-->
                      | NO             |
                      +----->[JOB_2]---+
*/
func If(predicate floc.Predicate, jobs ...floc.Job) floc.Job {
	count := len(jobs)
	if count == 1 {
		return func(ctx floc.Context, ctrl floc.Control) error {
			// Do not start the job if the execution is finished
			if ctrl.IsFinished() {
				return nil
			}

			// Test the predicate and run the job on success
			if predicate(ctx) {
				err := jobs[idxTrue](ctx, ctrl)
				if handledErr := handleResult(ctrl, err, locIf); handledErr != nil {
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
			if predicate(ctx) {
				err = jobs[idxTrue](ctx, ctrl)
			} else {
				err = jobs[idxFalse](ctx, ctrl)
			}

			if handlerErr := handleResult(ctrl, err, locIf); handlerErr != nil {
				return handlerErr
			}

			return nil
		}
	}

	panic("If requires one or two jobs")
}
