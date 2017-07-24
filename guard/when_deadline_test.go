package guard

import (
	"math/rand"
	"testing"
	"time"
)

func TestConstDeadline(t *testing.T) {
	const iterations = 100

	for i := 0; i < iterations; i++ {
		tpl := time.Now().Add(time.Duration(rand.Int()) * time.Second)
		constFunc := ConstDeadline(tpl)
		result := constFunc(nil, nil)

		if result != tpl {
			t.Fatalf("%s failed on iteration %d with template %v and result %v", t.Name(), i, tpl, result)
		}
	}
}

func TestDeadlineIn(t *testing.T) {
	const iterations = 100

	for i := 0; i < iterations; i++ {
		tpl := time.Duration(rand.Int()) * time.Second
		constFunc := DeadlineIn(tpl)
		result := constFunc(nil, nil)

		when := time.Now().Add(tpl)
		if when.Before(result) {
			t.Fatalf("%s failed on iteration %d with template %v and result %v", t.Name(), i, tpl, result)
		}
	}
}
