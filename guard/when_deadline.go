package guard

import (
	"time"

	floc "gopkg.in/workanator/go-floc.v1"
)

// WhenDeadlineFunc calculates the exact deadline passed in Deadline guards.
// The implementation may use the state and the id for accurate calculation
// of the deadline.
type WhenDeadlineFunc func(state floc.State, id interface{}) time.Time

// ConstDeadline returns constant deadline.
func ConstDeadline(deadline time.Time) WhenDeadlineFunc {
	return func(state floc.State, id interface{}) time.Time {
		return deadline
	}
}

// DeadlineIn calculates the deadline in the future with constant delay
// from now.
func DeadlineIn(delay time.Duration) WhenDeadlineFunc {
	return func(state floc.State, id interface{}) time.Time {
		return time.Now().Add(delay)
	}
}
