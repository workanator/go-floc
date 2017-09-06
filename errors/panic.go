package errors

import "fmt"

const errorPrefix = "panic with "

// ErrPanic contains information about panic.
type ErrPanic struct {
	data interface{}
}

// NewErrPanic constructs new instance of ErrPanic.
func NewErrPanic(data interface{}) ErrPanic {
	return ErrPanic{data}
}

// Data returns the data contained in panic caused the error.
func (err ErrPanic) Data() interface{} {
	return err.data
}

func (err ErrPanic) Error() string {
	switch v := err.data.(type) {
	case error:
		return errorPrefix + v.Error()

	case fmt.Stringer:
		return errorPrefix + v.String()

	default:
		return fmt.Sprintf("%s%v", errorPrefix, err.data)
	}
}
