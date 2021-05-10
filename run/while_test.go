package run

import (
	"testing"

	"fmt"

	"github.com/workanator/go-floc/v3"
)

func TestWhile_AlreadyFinished(t *testing.T) {
	ctx := floc.NewContext()
	defer ctx.Release()

	ctrl := floc.NewControl(ctx)
	defer ctrl.Release()

	flow := While(yes(), cancel(nil))

	ctrl.Complete(nil)

	result, _, _ := floc.RunWith(ctx, ctrl, flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects flow to be Completed but has %s", t.Name(), result.String())
	}
}

func TestWhile_None(t *testing.T) {
	const start = 100

	flow := While(countdown(start), noop())
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestWhile_Completed(t *testing.T) {
	flow := While(yes(), complete(nil))
	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestWhile_Canceled(t *testing.T) {
	flow := While(yes(), cancel(nil))
	result, data, err := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestWhile_Failed(t *testing.T) {
	flow := While(yes(), fail(nil, fmt.Errorf("err")))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestWhile_Error(t *testing.T) {
	flow := While(yes(), throw(fmt.Errorf("err")))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}
