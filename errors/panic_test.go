package errors

import (
	"fmt"
	"testing"
)

type Str string

func (s Str) String() string {
	return string(s)
}

func TestErrPanic_DataError(t *testing.T) {
	tpl := "ERROR"
	msg := fmt.Sprintf("%s%s", errorPrefix, tpl)

	err := NewErrPanic(fmt.Errorf(tpl))
	if err.Error() != msg {
		t.Fatalf("%s expects error to be %s but has %s", t.Name(), msg, err.Error())
	}

	if data, ok := err.Data().(error); !ok {
		t.Fatalf("%s expects data to be of type error but has %T", t.Name(), data)
	}
}

func TestErrPanic_DataStringer(t *testing.T) {
	tpl := "STRING"
	msg := fmt.Sprintf("%s%s", errorPrefix, tpl)

	err := NewErrPanic(Str(tpl))
	if err.Error() != msg {
		t.Fatalf("%s expects error to be %s but has %s", t.Name(), msg, err.Error())
	}

	if data, ok := err.Data().(fmt.Stringer); !ok {
		t.Fatalf("%s expects data to be of type error but has %T", t.Name(), data)
	}
}

func TestErrPanic_DataOther(t *testing.T) {
	var tpl int = 42
	msg := fmt.Sprintf("%s%d", errorPrefix, tpl)

	err := NewErrPanic(tpl)
	if err.Error() != msg {
		t.Fatalf("%s expects error to be %s but has %s", t.Name(), msg, err.Error())
	}

	if data, ok := err.Data().(int); !ok {
		t.Fatalf("%s expects data to be of type error but has %T", t.Name(), data)
	}
}
