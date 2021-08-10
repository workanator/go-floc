package floc

import (
	"testing"
)

func TestEmptyResultMask(t *testing.T) {
	set := EmptyResultMask()
	results := []Result{None, Completed, Canceled, Failed}
	length := len(results)
	maskeds := make(chan int, length)
	for idx, r := range results {
		go func(n int, r Result) {
			if set.IsMasked(r) {
				maskeds <- n
			} else {
				maskeds <- -1
			}
		}(idx, r)
	}

	for i := 0; i < length; i++ {
		index := <-maskeds
		if index >= 0 {
			t.Fatalf("%s expects %s to be not masked", t.Name(), results[index].String())
		}
	}
}

func TestResultMask_IsMasked(t *testing.T) {
	set := NewResultMask(None | Completed)

	if set.IsMasked(None) == false {
		t.Fatalf("%s expects None to be masked", t.Name())
	} else if set.IsMasked(Completed) == false {
		t.Fatalf("%s expects Completed to be masked", t.Name())
	} else if set.IsMasked(Canceled) == true {
		t.Fatalf("%s expects Canceled to be not masked", t.Name())
	} else if set.IsMasked(Failed) == true {
		t.Fatalf("%s expects Failed to be not masked", t.Name())
	}
}

func TestResultMask_IsEmpty(t *testing.T) {
	emptyMask := EmptyResultMask()
	if !emptyMask.IsEmpty() {
		t.Fatalf("%s expects mask to be empty", t.Name())
	}

	nonEmptyMask := NewResultMask(None)
	if nonEmptyMask.IsEmpty() {
		t.Fatalf("%s expects mask to be not empty", t.Name())
	}
}

func TestResultMask_String(t *testing.T) {
	emptyMask := EmptyResultMask()
	if emptyMask.String() != "[]" {
		t.Fatalf("%s expects mask to be [] but has %s", t.Name(), emptyMask.String())
	}

	mask1 := NewResultMask(None)
	if mask1.String() != "[None]" {
		t.Fatalf("%s expects mask to be [None] but has %s", t.Name(), mask1.String())
	}

	mask2 := NewResultMask(Canceled | Failed)
	if mask2.String() != "[Canceled,Failed]" {
		t.Fatalf("%s expects mask to be [Canceled,Failed] but has %s", t.Name(), mask2.String())
	}

	mask3 := NewResultMask(Canceled | Failed | None)
	if mask3.String() != "[None,Canceled,Failed]" {
		t.Fatalf("%s expects mask to be [None,Canceled,Failed] but has %s", t.Name(), mask3.String())
	}

	mask4 := NewResultMask(Canceled | Completed | Failed | None)
	if mask4.String() != "[None,Completed,Canceled,Failed]" {
		t.Fatalf("%s expects mask to be [None,Completed,Canceled,Failed] but has %s", t.Name(), mask4.String())
	}
}
