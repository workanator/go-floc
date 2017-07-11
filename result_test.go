package floc

import (
	"fmt"
	"testing"
)

func TestResultNone(t *testing.T) {
	r := None

	if r.IsNone() == false {
		t.Fatalf("%s expects None but has %s", t.Name(), r)
	}

	if r.IsValid() == false {
		t.Fatalf("%s expects None to be valid", t.Name())
	}

	if r.String() != "None" {
		t.Fatalf("%s expects None", t.Name())
	}
}

func TestResultCanceled(t *testing.T) {
	r := Canceled

	if r.IsCanceled() == false {
		t.Fatalf("%s expects Canceled but has %s", t.Name(), r)
	}

	if r.IsValid() == false {
		t.Fatalf("%s expects Canceled to be valid", t.Name())
	}

	if r.String() != "Canceled" {
		t.Fatalf("%s expects Canceled", t.Name())
	}
}

func TestResultCompleted(t *testing.T) {
	r := Completed

	if r.IsCompleted() == false {
		t.Fatalf("%s expects Completed but has %s", t.Name(), r)
	}

	if r.IsValid() == false {
		t.Fatalf("%s expects Completed to be valid", t.Name())
	}

	if r.String() != "Completed" {
		t.Fatalf("%s expects Completed", t.Name())
	}
}

func TestResultInvalid(t *testing.T) {
	const n = 100

	r := Result(n)

	if r.IsValid() == true {
		t.Fatalf("%s expects %s to be invalid", t.Name(), r)
	}

	s := fmt.Sprintf("Result(%d)", n)
	if r.String() != s {
		t.Fatalf("%s expects %s but has %s", t.Name(), s, r)
	}
}
