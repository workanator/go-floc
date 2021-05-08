package run

import (
	"testing"

	"fmt"

	"github.com/workanator/go-floc"
)

func TestBackground_AlreadyFinished(t *testing.T) {
	ctx := floc.NewContext()
	defer ctx.Release()

	ctrl := floc.NewControl(ctx)
	defer ctrl.Release()

	flow := Background(complete(nil))

	ctrl.Cancel(nil)

	result, _, _ := floc.RunWith(ctx, ctrl, flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	}
}

func TestBackground_Completed(t *testing.T) {
	flow := Sequence(Background(complete(nil)), waitUntilFinished())
	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestBackground_Canceled(t *testing.T) {
	flow := Sequence(Background(cancel(nil)), waitUntilFinished())
	result, data, err := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestBackground_Failed(t *testing.T) {
	flow := Sequence(Background(fail(nil, fmt.Errorf("err"))), waitUntilFinished())
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestBackground_Error(t *testing.T) {
	flow := Sequence(Background(throw(fmt.Errorf("err"))), waitUntilFinished())
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}
