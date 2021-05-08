package run

import (
	"testing"

	"fmt"

	"github.com/workanator/go-floc"
)

func TestIf_OneAlreadyFinished(t *testing.T) {
	ctx := floc.NewContext()
	defer ctx.Release()

	ctrl := floc.NewControl(ctx)
	defer ctrl.Release()

	flow := If(yes(), Then(cancel(nil)))

	ctrl.Complete(nil)

	result, _, _ := floc.RunWith(ctx, ctrl, flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects flow to be Completed but has %s", t.Name(), result.String())
	}
}

func TestIf_TwoAlreadyFinished(t *testing.T) {
	ctx := floc.NewContext()
	defer ctx.Release()

	ctrl := floc.NewControl(ctx)
	defer ctrl.Release()

	flow := If(no(), Then(noop()), Else(cancel(nil)))

	ctrl.Complete(nil)

	result, _, _ := floc.RunWith(ctx, ctrl, flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects flow to be Completed but has %s", t.Name(), result.String())
	}
}

func TestIf_PanicNoJobs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%s must panic when no jobs are given", t.Name())
		}
	}()

	If(yes())
}

func TestIf_PanicLotsJobs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%s must panic when more than two jobs are given", t.Name())
		}
	}()

	If(yes(), noop(), noop(), noop())
}

func TestIf_OneTrueNone(t *testing.T) {
	flow := If(yes(), noop())
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_OneTrueCompleted(t *testing.T) {
	flow := If(yes(), complete(nil))
	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_OneTrueCanceled(t *testing.T) {
	flow := If(yes(), cancel(nil))
	result, data, err := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_OneTrueFailed(t *testing.T) {
	flow := If(yes(), fail(nil, fmt.Errorf("err")))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestIf_OneTrueError(t *testing.T) {
	flow := If(yes(), throw(fmt.Errorf("err")))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestIf_OneFalseCompleted(t *testing.T) {
	flow := If(no(), complete(nil))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_OneFalseCanceled(t *testing.T) {
	flow := If(no(), cancel(nil))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_OneFalseFailed(t *testing.T) {
	flow := If(no(), fail(nil, fmt.Errorf("err")))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_OneFalseError(t *testing.T) {
	flow := If(no(), throw(fmt.Errorf("err")))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_TwoTrueNone(t *testing.T) {
	flow := If(yes(), Then(noop()), Else(fail(nil, nil)))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_TwoTrueCompleted(t *testing.T) {
	flow := If(yes(), Then(complete(nil)), Else(fail(nil, nil)))
	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_TwoTrueCanceled(t *testing.T) {
	flow := If(yes(), Then(cancel(nil)), Else(fail(nil, nil)))
	result, data, err := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_TwoTrueFailed(t *testing.T) {
	flow := If(yes(), Then(fail(nil, fmt.Errorf("err"))), Else(complete(nil)))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestIf_TwoTrueError(t *testing.T) {
	flow := If(yes(), Then(throw(fmt.Errorf("err"))), Else(complete(nil)))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestIf_TwoFalseNone(t *testing.T) {
	flow := If(no(), Then(fail(nil, nil)), Else(noop()))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_TwoFalseCompleted(t *testing.T) {
	flow := If(no(), Then(fail(nil, nil)), Else(complete(nil)))
	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_TwoFalseCanceled(t *testing.T) {
	flow := If(no(), Then(fail(nil, nil)), Else(cancel(nil)))
	result, data, err := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIf_TwoFalseFailed(t *testing.T) {
	flow := If(no(), Then(complete(nil)), Else(fail(nil, fmt.Errorf("err"))))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestIf_TwoFalseError(t *testing.T) {
	flow := If(no(), Then(complete(nil)), Else(throw(fmt.Errorf("err"))))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}
