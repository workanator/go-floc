package pred

import "testing"

func TestNotTrue(t *testing.T) {
	p := Not(alwaysTrue)

	if p(nil) == true {
		t.Fatalf("%s expects false", t.Name())
	}
}

func TestNotFalse(t *testing.T) {
	p := Not(alwaysFalse)

	if p(nil) == false {
		t.Fatalf("%s expects true", t.Name())
	}
}
