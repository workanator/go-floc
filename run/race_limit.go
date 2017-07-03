package run

import (
	"sync"

	floc "github.com/workanator/go-floc"

	"github.com/workanator/go-floc/flow"
)

// RaceLimit runs jobs in parallel and waits until the first N jobs finish.
func RaceLimit(limit int, jobs ...floc.Job) floc.Job {
	return func(theFlow floc.Flow, state floc.State, update floc.Update) {
		// Validate the winner limit
		if limit < 1 || limit > len(jobs) {
			panic("invalid amount of possible race winners")
		}

		// Do not start the race if the execution is finished
		if theFlow.IsFinished() {
			return
		}

		// Create the channel which will have the value when the race is won
		done := make(chan int, limit)
		defer close(done)

		// Wrap the flow into disablable flow so the calls to Cancel and Complete
		// can be disabled when the race is won
		disFlow, disable := flow.WithDisable(theFlow)

		// Wrap the trigger to a function which allows to hit the trigger only once
		mutex := sync.Mutex{}
		won := 0

		limitedUpdate := func(flow floc.Flow, state floc.State, key string, value interface{}) {
			mutex.Lock()
			defer mutex.Unlock()

			if won < limit {
				won++
				update(flow, state, key, value)
				done <- won

				if won == limit {
					disable()
				}
			}
		}

		// Run jobs in parallel
		for _, job := range jobs {
			go job(disFlow, state, limitedUpdate)
		}

		// Wait until one job done
		winners := 0
		for winners < limit {
			select {
			case <-disFlow.Done():
				// The execution has been finished or canceled so the trigger
				// should not be triggered therefore we disable it
				disable()
				return

			case <-done:
				// One of the jobs finished. We have a winner! Count it.
				winners++
			}
		}
	}
}
