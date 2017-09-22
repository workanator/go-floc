package floc

import "fmt"

/*
Result identifies the result of execution.
*/
type Result int32

/*
Possible results.
*/
const (
	None         Result = 1
	Completed    Result = 2
	Canceled     Result = 4
	Failed       Result = 8
	usedBitsMask Result = None | Completed | Canceled | Failed
	finishedMask Result = Completed | Canceled | Failed
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
	return result&finishedMask != 0
}

// IsValid tests if the result is a valid value.
func (result Result) IsValid() bool {
	return result == None || result == Completed || result == Canceled || result == Failed
}

// Mask constructs ResultMask with only one result masked.
func (result Result) Mask() ResultMask {
	return NewResultMask(result)
}

// i32 returns the underlying value as int32.
func (result Result) i32() int32 {
	return int32(result)
}

func (result Result) String() string {
	switch result {
	case None:
		return "None"
	case Completed:
		return "Completed"
	case Canceled:
		return "Canceled"
	case Failed:
		return "Failed"
	default:
		return fmt.Sprintf("Result(%d)", result.i32())
	}
}
