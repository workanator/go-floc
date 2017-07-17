package guard

import "fmt"

const errorPrefix = "paniced with "

// ErrPanic is thrown with Cancel if no panic trigger is provided to Panic.
type ErrPanic struct {
	Data interface{}
}

func (err ErrPanic) Error() string {
	switch v := err.Data.(type) {
	case error:
		return fmt.Sprintf("%s%s", errorPrefix, v.Error())

	case fmt.Stringer:
		return fmt.Sprintf("%s%s", errorPrefix, v.String())

	default:
		return fmt.Sprintf("%s%v", errorPrefix, err.Data)
	}
}
