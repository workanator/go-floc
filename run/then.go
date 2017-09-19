package run

import "gopkg.in/workanator/go-floc.v2"

/*
Then just returns the job unmodified. Then is used for expressiveness
and can be omitted.
*/
func Then(job floc.Job) floc.Job {
	return job
}
