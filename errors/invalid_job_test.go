package errors

import "testing"

func TestErrInvalidJob_Error(t *testing.T) {
	err := ErrInvalidJob{}
	if err.Error() != invalidJobMessage {
		t.Fatalf("%s expects message to be %s but has %s", t.Name(), invalidJobMessage, err.Error())
	}
}
