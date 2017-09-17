package errors

import (
	"fmt"
	"testing"
)

func TestErrLocation_What(t *testing.T) {
	const max = 100

	for i := 0; i < max; i++ {
		what := fmt.Errorf("%d", i)
		err := NewErrLocation(what, "")
		if err.What().Error() != what.Error() {
			t.Fatalf("%s expects What to be %s but has %s", t.Name(), what.Error(), err.What().Error())
		}
	}
}

func TestErrLocation_Where(t *testing.T) {
	const max = 100

	for i := 0; i < max; i++ {
		where := fmt.Sprintf("%d", i)
		err := NewErrLocation(nil, where)
		if err.Where() != where {
			t.Fatalf("%s expects Where to be %s but has %s", t.Name(), where, err.Where())
		}
	}
}

func TestErrLocation_Error(t *testing.T) {
	const max = 100

	for i := 0; i < max; i++ {
		what := fmt.Errorf("%d", i)
		where := fmt.Sprintf("%d", max-i)
		msg := where + locationDelimiter + what.Error()

		err := NewErrLocation(what, where)
		if err.Error() != msg {
			t.Fatalf("%s expects error to be %s but has %s", t.Name(), msg, err.Error())
		}
	}
}
