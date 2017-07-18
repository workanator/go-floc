package state

import "testing"

type Releaseable bool

func (r *Releaseable) Release() {
	*r = true
}

func TestDefault(t *testing.T) {
	New(new(int)).Release()
	New(true).Release()
	New(func() string { return "Hello" }).Release()
	New(nil).Release()
}

func TestDefaultRead(t *testing.T) {
	state := New("Hello")
	defer state.Release()

	data, lock := state.Get()
	str := data.(string)

	lock.Lock()
	defer lock.Unlock()

	if str != "Hello" {
		t.Fatalf("%s expects Hello but has %s", t.Name(), str)
	}
}

func TestDefaultWrite(t *testing.T) {
	const max = 100

	state := New(new(int))
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

func TestDefaultInvalidCast(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%s must panic", t.Name())
		}
	}()

	state := New("Hello")
	defer state.Release()

	data, _ := state.Get()
	_ = data.(*string)
}

func TestDefaultReleaser(t *testing.T) {
	d := new(Releaseable)

	s := New(d)
	if *d != false {
		t.Fatalf("%s expects false but has %t", t.Name(), *d)
	}

	s.Release()
	if *d != true {
		t.Fatalf("%s expects true but has %t", t.Name(), *d)
	}
}
