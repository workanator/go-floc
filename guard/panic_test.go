package guard

import (
	"testing"

	"github.com/workanator/go-floc"
	"github.com/workanator/go-floc/errors"
)

func TestPanic(t *testing.T) {
	const tpl int = 1

	flow := Panic(panicWith(tpl))

	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data == nil {
		t.Fatalf("%s expects data to be not nil", t.Name())
	} else if v, ok := data.(int); !ok {
		t.Fatalf("%s expects data to be of type int but has %T", t.Name(), data)
	} else if v != tpl {
		t.Fatalf("%s expects data to be %d but has %d", t.Name(), tpl, v)
	} else if e, ok := err.(errors.ErrPanic); !ok {
		t.Fatalf("%s expects error to be ErrPanic but has %T", t.Name(), e)
	} else if e.Data() == nil {
		t.Fatalf("%s expects error data to be not nil", t.Name())
	} else if d, ok := e.Data().(int); !ok {
		t.Fatalf("%s expects data to be of type int but has %T", t.Name(), e.Data())
	} else if d != tpl {
		t.Fatalf("%s expects data to be %d but has %d", t.Name(), tpl, d)
	}

}

func TestIgnorePanic(t *testing.T) {
	const tpl int = 1

	flow := IgnorePanic(panicWith(tpl))

	result, _, _ := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	}
}

func TestOnPanic(t *testing.T) {
	const tpl int = 2

	flow := OnPanic(
		panicWith(tpl),
		func(ctx floc.Context, ctrl floc.Control, v interface{}) {
			ctrl.Complete(tpl)
		},
	)

	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	} else if d, ok := data.(int); !ok {
		t.Fatalf("%s expects data to be of type int but has %T", t.Name(), data)
	} else if d != tpl {
		t.Fatalf("%s expects data to be %d but has %d", t.Name(), tpl, d)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func panicWith(v interface{}) floc.Job {
	return func(floc.Context, floc.Control) error {
		panic(v)
	}
}
