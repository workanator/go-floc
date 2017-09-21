package floc

import (
	"fmt"
	"testing"
)

func TestResult_IsNone(t *testing.T) {
	r := None
	if r.IsNone() == false {
		t.Fatalf("%s expects None but has %s", t.Name(), r.String())
	} else if r.IsValid() == false {
		t.Fatalf("%s expects None to be valid", t.Name())
	}
}

func TestResult_IsCanceled(t *testing.T) {
	r := Canceled
	if r.IsCanceled() == false {
		t.Fatalf("%s expects Canceled but has %s", t.Name(), r.String())
	} else if r.IsValid() == false {
		t.Fatalf("%s expects Canceled to be valid", t.Name())
	}
}

func TestResult_IsCompleted(t *testing.T) {
	r := Completed
	if r.IsCompleted() == false {
		t.Fatalf("%s expects Completed but has %s", t.Name(), r.String())
	} else if r.IsValid() == false {
		t.Fatalf("%s expects Completed to be valid", t.Name())
	}
}

func TestResult_IsFailed(t *testing.T) {
	r := Failed
	if r.IsFailed() == false {
		t.Fatalf("%s expects Failed but has %s", t.Name(), r.String())
	} else if r.IsValid() == false {
		t.Fatalf("%s expects Failed to be valid", t.Name())
	}
}

func TestResult_IsValid(t *testing.T) {
	const n = 1000

	r := Result(n)
	if r.IsValid() == true {
		t.Fatalf("%s expects %s to be invalid", t.Name(), r.String())
	}

	s := fmt.Sprintf("Result(%d)", n)
	if r.String() != s {
		t.Fatalf("%s expects %s but has %s", t.Name(), s, r.String())
	}
}

func TestResult_IsFinished(t *testing.T) {
	if None.IsFinished() == true {
		t.Fatalf("%s expects None to be not finished", t.Name())
	}

	if Completed.IsFinished() == false {
		t.Fatalf("%s expects Completed to be finished", t.Name())
	}

	if Canceled.IsFinished() == false {
		t.Fatalf("%s expects Canceled to be finished", t.Name())
	}

	if Failed.IsFinished() == false {
		t.Fatalf("%s expects Failed to be finished", t.Name())
	}
}

func TestResult_Int32(t *testing.T) {
	var n int32
	for n = 0; n < 1000; n++ {
		r := Result(n)
		if r.i32() != n {
			t.Fatalf("%s expects Result to be %d but has %d", t.Name(), n, r.i32())
		}
	}
}

func TestResult_String(t *testing.T) {
	if None.String() != "None" {
		t.Fatalf("%s expects None bu has %s", t.Name(), None.String())
	}

	if Completed.String() != "Completed" {
		t.Fatalf("%s expects Completed bu has %s", t.Name(), Completed.String())
	}

	if Canceled.String() != "Canceled" {
		t.Fatalf("%s expects Canceled bu has %s", t.Name(), Canceled.String())
	}

	if Failed.String() != "Failed" {
		t.Fatalf("%s expects Failed bu has %s", t.Name(), Failed.String())
	}
}
