package guard

import (
	"math/rand"
	"testing"
	"time"
)

func TestConstTimeout(t *testing.T) {
	const max = 100

	for i := 0; i < max; i++ {
		tpl := time.Duration(rand.Int()) * time.Millisecond
		constFunc := ConstTimeout(tpl)
		result := constFunc(nil, nil)

		if result != tpl {
			t.Fatalf("%s:%d failed with template %v and result %v", t.Name(), i, tpl, result)
		}
	}
}
