package run

import (
	"testing"

	"fmt"

	"gopkg.in/devishot/go-floc.v2"
)

func TestIfNot_OneAlreadyFinished(t *testing.T) {
	ctx := floc.NewContext()
	defer ctx.Release()

	ctrl := floc.NewControl(ctx)
	defer ctrl.Release()

	flow := IfNot(no(), Then(cancel(nil)))

	ctrl.Complete(nil)

	result, _, _ := floc.RunWith(ctx, ctrl, flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects flow to be Completed but has %s", t.Name(), result.String())
	}
}

func TestIfNot_TwoAlreadyFinished(t *testing.T) {
	ctx := floc.NewContext()
	defer ctx.Release()

	ctrl := floc.NewControl(ctx)
	defer ctrl.Release()

	flow := IfNot(yes(), Then(noop()), Else(cancel(nil)))

	ctrl.Complete(nil)

	result, _, _ := floc.RunWith(ctx, ctrl, flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects flow to be Completed but has %s", t.Name(), result.String())
	}
}

func TestIfNot_PanicNoJobs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%s must panic when no jobs are given", t.Name())
		}
	}()

	IfNot(yes())
}

func TestIfNot_PanicLotsJobs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%s must panic when more than two jobs are given", t.Name())
		}
	}()

	IfNot(yes(), noop(), noop(), noop())
}

func TestIfNot_OneTrueNone(t *testing.T) {
	flow := IfNot(no(), noop())
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_OneTrueCompleted(t *testing.T) {
	flow := IfNot(no(), complete(nil))
	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_OneTrueCanceled(t *testing.T) {
	flow := IfNot(no(), cancel(nil))
	result, data, err := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_OneTrueFailed(t *testing.T) {
	flow := IfNot(no(), fail(nil, fmt.Errorf("err")))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestIfNot_OneTrueError(t *testing.T) {
	flow := IfNot(no(), throw(fmt.Errorf("err")))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestIfNo_OneFalseCompleted(t *testing.T) {
	flow := IfNot(yes(), complete(nil))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_OneFalseCanceled(t *testing.T) {
	flow := IfNot(yes(), cancel(nil))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_OneFalseFailed(t *testing.T) {
	flow := IfNot(yes(), fail(nil, fmt.Errorf("err")))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_OneFalseError(t *testing.T) {
	flow := IfNot(yes(), throw(fmt.Errorf("err")))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_TwoTrueNone(t *testing.T) {
	flow := IfNot(no(), Then(noop()), Else(fail(nil, nil)))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_TwoTrueCompleted(t *testing.T) {
	flow := IfNot(no(), Then(complete(nil)), Else(fail(nil, nil)))
	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_TwoTrueCanceled(t *testing.T) {
	flow := IfNot(no(), Then(cancel(nil)), Else(fail(nil, nil)))
	result, data, err := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_TwoTrueFailed(t *testing.T) {
	flow := IfNot(no(), Then(fail(nil, fmt.Errorf("err"))), Else(complete(nil)))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestIfNot_TwoTrueError(t *testing.T) {
	flow := IfNot(no(), Then(throw(fmt.Errorf("err"))), Else(complete(nil)))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestIfNot_TwoFalseNone(t *testing.T) {
	flow := IfNot(yes(), Then(fail(nil, nil)), Else(noop()))
	result, data, err := floc.Run(flow)
	if !result.IsNone() {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_TwoFalseCompleted(t *testing.T) {
	flow := IfNot(yes(), Then(fail(nil, nil)), Else(complete(nil)))
	result, data, err := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_TwoFalseCanceled(t *testing.T) {
	flow := IfNot(yes(), Then(fail(nil, nil)), Else(cancel(nil)))
	result, data, err := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err != nil {
		t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
	}
}

func TestIfNot_TwoFalseFailed(t *testing.T) {
	flow := IfNot(yes(), Then(complete(nil)), Else(fail(nil, fmt.Errorf("err"))))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}

func TestIfNot_TwoFalseError(t *testing.T) {
	flow := IfNot(yes(), Then(complete(nil)), Else(throw(fmt.Errorf("err"))))
	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data != nil {
		t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	}
}
