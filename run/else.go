package run

import "gopkg.in/workanator/go-floc.v2"

/*
Else just returns the job unmodified. Else is used for expressiveness
and can be omitted.
*/
func Else(job floc.Job) floc.Job {
	return job
}
