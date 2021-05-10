package pred

import (
	"testing"

	"github.com/workanator/go-floc/v3"
)

func TestXor_True(t *testing.T) {
	const max = 100

	for i := 2; i < max; i++ {
		tests := make([]floc.Predicate, i)
		for n := 0; n < i; n++ {
			tests[n] = alwaysTrue
		}

		p := Xor(tests...)

		// Values: 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 ...
		// Result:   0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 ...
		mustBe := i%2 != 0
		result := p(nil)

		if result != mustBe {
			t.Fatalf("%s expects %t with %d tests but has %t", t.Name(), mustBe, i, result)
		}
	}
}

func TestXor_False(t *testing.T) {
	const max = 100

	for i := 2; i < max; i++ {
		tests := make([]floc.Predicate, i)
		for n := 0; n < i; n++ {
			tests[n] = alwaysFalse
		}

		p := Xor(tests...)

		// Values: 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 ...
		// Result:   0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 ...
		result := p(nil)
		if result != false {
			t.Fatalf("%s expects false with %d tests but has %t", t.Name(), i, result)
		}
	}
}

func TestXor_Mixed(t *testing.T) {
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

		p := Xor(tests...)

		// Values: 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 ...
		// Result:   1 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 ...
		mustBe := (i%4 == 1) || (i%4 == 2)
		result := p(nil)

		if result != mustBe {
			t.Fatalf("%s expects %t with %d tests but has %t", t.Name(), mustBe, i, result)
		}
	}
}

func TestXor_Panic(t *testing.T) {
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

		Xor(tests...)
	}

	panicFunc(0)
	panicFunc(1)
	panicFunc(2)
}
