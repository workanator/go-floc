package guard

import (
	"time"

	"gopkg.in/devishot/go-floc.v2"
)

// WhenTimeoutFunc calculates the exact timeout passed in Timeout guards.
// The implementation may use the state and the id for accurate calculation
// of the timeout.
type WhenTimeoutFunc func(ctx floc.Context, id interface{}) time.Duration

// ConstTimeout returns constant timeout.
func ConstTimeout(timeout time.Duration) WhenTimeoutFunc {
	return func(ctx floc.Context, id interface{}) time.Duration {
		return timeout
	}
}
