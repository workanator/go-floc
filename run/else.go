package run

import "github.com/workanator/go-floc"

/*
Else just returns the job unmodified. Else is used for expressiveness
and can be omitted.

Summary:
	- Run jobs in goroutines : N/A
	- Wait all jobs finish   : N/A
	- Run order              : N/A

Diagram:
  ----[JOB]--->
*/
func Else(job floc.Job) floc.Job {
	return job
}
