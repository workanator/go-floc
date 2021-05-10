package run

import (
	"time"

	"github.com/workanator/go-floc/v3"
)

/*
Wait waits until the condition is met. The function falls into sleep with the
duration given between condition checks. The function does not run any job
actually and just repeatedly checks predicate's return value. When the predicate
returns true the function finishes.

Summary:
	- Run jobs in goroutines : N/A
	- Wait all jobs finish   : N/A
	- Run order              : N/A

Diagram:
                    NO
    +------(SLEEP)------+
    |                   |
    V                   | YES
  ----(CONDITION MET?)--+----->
*/
func Wait(predicate floc.Predicate, sleep time.Duration) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		for !ctrl.IsFinished() && !predicate(ctx) {
			time.Sleep(sleep)
		}

		return nil
	}
}
