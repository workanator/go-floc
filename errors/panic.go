package errors

import "fmt"

const errorPrefix = "panic with "

// ErrPanic is thrown with Cancel if no panic trigger is provided to Panic.
type ErrPanic struct {
	Data interface{}
}

func (err ErrPanic) Error() string {
	switch v := err.Data.(type) {
	case error:
		return errorPrefix + v.Error()

	case fmt.Stringer:
		return errorPrefix + v.String()

	default:
		return fmt.Sprintf("%s%v", errorPrefix, err.Data)
	}
}
