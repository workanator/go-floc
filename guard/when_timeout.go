package guard

import (
	"time"

	floc "github.com/workanator/go-floc.v1"
)

// WhenTimeoutFunc calculates the exact timeout passed in Timeout guards.
// The implementation may use the state and the id for accurate calculation
// of the timeout.
type WhenTimeoutFunc func(state floc.State, id interface{}) time.Duration

// ConstTimeout returns constant timeout.
func ConstTimeout(timeout time.Duration) WhenTimeoutFunc {
	return func(state floc.State, id interface{}) time.Duration {
		return timeout
	}
}
