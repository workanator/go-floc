package errors

import (
	"fmt"
	"testing"
	"time"
)

func TestErrTimeout_ID(t *testing.T) {
	const max int = 100

	for i := 0; i < max; i++ {
		err := NewErrTimeout(i, time.Now())
		if err.ID() == nil {
			t.Fatalf("%s expects ID to be not nil", t.Name())
		} else if id, ok := err.ID().(int); !ok {
			t.Fatalf("%s expects ID to be of type int but has %v", t.Name(), id)
		} else if id != i {
			t.Fatalf("%se expects ID to be %d but has %d", t.Name(), i, id)
		}
	}
}

func TestErrTimeout_At(t *testing.T) {
	const max int = 100

	for i := 0; i < max; i++ {
		now := time.Now()
		err := NewErrTimeout(nil, now)

		if !err.At().Equal(now) {
			t.Fatalf("%s expects time to be %s but has %s", t.Name(), now.Format(time.RFC3339), err.At().Format(time.RFC3339))
		}
	}
}

func TestErrTimeout_Error(t *testing.T) {
	id := 1
	now := time.Now()
	msg := fmt.Sprintf(tplTimeoutMessage, id, now)

	err := NewErrTimeout(id, now)
	if err.Error() != msg {
		t.Fatalf("%s expects error to be %s but has %s", t.Name(), msg, err.Error())
	}
}
