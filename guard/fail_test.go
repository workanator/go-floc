package guard

import (
	"testing"

	"fmt"

	"gopkg.in/workanator/go-floc.v2"
)

func TestFail(t *testing.T) {
	const dataTpl = "failed"
	var errTpl = fmt.Errorf("err %s", t.Name())

	flow := Fail(dataTpl, errTpl)

	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data == nil {
		t.Fatalf("%s expects data to be not nil", t.Name())
	} else if s, ok := data.(string); !ok {
		t.Fatalf("%s expects data to be of type string but has %v", t.Name(), s)
	} else if s != dataTpl {
		t.Fatalf("%s expects data to be %s but has %s", t.Name(), dataTpl, s)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	} else if err.Error() != errTpl.Error() {
		t.Fatalf("%s expects error to be %s but has %s", t.Name(), errTpl.Error(), err.Error())
	}
}
