package guard

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeoutError(t *testing.T) {
	id := 1
	now := time.Now()
	err := ErrTimeout{ID: id, At: now}

	manual := fmt.Sprintf("%v timed out at %s", id, now)
	formatted := err.Error()
	if manual != formatted {
		t.Fatalf("%s expects output to be '%s' but has '%s'", t.Name(), manual, formatted)
	}
}
