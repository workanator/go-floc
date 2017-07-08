package run

import (
	"time"

	floc "github.com/workanator/go-floc"
)

/*
Wait waits until the condition is met. The function falls into sleep with the
duration given between condition checks. The function does not run any job
actually and just repeatedly checks predicate return value. When the predicate
returns true the function finishes.

It's a good idea to use synchronization mechanizm like sync.Cond to make
the function sleep until the contidion is ready.

Summary:
	- Run jobs in goroutines : N/A
	- Wait all jobs finish   : N/A
	- Run order              : N/A

Visual Representation:
                    NO
    +-------------------+
    |                   |
    V                   | YES
  ----(CONDITION MET?)--+----->
*/
func Wait(predicate floc.Predicate, duration time.Duration) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		for !predicate(flow, state) && !flow.IsFinished() {
			time.Sleep(duration)
		}
	}
}
