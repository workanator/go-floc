package guard

import (
	"testing"

	"github.com/workanator/go-floc/v3"
)

func TestComplete(t *testing.T) {
	const tpl = "completed"

	flow := Complete(tpl)

	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	} else if data == nil {
		t.Fatalf("%s expects data to be not nil", t.Name())
	} else if s, ok := data.(string); !ok {
		t.Fatalf("%s expects data to be of type string but has %T", t.Name(), data)
	} else if s != tpl {
		t.Fatalf("%s expects data to be %s but has %s", t.Name(), tpl, s)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}
