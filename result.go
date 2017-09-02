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

	resultFirst = None
	resultLast  = Failed
)

// IsNone tests if the result is None.
func (result Result) IsNone() bool {
	return result == None
}

// IsCompleted tests if the result is Completed.
func (result Result) IsCompleted() bool {
	return result == Completed
}

// IsCanceled tests if the result is Canceled.
func (result Result) IsCanceled() bool {
	return result == Canceled
}

// IsFailed tests if the result is Failed.
func (result Result) IsFailed() bool {
	return result == Failed
}

// IsFinished tests if the result is either Completed or Canceled or Failed.
func (result Result) IsFinished() bool {
	return result == Completed || result == Canceled || result == Failed
}

// IsValid tests if the result is a valid value.
func (result Result) IsValid() bool {
	return result >= resultFirst && result <= resultLast
}

// Int32 returns the underlying value as int32. That is handy while working
// with atomic operations.
func (result Result) Int32() int32 {
	return int32(result)
}
