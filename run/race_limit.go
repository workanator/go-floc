package run

import (
	"sync"

	floc "github.com/workanator/go-floc"
)

/*
RaceLimit runs jobs in their own goroutines and waits until first N jobs finish.
During the race only first N calls to update are allowed while further calls
are discarded. Before starting the race the function synchronizes start of the
each job putting them in equal conditions.

If limit is less than 1 or greater than the number of jobs the function will
panic.

Summary:
	- Run jobs in goroutines : YES
	- Wait all jobs finish   : YES
	- Run order              : PARALLEL with synchronization of start

Diagram:
    +-->[JOB_1]--+
    |            |
  --+-->  ..   --+-->
    |            |
    +-->[JOB_N]--+
*/
func RaceLimit(limit int, jobs ...floc.Job) floc.Job {
	// Validate the winner limit
	if limit < 1 || limit > len(jobs) {
		panic("invalid amount of possible race winners")
	}

	return func(flow floc.Flow, state floc.State, update floc.Update) {
		// Do not start the race if the execution is finished
		if flow.IsFinished() {
			return
		}

		// Create the channel which will have the value when the race is won
		done := make(chan struct{}, len(jobs))
		defer close(done)

		// Wrap the flow into disablable flow so the calls to Cancel and Complete
		// can be disabled when the race is won
		disFlow, disable := floc.NewFlowWithDisable(flow)

		// Wrap the trigger to a function which allows to hit the update only
		// `limit` time(s)
		mutex := sync.Mutex{}
		winnerJobs := 0

		limitedUpdate := func(flow floc.Flow, state floc.State, key string, value interface{}) {
			mutex.Lock()
			defer mutex.Unlock()

			if winnerJobs < limit {
				winnerJobs++
				update(flow, state, key, value)

				if winnerJobs == limit {
					disable()
				}
			}
		}

		// Condition is used to synchronize start of jobs
		canStart := false
		startMutex := &sync.RWMutex{}
		startCond := sync.NewCond(startMutex.RLocker())

		// Run jobs in parallel and wait until all of them ready to start
		runningJobs := 0
		for _, job := range jobs {
			runningJobs++

			go func(job floc.Job) {
				defer func() { done <- struct{}{} }()

				// Wait for the start of the race
				startCond.L.Lock()
				for !canStart && !flow.IsFinished() {
					startCond.Wait()
				}
				startCond.L.Unlock()

				// Perform the job if the flow is not finished
				if !flow.IsFinished() {
					job(disFlow, state, limitedUpdate)
				}
			}(job)
		}

		// Notify all jobs they can start the race
		startMutex.Lock()
		canStart = true
		startCond.Broadcast()
		startMutex.Unlock()

		// Wait until `limit` job(s) done
		finishedJobs := 0
		for finishedJobs < runningJobs {
			select {
			case <-disFlow.Done():
				// The execution has been finished or canceled so the trigger
				// should not be triggered anymore, therefore we disable it. But we
				// do not return and wait until all jobs finish the race.
				disable()

			case <-done:
				// One of the jobs finished.
				finishedJobs++
				if finishedJobs == limit {
					disable()
				}
			}
		}
	}
}
