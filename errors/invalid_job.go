package errors

// ErrInvalidJob indicates that the job is invalid.
type ErrInvalidJob struct{}

const invalidJobMessage = "job is invalid"

func (ErrInvalidJob) Error() string {
	return invalidJobMessage
}
