package guard

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/state"
)

func TestCancel(t *testing.T) {
	const tpl = "canceled"

	f := flow.New()
	s := state.New(nil)
	job := Cancel(tpl)

	floc.Run(f, s, nil, job)

	result, data := f.Result()
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled, result)
	}

	if data.(string) != tpl {
		t.Fatalf("%s expects data to be string '%s' but has %v", t.Name(), tpl, data)
	}
}
