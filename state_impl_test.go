package floc

import (
	"fmt"
	"testing"
)

type Releaseable bool

func (r *Releaseable) Release() {
	*r = true
}

func TestState(t *testing.T) {
	const max = 100

	for i := 0; i < max; i++ {
		// Integer
		sInt := NewState(i)
		if sInt.Data().(int) != i {
			t.Fatalf("%s expects int to be %d but has %d", t.Name(), i, sInt.Data().(int))
		}

		// Allocated integer
		vInt := new(int)
		*vInt = i

		sAllocInt := NewState(vInt)
		if *(sAllocInt.Data().(*int)) != i {
			t.Fatalf("%s expects allocated int to be %d but has %d", t.Name(), i, *(sAllocInt.Data().(*int)))
		}

		// Boolean
		b := i%2 == 0

		sBool := NewState(b)
		if sBool.Data().(bool) != b {
			t.Fatalf("%s expects bool to be %t but has %t", t.Name(), b, sBool.Data().(bool))
		}

		// String
		s := fmt.Sprintf("STR%d", i)

		sStr := NewState(s)
		if sStr.Data().(string) != s {
			t.Fatalf("%s expects string to be %s but has %s", t.Name(), s, sStr.Data().(string))
		}

		// Func
		f := func() string { return s }

		sFunc := NewState(f)
		if (sFunc.Data().(func() string))() != f() {
			t.Fatalf("%s expects func to return %s but has %s", t.Name(), f(), (sFunc.Data().(func() string))())
		}

		// Nil
		sNil := NewState(nil)
		if sNil.Data() != nil {
			t.Fatalf("%s expects nil but has %v", t.Name(), sNil.Data())
		}
	}
}

func TestStateRead(t *testing.T) {
	state := NewState("Hello")
	defer state.Release()

	data, lock := state.DataWithReadLocker()
	str := data.(string)

	lock.Lock()
	defer lock.Unlock()

	if str != "Hello" {
		t.Fatalf("%s expects Hello but has %s", t.Name(), str)
	}
}

func TestStateWrite(t *testing.T) {
	const max = 100

	state := NewState(new(int))
	defer state.Release()

	// Increment 100 times
	dataEx, lockEx := state.DataWithWriteLocker()
	counter := dataEx.(*int)

	for i := 0; i < max; i++ {
		lockEx.Lock()
		*counter++
		lockEx.Unlock()
	}

	// Read the result
	data, lock := state.DataWithReadLocker()
	result := data.(*int)

	lock.Lock()
	defer lock.Unlock()

	if *result != max {
		t.Fatalf("%s expects %d but has %d", t.Name(), max, *result)
	}
}

func TestStateReadWrite(t *testing.T) {
	const max = 100

	state := NewState(new(int))
	defer state.Release()

	data, readLocker, writeLocker := state.DataWithReadAndWriteLockers()
	counter := data.(*int)

	for i := 0; i < max; i++ {
		// Write
		writeLocker.Lock()
		*counter = i
		writeLocker.Unlock()

		// Read
		readLocker.Lock()
		current := *counter
		readLocker.Unlock()

		// Compare
		if current != i {
			t.Fatalf("%s expects %d but has %d", t.Name(), i, current)
		}
	}
}

func TestStateInvalidCast(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%s must panic", t.Name())
		}
	}()

	state := NewState("Hello")
	defer state.Release()

	data := state.Data()
	_ = data.(*string)
}

func TestStateReleaser(t *testing.T) {
	d := new(Releaseable)

	s := NewState(d)
	if *d != false {
		t.Fatalf("%s expects false but has %t", t.Name(), *d)
	}

	s.Release()
	if *d != true {
		t.Fatalf("%s expects true but has %t", t.Name(), *d)
	}
}
