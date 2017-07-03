package floc

import "sync"

// State is the container of data shared amongst jobs. Depending on
// implementation the data can be thread-safe or not.
type State interface {
	// Returns the underlying state data with non-exclusive lock.
	Get() (data interface{}, locker sync.Locker)

	// Returns the underlying state data with exclusive lock.
	GetExclusive() (data interface{}, locker sync.Locker)
}
