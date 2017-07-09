package run

import floc "github.com/workanator/go-floc"

/*
Repeat repeats running jobs for N times. Jobs are start sequentially.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
                          NO
    +-----------[JOB]<---------+
    |                          |
    V                          | YES
  ----(ITERATED COUNT TIMES?)--+---->
*/
func Repeat(count int, jobs ...floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		for n := 1; n <= count; n++ {
			for _, job := range jobs {
				// Do not start the next job if the execution is finished
				if flow.IsFinished() {
					return
				}

				// Do the job
				job(flow, state, update)
			}
		}
	}
}
