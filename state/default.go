package state

import (
	"sync"

	floc "github.com/workanator/go-floc"
)

type defaultState struct {
	sync.RWMutex
	data interface{}
}

// New create a new instance of the default state container which can
// contain any arbitrary data.
func New(data interface{}) floc.State {
	return &defaultState{
		data: data,
	}
}

// Release releases all underlying resources. If the data contained implements
// floc.Releaser interface then Release() method of it is called.
func (s *defaultState) Release() {
	if s.data != nil {
		if releaser, ok := s.data.(floc.Releaser); ok {
			releaser.Release()
		}
	}
}

// Returns the underlying state data with non-exclusive locker.
func (s *defaultState) Get() (data interface{}, locker sync.Locker) {
	return s.data, (*defaultStateRLocker)(s)
}

// Returns the underlying state data with exclusive locker.
func (s *defaultState) GetExclusive() (data interface{}, locker sync.Locker) {
	return s.data, s
}

type defaultStateRLocker defaultState

func (r *defaultStateRLocker) Lock() {
	(*defaultState)(r).RLock()
}

func (r *defaultStateRLocker) Unlock() {
	(*defaultState)(r).RUnlock()
}
