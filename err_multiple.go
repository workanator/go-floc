package floc

import (
	"strconv"
	"strings"
)

// ErrMultiple contains multiple errors.
type ErrMultiple struct {
	list []error
}

// NewErrMultiple construct error instance from the list of errors given.
func NewErrMultiple(errs ...error) ErrMultiple {
	return ErrMultiple{
		list: errs,
	}
}

// Top returns the top most error from the contained list.
func (err ErrMultiple) Top() error {
	if len(err.list) > 0 {
		return err.list[0]
	}
	return nil
}

// List returns the list of errors contained.
func (err ErrMultiple) List() []error {
	return err.list
}

// Len returns the count of errors contained.
func (err ErrMultiple) Len() int {
	return len(err.list)
}

func (err ErrMultiple) Error() string {
	// Return the first error if only one error is contained.
	if len(err.list) == 1 {
		return err.list[0].Error()
	}

	// Build the string from all errors contained.
	sb := strings.Builder{}

	sb.WriteString(strconv.Itoa(len(err.list)))
	sb.WriteString(" errors: ")
	for i, err := range err.list {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteByte('"')
		sb.WriteString(err.Error())
		sb.WriteByte('"')
	}

	return sb.String()
}
