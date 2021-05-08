package floc

import (
	"fmt"
	"testing"
)

func TestErrMultiple_Top(t *testing.T) {
	const max = 100

	for n := 1; n <= max; n++ {
		list := make([]error, n)
		for i := 0; i < n; i++ {
			list[i] = fmt.Errorf("%d", i)
		}

		err := NewErrMultiple(list...)
		if err.Top().Error() != list[0].Error() {
			t.Fatalf("%s expects error to be %s but has %s", t.Name(), list[0].Error(), err.Top().Error())
		}
	}
}

func TestErrMultiple_List(t *testing.T) {
	const max = 100

	for n := 1; n <= max; n++ {
		list := make([]error, n)
		for i := 0; i < n; i++ {
			list[i] = fmt.Errorf("%d", i)
		}

		err := NewErrMultiple(list...)
		if err.Len() != len(list) {
			t.Fatalf("%s expects error count to be %d but has %d", t.Name(), len(list), err.Len())
		}

		for k, e := range err.List() {
			if e.Error() != list[k].Error() {
				t.Fatalf("%s expects error to be %s but has %s", t.Name(), list[k].Error(), e.Error())
			}
		}
	}
}

func TestErrMultiple_ErrorOne(t *testing.T) {
	e := fmt.Errorf("ERROR")
	err := NewErrMultiple(e)
	if err.Error() != e.Error() {
		t.Fatalf("%s expects error to be %s but has %s", t.Name(), e.Error(), err.Error())
	}
}

func TestErrMultiple_ErrorMany(t *testing.T) {
	e1 := fmt.Errorf("err1")
	e2 := fmt.Errorf("err2")
	e3 := fmt.Errorf("err3")

	msg := fmt.Sprintf(`3 errors: "%v", "%v", "%v"`, e1, e2, e3)

	err := NewErrMultiple(e1, e2, e3)
	if err.Error() != msg {
		t.Fatalf("%s expects error to be %s but has %s", t.Name(), msg, err.Error())
	}
}
