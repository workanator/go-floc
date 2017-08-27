package floc

/*
Result identifies the result of execution.
 */
type Result int32

/*
Possible results.
 */
const (
	None Result = iota
	Completed
	Canceled
	Failed
)