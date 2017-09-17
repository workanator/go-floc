package errors

import (
	"bytes"
	"fmt"
)

// ErrMultiple contains multiple errors.
type ErrMultiple struct {
	list []error
}

// NewErrMultiple construct error instance from the list of errors given.
func NewErrMultiple(err error, errs ...error) ErrMultiple {
	errCount := len(errs)
	if errCount > 0 {
		list := make([]error, 1+errCount)
		list[0] = err

		for i, e := range errs {
			list[1+i] = e
		}

		return ErrMultiple{list}
	}

	return ErrMultiple{[]error{err}}
}

// Top returns the top most error from the contained list.
func (err ErrMultiple) Top() error {
	return err.list[0]
}

// List returns the list of errors contained.
func (err ErrMultiple) List() []error {
	return err.list[:]
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
	buf := &bytes.Buffer{}

	fmt.Fprintf(buf, "%d errors: ", len(err.list))
	for i, err := range err.list {
		if i != 0 {
			fmt.Fprint(buf, ", ")
		}
		fmt.Fprintf(buf, `"%v"`, err)
	}

	return buf.String()
}
