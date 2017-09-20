package guard

import (
	"testing"

	"gopkg.in/workanator/go-floc.v2"
)

func TestCancel(t *testing.T) {
	const tpl = "canceled"

	flow := Cancel(tpl)

	result, data, err := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	} else if data == nil {
		t.Fatalf("%s expects data to be not nil", t.Name())
	} else if s, ok := data.(string); !ok {
		t.Fatalf("%s expects data to be of type string but has %v", t.Name(), s)
	} else if s != tpl {
		t.Fatalf("%s expects data to be %s but has %s", t.Name(), tpl, s)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}
