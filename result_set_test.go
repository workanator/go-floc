package floc

import "testing"

func TestResultSet_Contains(t *testing.T) {
	set := NewResultSet(None, Completed)

	if set.Contains(None) == false {
		t.Fatalf("%s expects None to be in set", t.Name())
	} else if set.Contains(Completed) == false {
		t.Fatalf("%s expects Completed to be in set", t.Name())
	} else if set.Contains(Canceled) == true {
		t.Fatalf("%s expects Canceled to be not in set", t.Name())
	} else if set.Contains(Failed) == true {
		t.Fatalf("%s expects Failed to be not in set", t.Name())
	}
}

func TestResultSet_IsEmpty(t *testing.T) {
	emptySet := NewResultSet()
	if !emptySet.IsEmpty() {
		t.Fatalf("%s expects set to be empty", t.Name())
	}

	nonEmptySet := NewResultSet(None)
	if nonEmptySet.IsEmpty() {
		t.Fatalf("%s expects set to be not empty", t.Name())
	}
}

func TestResultSet_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%s must panic", t.Name())
		}
	}()

	NewResultSet(None, Completed, Canceled, Failed, Result(100))
}
