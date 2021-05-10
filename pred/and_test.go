package pred

import (
	"testing"

	"github.com/workanator/go-floc/v3"
)

func TestAnd_True(t *testing.T) {
	const max = 100

	for i := 2; i < max; i++ {
		tests := make([]floc.Predicate, i)
		for n := 0; n < i; n++ {
			tests[n] = alwaysTrue
		}

		p := And(tests...)

		if p(nil) == false {
			t.Fatalf("%s expects true with %d tests", t.Name(), i)
		}
	}
}

func TestAnd_False(t *testing.T) {
	const max = 100

	for i := 2; i < max; i++ {
		tests := make([]floc.Predicate, i)
		for n := 0; n < i; n++ {
			tests[n] = alwaysFalse
		}

		p := And(tests...)

		if p(nil) == true {
			t.Fatalf("%s expects false with %d tests", t.Name(), i)
		}
	}
}

func TestAnd_Mixed(t *testing.T) {
	const max = 100

	for i := 2; i < max; i++ {
		tests := make([]floc.Predicate, i)
		for n := 0; n < i; n++ {
			if n%2 == 0 {
				tests[n] = alwaysTrue
			} else {
				tests[n] = alwaysFalse
			}
		}

		p := And(tests...)

		if p(nil) == true {
			t.Fatalf("%s expects false with %d tests", t.Name(), i)
		}
	}
}

func TestAnd_Panic(t *testing.T) {
	panicFunc := func(n int) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("%s must panic with %d tests", t.Name(), n)
			}
		}()

		tests := make([]floc.Predicate, n)
		for i := 0; i < n; i++ {
			tests[n] = alwaysFalse
		}

		And(tests...)
	}

	panicFunc(0)
	panicFunc(1)
	panicFunc(2)
}
