package pred

import "testing"

func TestNot_True(t *testing.T) {
	p := Not(alwaysTrue)

	if p(nil) == true {
		t.Fatalf("%s expects false", t.Name())
	}
}

func TestNot_False(t *testing.T) {
	p := Not(alwaysFalse)

	if p(nil) == false {
		t.Fatalf("%s expects true", t.Name())
	}
}
