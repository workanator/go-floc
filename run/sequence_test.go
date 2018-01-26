package run

import (
	"testing"

	"fmt"

	"gopkg.in/devishot/go-floc.v2"
)

func TestSequence_AlreadyFinished(t *testing.T) {
	ctx := floc.NewContext()
	defer ctx.Release()

	ctrl := floc.NewControl(ctx)
	defer ctrl.Release()

	flow := Sequence(noop())

	ctrl.Complete(nil)

	result, _, _ := floc.RunWith(ctx, ctrl, flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects flow to be Completed but has %s", t.Name(), result.String())
	}
}

func TestSequence_None(t *testing.T) {
	const max = 100

	for n := 1; n <= max; n++ {
		go func(t *testing.T, n int) {
			jobs := make([]floc.Job, n)
			for i := 0; i < n; i++ {
				jobs[i] = noop()
			}

			flow := Sequence(jobs...)
			result, data, err := floc.Run(flow)
			if !result.IsNone() {
				t.Fatalf("%s:%d expects result to be None but has %s", t.Name(), n, result.String())
			} else if data != nil {
				t.Fatalf("%s:%d expects data to be nil but has %v", t.Name(), n, data)
			} else if err != nil {
				t.Fatalf("%s:%d expects error to be nil but has %v", t.Name(), n, err)
			}
		}(t, n)
	}
}

func TestSequence_Completed(t *testing.T) {
	const max = 100

	for n := 1; n <= max; n++ {
		go func(t *testing.T, n int) {
			jobs := make([]floc.Job, n)
			for i := 0; i < n; i++ {
				if i == n-1 {
					jobs[i] = complete(nil)
				} else {
					jobs[i] = noop()
				}
			}

			flow := Sequence(jobs...)
			result, data, err := floc.Run(flow)
			if !result.IsCompleted() {
				t.Fatalf("%s:%d expects result to be Completed but has %s", t.Name(), n, result.String())
			} else if data != nil {
				t.Fatalf("%s:%d expects data to be nil but has %v", t.Name(), n, data)
			} else if err != nil {
				t.Fatalf("%s:%d expects error to be nil but has %v", t.Name(), n, err)
			}
		}(t, n)
	}
}

func TestSequence_Canceled(t *testing.T) {
	const max = 100

	for n := 1; n <= max; n++ {
		go func(t *testing.T, n int) {
			jobs := make([]floc.Job, n)
			for i := 0; i < n; i++ {
				if i == n-1 {
					jobs[i] = cancel(nil)
				} else {
					jobs[i] = noop()
				}
			}

			flow := Sequence(jobs...)
			result, data, err := floc.Run(flow)
			if !result.IsCanceled() {
				t.Fatalf("%s:%d expects result to be Canceled but has %s", t.Name(), n, result.String())
			} else if data != nil {
				t.Fatalf("%s:%d expects data to be nil but has %v", t.Name(), n, data)
			} else if err != nil {
				t.Fatalf("%s:%d expects error to be nil but has %v", t.Name(), n, err)
			}
		}(t, n)
	}
}

func TestSequence_Failed(t *testing.T) {
	const max = 100

	for n := 1; n <= max; n++ {
		go func(t *testing.T, n int) {
			jobs := make([]floc.Job, n)
			for i := 0; i < n; i++ {
				if i == n-1 {
					jobs[i] = fail(nil, fmt.Errorf("err%d", n))
				} else {
					jobs[i] = noop()
				}
			}

			flow := Sequence(jobs...)
			result, data, err := floc.Run(flow)
			if !result.IsFailed() {
				t.Fatalf("%s:%d expects result to be Failed but has %s", t.Name(), n, result.String())
			} else if data != nil {
				t.Fatalf("%s:%d expects data to be nil but has %v", t.Name(), n, data)
			} else if err == nil {
				t.Fatalf("%s:%d expects error to be not nil", t.Name(), n)
			}
		}(t, n)
	}
}

func TestSequence_Error(t *testing.T) {
	const max = 100

	for n := 1; n <= max; n++ {
		go func(t *testing.T, n int) {
			jobs := make([]floc.Job, n)
			for i := 0; i < n; i++ {
				if i == n-1 {
					jobs[i] = throw(fmt.Errorf("err%d", n))
				} else {
					jobs[i] = noop()
				}
			}

			flow := Sequence(jobs...)
			result, data, err := floc.Run(flow)
			if !result.IsFailed() {
				t.Fatalf("%s:%d expects result to be Failed but has %s", t.Name(), n, result.String())
			} else if data != nil {
				t.Fatalf("%s:%d expects data to be nil but has %v", t.Name(), n, data)
			} else if err == nil {
				t.Fatalf("%s:%d expects error to be not nil", t.Name(), n)
			}
		}(t, n)
	}
}
