package state

import "testing"

func TestDefault(t *testing.T) {
	New(new(int))
	New(true)
	New(func() string { return "Hello" })
	New(nil)
}

func TestDefaultRead(t *testing.T) {
	state := New("Hello")

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

	data, _ := state.Get()
	_ = data.(*string)
}
