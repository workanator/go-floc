package floc

import (
	"context"
	"sync"
	"testing"
)

func TestBorrowContext(t *testing.T) {
	ctx := BorrowContext(context.Background())
	ctx.Release()
}

func TestBorrowContext_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%s expects panic", t.Name())
		}
	}()

	BorrowContext(nil)
}

func TestNewContext(t *testing.T) {
	ctx := NewContext()
	ctx.Release()
}

func TestFlowContext_Ctx(t *testing.T) {
	type customKeyType int

	const (
		key   customKeyType = 1
		value               = "VALUE"
	)

	ctx := NewContext()
	defer ctx.Release()

	innerCtx := ctx.Ctx()
	ctx.UpdateCtx(context.WithValue(innerCtx, key, value))

	v := ctx.Ctx().Value(key)
	if v == nil {
		t.Fatalf("%s expects value to be not nil", t.Name())
	}

	if s, ok := v.(string); !ok {
		t.Fatalf("%s expects value to be of type string but has %T", t.Name(), v)
	} else if s != value {
		t.Fatalf("%s expects value to be %s but has %s", t.Name(), value, s)
	}
}

func TestFlowContext_Value(t *testing.T) {
	const max = 1000

	ctx := NewContext()
	defer ctx.Release()

	var wg sync.WaitGroup
	var n int
	for n = 0; n < max; n++ {
		wg.Add(1)

		go func(n int) {
			ctx.AddValue(n, n)
			wg.Done()
		}(n)
	}

	wg.Wait()

	for n = 0; n < max; n++ {
		v := ctx.Value(n)
		if v == nil {
			t.Fatalf("%s expects value to be not nil", t.Name())
		}

		if d, ok := v.(int); !ok {
			t.Fatalf("%s expects value %d to be of type int but has %T", t.Name(), n, v)
		} else if d != n {
			t.Fatalf("%s expects value %d to be %d but has %d", t.Name(), n, n, d)
		}
	}
}

func TestFlowContext_Done(t *testing.T) {
	ctx := NewContext()
	defer ctx.Release()

	cancelCtx, cancel := context.WithCancel(ctx.Ctx())
	ctx.UpdateCtx(cancelCtx)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			<-ctx.Done()
			wg.Done()
		}()
	}

	cancel()

	wg.Wait()
}
