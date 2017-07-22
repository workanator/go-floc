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

  state := floc.NewStateContainer(new(Events))

The container can contain nil value as well if no contained data is required.

  state := floc.NewStateContainer(nil)
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

// Get returns the contained data with non-exclusive locker.
func (s *stateContainer) Get() (data interface{}, locker sync.Locker) {
	return s.data, (*stateRLocker)(s)
}

// GetExclusive returns the contained data with exclusive locker.
func (s *stateContainer) GetExclusive() (data interface{}, locker sync.Locker) {
	return s.data, s
}

type stateRLocker stateContainer

func (r *stateRLocker) Lock() {
	(*stateContainer)(r).RLock()
}

func (r *stateRLocker) Unlock() {
	(*stateContainer)(r).RUnlock()
}
