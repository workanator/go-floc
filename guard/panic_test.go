package guard

import (
	"testing"

	floc "github.com/workanator/go-floc"
)

func TestPanic(t *testing.T) {
	const tpl int = 1

	f := floc.NewFlow()
	s := floc.NewStateContainer(nil)
	job := Panic(func(floc.Flow, floc.State, floc.Update) {
		panic(tpl)
	})

	floc.Run(f, s, nil, job)

	result, data := f.Result()
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled.String(), result)
	}

	e, ok := data.(ErrPanic)
	if !ok {
		t.Fatalf("%s expects data to be ErrPanic but has %T", t.Name(), data)
	}

	d, ok := e.Data.(int)
	if !ok {
		t.Fatalf("%s expects error data to be int but has %T", t.Name(), e.Data)
	}

	if d != tpl {
		t.Fatalf("%s expects error data to be %d but has %d", t.Name(), tpl, d)
	}
}

func TestPanicIgnore(t *testing.T) {
	const tpl int = 1

	f := floc.NewFlow()
	s := floc.NewStateContainer(nil)
	job := IgnorePanic(func(floc.Flow, floc.State, floc.Update) {
		panic(tpl)
	})

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsNone() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.None.String(), result)
	}
}

func TestPanicWithTrigger(t *testing.T) {
	const tpl int = 2

	f := floc.NewFlow()
	s := floc.NewStateContainer(nil)
	job := PanicWithTrigger(
		func(floc.Flow, floc.State, floc.Update) {
			panic(tpl)
		},
		func(flow floc.Flow, state floc.State, v interface{}) {
			flow.Complete(tpl)
		},
	)

	floc.Run(f, s, nil, job)

	result, data := f.Result()
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed.String(), result)
	}

	d, ok := data.(int)
	if !ok {
		t.Fatalf("%s expects data to be int but has %T", t.Name(), data)
	}

	if d != tpl {
		t.Fatalf("%s expects error data to be %d but has %d", t.Name(), tpl, d)
	}
}
