package run

import "github.com/workanator/go-floc/v3"

/*
Then just returns the job unmodified. Then is used for expressiveness
and can be omitted.

Summary:
	- Run jobs in goroutines : N/A
	- Wait all jobs finish   : N/A
	- Run order              : N/A

Diagram:
  ----[JOB]--->
*/
func Then(job floc.Job) floc.Job {
	return job
}
