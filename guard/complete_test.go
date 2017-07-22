package guard

import (
	"testing"

	floc "github.com/workanator/go-floc"
)

func TestComplete(t *testing.T) {
	const tpl = "completed"

	f := floc.NewFlow()
	s := floc.NewState(nil)
	job := Complete(tpl)

	floc.Run(f, s, nil, job)

	result, data := f.Result()
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed.String(), result)
	}

	if data.(string) != tpl {
		t.Fatalf("%s expects data to be string '%s' but has %v", t.Name(), tpl, data)
	}
}
