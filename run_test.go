package floc

import (
	"fmt"
	"testing"
)

func TestRun_NilJob(t *testing.T) {
	result, data, err := Run(nil)

	if result.IsNone() == false {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	}

	if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	}

	if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestRun_None(t *testing.T) {
	flow := func(Context, Control) error {
		return nil
	}

	result, data, err := Run(flow)

	if result.IsNone() == false {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	}

	if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	}

	if err != nil {
		t.Fatalf("%s expects error to be nil but has %s", t.Name(), err.Error())
	}
}

func TestRun_Completed(t *testing.T) {
	const tpl int = 1

	flow := func(ctx Context, ctrl Control) error {
		ctrl.Complete(tpl)
		return nil
	}

	result, data, err := Run(flow)

	if result.IsCompleted() == false {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	}

	if data == nil {
		t.Fatalf("%s expects data to be not nil", t.Name())
	} else if d, ok := data.(int); !ok {
		t.Fatalf("%s expects data to be of type int but has %T", t.Name(), data)
	} else if d != tpl {
		t.Fatalf("%s expects data to be %d but has %d", t.Name(), tpl, d)
	}

	if err != nil {
		t.Fatalf("%s expects error to be nil but has %s", t.Name(), err.Error())
	}
}

func TestRun_Canceled(t *testing.T) {
	const tpl float64 = 3.1415

	flow := func(ctx Context, ctrl Control) error {
		ctrl.Cancel(tpl)
		return nil
	}

	result, data, err := Run(flow)

	if result.IsCanceled() == false {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	}

	if data == nil {
		t.Fatalf("%s expects data to be not nil", t.Name())
	} else if f, ok := data.(float64); !ok {
		t.Fatalf("%s expects data to be of type float64 but has %T", t.Name(), data)
	} else if f < tpl-0.0001 || f > tpl+0.0001 {
		t.Fatalf("%s expects data to be %f but has %f", t.Name(), tpl, f)
	}

	if err != nil {
		t.Fatalf("%s expects error to be nil but has %s", t.Name(), err.Error())
	}
}

func TestRun_Failed(t *testing.T) {
	const tplData string = "REASON"
	var tplError error = fmt.Errorf("failed because of %s", tplData)

	flow := func(ctx Context, ctrl Control) error {
		ctrl.Fail(tplData, tplError)
		return nil
	}

	result, data, err := Run(flow)

	if result.IsFailed() == false {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	}

	if data == nil {
		t.Fatalf("%s expects data to be not nil", t.Name())
	} else if s, ok := data.(string); !ok {
		t.Fatalf("%s expects data to be of type string but has %T", t.Name(), data)
	} else if s != tplData {
		t.Fatalf("%s expects data to be %s but has %s", t.Name(), tplData, s)
	}

	if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	} else if err.Error() != tplError.Error() {
		t.Fatalf("%s expects error message to be %s but has %s", t.Name(), tplError.Error(), err.Error())
	}
}

func TestRun_UnhandledError(t *testing.T) {
	var tplError error = fmt.Errorf("something happened")

	flow := func(ctx Context, ctrl Control) error {
		return tplError
	}

	result, data, err := Run(flow)

	if result.IsFailed() == false {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	}

	if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	}

	if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	} else if err.Error() != tplError.Error() {
		t.Fatalf("%s expects error message to be %s but has %s", t.Name(), tplError.Error(), err.Error())
	}
}
