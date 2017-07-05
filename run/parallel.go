package run

import floc "github.com/workanator/go-floc"

/*
Parallel runs jobs in their own goroutines and waits until all of them finish.

Summary:
	- Run jobs in goroutines : YES
	- Wait all jobs finish   : YES
	- Run order              : PARALLEL

Visual Representation:
    +-->[JOB_1]--+
    |            |
  --+-->  ..   --+-->
    |            |
    +-->[JOB_N]--+
*/
func Parallel(jobs ...floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		// Do not start parallel jobs if the execution is finished
		if flow.IsFinished() {
			return
		}

		// Create channel which is used for back counting of finished jobs
		done := make(chan struct{}, len(jobs))
		defer close(done)

		// Run jobs in parallel
		running := 0
		for _, job := range jobs {
			running++

			// Run the job in it's own goroutine
			go func(job floc.Job) {
				defer func() { done <- struct{}{} }()
				job(flow, state, update)
			}(job)
		}

		// Wait until all jobs done
		for running > 0 {
			select {
			case <-flow.Done():
				// The execution finished but we should wait until all jobs finished
				// and we assume all jobs are aware of the flow state. If we do
				// not wait that may lead to unpredicted behavior.

			case <-done:
				// One of the jobs finished
				running--
			}
		}
	}
}
