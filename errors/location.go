package errors

// ErrLocation contains the original error and the location where it happened.
type ErrLocation struct {
	what  error
	where string
}

// NewErrLocation constructs error instance from the original error and the location.
func NewErrLocation(what error, where string) ErrLocation {
	return ErrLocation{what, where}
}

// Where returns the location of the error.
func (err ErrLocation) Where() string {
	return err.where
}

// What returns the original error.
func (err ErrLocation) What() error {
	return err.what
}

func (err ErrLocation) Error() string {
	return err.where + ": " + err.what.Error()
}
