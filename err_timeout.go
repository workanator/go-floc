package floc

import (
	"fmt"
	"time"
)

// ErrTimeout is thrown with Cancel if no panic trigger is provided to Timeout.
type ErrTimeout struct {
	id interface{}
	at time.Time
}

const tplTimeoutMessage = "%v timed out at %s"

// NewErrTimeout constructs new instance of ErrTimeout.
func NewErrTimeout(id interface{}, at time.Time) ErrTimeout {
	return ErrTimeout{id, at}
}

// ID returns the ID of the timeout happened.
func (err ErrTimeout) ID() interface{} {
	return err.id
}

// At returns the time when the timeout happened.
func (err ErrTimeout) At() time.Time {
	return err.at
}

func (err ErrTimeout) Error() string {
	return fmt.Sprintf(tplTimeoutMessage, err.id, err.at)
}
