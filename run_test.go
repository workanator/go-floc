package floc

import "testing"

func TestRun(t *testing.T) {
	finished := false

	Run(nil, nil, nil, func(flow Flow, state State, update Update) {
		finished = true
	})

	if finished == false {
		t.Fatalf("%s expects true", t.Name())
	}
}
