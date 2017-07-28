package floc

import "sync"

/*
State is the container of data shared amongst jobs. Depending on
implementation the data can be thread-safe or not.

The state is aware of possible implementation of Releaser interface by contained
data. So if the contained data implements Releaser call to state.Release() will
be propagated to data.Release() as well.

  type Data struct{}

  func (Data) Release() {
    fmt.Println("Data released")
  }

  state := floc.NewState(Data{})
  state.Release()

  // Output: Data released
*/
type State interface {
	Releaser

	// Returns the underlying state data with non-exclusive lock.
	Get() (data interface{}, locker sync.Locker)

	// Returns the underlying state data with exclusive lock.
	GetExclusive() (data interface{}, locker sync.Locker)
}
