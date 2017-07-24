package guard

import (
	"math/rand"
	"testing"
	"time"
)

func TestConstTimeout(t *testing.T) {
	const iterations = 100

	for i := 0; i < iterations; i++ {
		tpl := time.Duration(rand.Int()) * time.Millisecond
		constFunc := ConstTimeout(tpl)
		result := constFunc(nil, nil)

		if result != tpl {
			t.Fatalf("%s failed on iteration %d with template %v and result %v", t.Name(), i, tpl, result)
		}
	}
}
