package floc

import (
	"sync"
)

type stateContainer struct {
	sync.RWMutex
	data interface{}
}

/*
NewState create a new instance of the state container which can contain any
arbitrary data. Data can either be of primitive type or complex structure or
even interface or function. What the state should contain depends on task.

  type Events struct {
    HeaderReady bool
    BodyReady bool
    DataReady bool
  }

  state := floc.NewState(new(Events))

The container can contain nil value as well if no contained data is required.

  state := floc.NewState(nil)
*/
func NewState(data interface{}) State {
	return &stateContainer{
		data: data,
	}
}

// Release releases all underlying resources. If the data contained implements
// floc.Releaser interface then Release() method of it is called.
func (s *stateContainer) Release() {
	if s.data != nil {
		if releaser, ok := s.data.(Releaser); ok {
			releaser.Release()
		}
	}
}

// Date returns the contained data.
func (s *stateContainer) Data() (data interface{}) {
	return s.data
}

// DataWithReadLocker returns the contained data with read-only locker.
func (s *stateContainer) DataWithReadLocker() (data interface{}, readLocker sync.Locker) {
	return s.data, (*stateRLocker)(s)
}

// DataWithWriteLocker returns the contained data with read-write locker.
func (s *stateContainer) DataWithWriteLocker() (data interface{}, writeLocker sync.Locker) {
	return s.data, s
}

// DataWithReadAndWriteLockers returns the contained data with read-only
// and read/write lockers.
func (s *stateContainer) DataWithReadAndWriteLockers() (data interface{}, readLocker, writeLocker sync.Locker) {
	return s.data, (*stateRLocker)(s), s
}

type stateRLocker stateContainer

func (r *stateRLocker) Lock() {
	(*stateContainer)(r).RLock()
}

func (r *stateRLocker) Unlock() {
	(*stateContainer)(r).RUnlock()
}
