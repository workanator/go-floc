package errors

// ErrLocation contains the original error and the location where it happened.
type ErrLocation struct {
	err   error
	where string
}

// NewErrLocation constructs error instance from the original error and the location.
func NewErrLocation(err error, where string) ErrLocation {
	return ErrLocation{err, where}
}

// Where returns the location of the error.
func (err ErrLocation) Where() string {
	return err.where
}

// Err returns the original error.
func (err ErrLocation) Err() error {
	return err.err
}

func (err ErrLocation) Error() string {
	return err.where + ": " + err.err.Error()
}
