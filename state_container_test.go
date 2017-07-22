package floc

import "testing"

type Releaseable bool

func (r *Releaseable) Release() {
	*r = true
}

func TestStateContainer(t *testing.T) {
	NewStateContainer(new(int)).Release()
	NewStateContainer(true).Release()
	NewStateContainer(func() string { return "Hello" }).Release()
	NewStateContainer(nil).Release()
}

func TestStateContainerRead(t *testing.T) {
	state := NewStateContainer("Hello")
	defer state.Release()

	data, lock := state.Get()
	str := data.(string)

	lock.Lock()
	defer lock.Unlock()

	if str != "Hello" {
		t.Fatalf("%s expects Hello but has %s", t.Name(), str)
	}
}

func TestStateContainerWrite(t *testing.T) {
	const max = 100

	state := NewStateContainer(new(int))
	defer state.Release()

	// Increment 100 times
	dataEx, lockEx := state.GetExclusive()
	counter := dataEx.(*int)

	for i := 0; i < max; i++ {
		lockEx.Lock()
		*counter++
		lockEx.Unlock()
	}

	// Read the result
	data, lock := state.Get()
	result := data.(*int)

	lock.Lock()
	defer lock.Unlock()

	if *result != max {
		t.Fatalf("%s expects %d but has %d", t.Name(), max, *result)
	}
}

func TestStateContainerInvalidCast(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%s must panic", t.Name())
		}
	}()

	state := NewStateContainer("Hello")
	defer state.Release()

	data, _ := state.Get()
	_ = data.(*string)
}

func TestStateContainerReleaser(t *testing.T) {
	d := new(Releaseable)

	s := NewStateContainer(d)
	if *d != false {
		t.Fatalf("%s expects false but has %t", t.Name(), *d)
	}

	s.Release()
	if *d != true {
		t.Fatalf("%s expects true but has %t", t.Name(), *d)
	}
}
