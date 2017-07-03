package run

import floc "github.com/workanator/go-floc"

// Race runs jobs in parallel and waits until the first job finish.
func Race(jobs ...floc.Job) floc.Job {
	return RaceLimit(1, jobs...)
}
