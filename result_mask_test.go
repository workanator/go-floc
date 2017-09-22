package floc

import "testing"

func TestEmptyResultMask(t *testing.T) {
	set := EmptyResultMask()

	for _, r := range []Result{None, Completed, Canceled, Failed} {
		go func(r Result) {
			if set.IsMasked(r) {
				t.Fatalf("%s expects set to not contain %s", t.Name(), r.String())
			}
		}(r)
	}
}

func TestResultMask_IsMasked(t *testing.T) {
	set := NewResultMask(None | Completed)

	if set.IsMasked(None) == false {
		t.Fatalf("%s expects None to be in set", t.Name())
	} else if set.IsMasked(Completed) == false {
		t.Fatalf("%s expects Completed to be in set", t.Name())
	} else if set.IsMasked(Canceled) == true {
		t.Fatalf("%s expects Canceled to be not in set", t.Name())
	} else if set.IsMasked(Failed) == true {
		t.Fatalf("%s expects Failed to be not in set", t.Name())
	}
}

func TestResultMask_IsEmpty(t *testing.T) {
	emptySet := EmptyResultMask()
	if !emptySet.IsEmpty() {
		t.Fatalf("%s expects set to be empty", t.Name())
	}

	nonEmptySet := NewResultMask(None)
	if nonEmptySet.IsEmpty() {
		t.Fatalf("%s expects set to be not empty", t.Name())
	}
}
