package errors

import (
	"bytes"
	"fmt"
)

// ErrMultiple contains multiple errors.
type ErrMultiple struct {
	errs []error
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

// What returns the top most error from the contained list.
func (err ErrMultiple) Err() error {
	return err.errs[0]
}

// Errs returns the list of errors contained.
func (err ErrMultiple) Errs() []error {
	return err.errs[:]
}

func (err ErrMultiple) Error() string {
	// Return the first error if only one error is contained.
	if len(err.errs) == 1 {
		return err.errs[0].Error()
	}

	// Build the string from all errors contained.
	buf := &bytes.Buffer{}

	fmt.Fprintf(buf, "%d errors: ", len(err.errs))
	for i, err := range err.errs {
		if i != 0 {
			fmt.Fprint(buf, ", ")
		}
		fmt.Fprintf(buf, `"%v"`, err)
	}

	return buf.String()
}
