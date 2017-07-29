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

	// Returns the contained data.
	Data() (data interface{})

	// Returns the contained data with read-only locker.
	DataWithReadLocker() (data interface{}, readLocker sync.Locker)

	// Returns the contained data with read/write locker.
	DataWithWriteLocker() (data interface{}, writeLocker sync.Locker)

	// Returns the contained data with read-only and read/write lockers.
	DataWithReadAndWriteLockers() (data interface{}, readLocker, writeLocker sync.Locker)
}
