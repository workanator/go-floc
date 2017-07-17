package guard

import (
	"fmt"
	"time"
)

// ErrTimeout is thrown with Cancel if no panic trigger is provided to Timeout.
type ErrTimeout struct {
	ID interface{}
	At time.Time
}

func (err ErrTimeout) Error() string {
	return fmt.Sprintf("%v timed out at %s", err.ID, err.At)
}
