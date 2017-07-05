/*
Package run is the collection of jobs which make the architecture of the flow.

Each function in the package is middleware which always takes at leats one
floc.Job to run and constructs and returns another floc.Job. That allows to
organize jobs in any combination and in result is only one floc.Job which can
be run with floc.Run().
*/
package run
