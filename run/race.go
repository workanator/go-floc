package run

import floc "github.com/workanator/go-floc"

/*
Race runs jobs in their own goroutines and waits until the first job finishes.
During the race only the first call to update is allowed while further calls
are discarded. Before starting the race the function synchronizes start of the
each job putting them in equal conditions.

Summary:
	- Run jobs in goroutines : YES
	- Wait all jobs finish   : YES
	- Run order              : PARALLEL with syncronization of start
*/
func Race(jobs ...floc.Job) floc.Job {
	return RaceLimit(1, jobs...)
}
