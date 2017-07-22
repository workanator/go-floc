package floc

import (
	"sync"
)

/*
StateContainer allows the state to contain any arbitrary data.
Data can either be of primitive type or complex structure or even
interface or function. What the state should contain depends on task.

  type Events struct {
    HeaderReady bool
    BodyReady bool
    DataReady bool
  }

  theState := state.New(new(Events))
*/
type StateContainer struct {
	sync.RWMutex
	data interface{}
}

// NewStateContainer create a new instance of the state container which can
// contain any arbitrary data.
func NewStateContainer(data interface{}) State {
	return &StateContainer{
		data: data,
	}
}

// Release releases all underlying resources. If the data contained implements
// floc.Releaser interface then Release() method of it is called.
func (s *StateContainer) Release() {
	if s.data != nil {
		if releaser, ok := s.data.(Releaser); ok {
			releaser.Release()
		}
	}
}

// Get returns the contained data with non-exclusive locker.
func (s *StateContainer) Get() (data interface{}, locker sync.Locker) {
	return s.data, (*stateRLocker)(s)
}

// GetExclusive returns the contained data with exclusive locker.
func (s *StateContainer) GetExclusive() (data interface{}, locker sync.Locker) {
	return s.data, s
}

type stateRLocker StateContainer

func (r *stateRLocker) Lock() {
	(*StateContainer)(r).RLock()
}

func (r *stateRLocker) Unlock() {
	(*StateContainer)(r).RUnlock()
}
