package pred

import (
	"testing"

	floc "github.com/workanator/go-floc.v1"
)

func TestXorAllTrue(t *testing.T) {
	const max = 100

	for i := 2; i < max; i++ {
		predicates := make([]floc.Predicate, i)
		for n := 0; n < i; n++ {
			predicates[n] = alwaysTrue
		}

		p := Xor(predicates...)

		// Values: 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 ...
		// Result:   0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 ...
		mustBe := i%2 != 0
		result := p(nil)

		if result != mustBe {
			t.Fatalf("%s expects %t with %d predicates but has %t", t.Name(), mustBe, i, result)
		}
	}
}

func TestXorAllFalse(t *testing.T) {
	const max = 100

	for i := 2; i < max; i++ {
		predicates := make([]floc.Predicate, i)
		for n := 0; n < i; n++ {
			predicates[n] = alwaysFalse
		}

		p := Xor(predicates...)

		// Values: 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 ...
		// Result:   0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 ...
		mustBe := false
		result := p(nil)

		if result != mustBe {
			t.Fatalf("%s expects %t with %d predicates but has %t", t.Name(), mustBe, i, result)
		}
	}
}

func TestXorMixed(t *testing.T) {
	const max = 100

	for i := 2; i < max; i++ {
		predicates := make([]floc.Predicate, i)
		for n := 0; n < i; n++ {
			if n%2 == 0 {
				predicates[n] = alwaysTrue
			} else {
				predicates[n] = alwaysFalse
			}
		}

		p := Xor(predicates...)

		// Values: 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 ...
		// Result:   1 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 ...
		mustBe := (i%4 == 1) || (i%4 == 2)
		result := p(nil)

		if result != mustBe {
			t.Fatalf("%s expects %t with %d predicates but has %t", t.Name(), mustBe, i, result)
		}
	}
}

func TestXorPanic(t *testing.T) {
	panicFunc := func(n int) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("%s must panic with %d predicates", t.Name(), n)
			}
		}()

		predicates := make([]floc.Predicate, n)
		for i := 0; i < n; i++ {
			predicates[n] = alwaysFalse
		}

		Xor(predicates...)
	}

	panicFunc(0)
	panicFunc(1)
	panicFunc(2)
}
