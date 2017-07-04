package guard

import "fmt"

// ErrPanic is thrown with Cancel if no panic trigger is provided to Panic.
type ErrPanic struct {
	err interface{}
}

func (err ErrPanic) Error() string {
	return fmt.Sprintf("%v", err.err)
}
