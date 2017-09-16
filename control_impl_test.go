package floc

import (
	"fmt"
	"testing"
)

var templateData = []struct {
	data interface{}
	test func(sample interface{}) bool
}{
	{nil, func(sample interface{}) bool { return sample == nil }},
	{string("ABC"), func(sample interface{}) bool { s, ok := sample.(string); return ok && s == "ABC" }},
	{int(123), func(sample interface{}) bool { d, ok := sample.(int); return ok && d == 123 }},
	{float64(3.1415), func(sample interface{}) bool { f, ok := sample.(float64); return ok && f >= 3.1415 && f <= 3.14151 }},
}

func TestNewControl(t *testing.T) {
	ctx := NewContext()
	defer ctx.Release()

	ctrl := NewControl(ctx)
	defer ctrl.Release()

	if ctrl.IsFinished() == true {
		t.Fatalf("%s expects Control to be not finished", t.Name())
	}

	result, _, _ := ctrl.Result()
	if result.IsNone() == false {
		t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
	}
}

func TestNewControl_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%s must panic when ctx is nil", t.Name())
		}
	}()

	ctrl := NewControl(nil)
	defer ctrl.Release()
}

func TestFlowControl_Release(t *testing.T) {
	ctx := NewContext()
	defer ctx.Release()

	ctrl := NewControl(ctx)
	ctrl.Release()

	if ctrl.IsFinished() == false {
		result, _, _ := ctrl.Result()
		t.Fatalf("%s expects Control to be finished on Release but has state %s", t.Name(), result.String())
	}

	result, _, _ := ctrl.Result()
	if result.IsCanceled() == false {
		t.Fatalf("%s expects Control to be Canceled bu has %s", t.Name(), result.String())
	}
}

func TestFlowControl_Release2(t *testing.T) {
	ctx := NewContext()
	defer ctx.Release()

	ctrl := NewControl(ctx)
	ctrl.Complete(nil)
	ctrl.Release()

	if ctrl.IsFinished() == false {
		result, _, _ := ctrl.Result()
		t.Fatalf("%s expects Control to be finished on Release but has state %s", t.Name(), result.String())
	}

	result, _, _ := ctrl.Result()
	if result.IsCompleted() == false {
		t.Fatalf("%s expects Control to be Completed but has %s", t.Name(), result.String())
	}
}

func TestFlowControl_Cancel(t *testing.T) {
	for i, td := range templateData {
		ctx := NewContext()
		ctrl := NewControl(ctx)

		ctrl.Cancel(td.data)
		result, data, err := ctrl.Result()
		if err != nil {
			t.Fatalf("%s expects error to be nil bu has %s", t.Name(), err.Error())
		} else if result.IsCanceled() == false {
			t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
		} else if td.test(data) == false {
			t.Fatalf("%s failed to test data sample %d:%v", t.Name(), i, data)
		}

		ctrl.Release()
		ctx.Release()
	}
}

func TestFlowControl_Complete(t *testing.T) {
	for i, td := range templateData {
		ctx := NewContext()
		ctrl := NewControl(ctx)

		ctrl.Complete(td.data)
		result, data, err := ctrl.Result()
		if err != nil {
			t.Fatalf("%s expects error to be nil bu has %s", t.Name(), err.Error())
		} else if result.IsCompleted() == false {
			t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
		} else if td.test(data) == false {
			t.Fatalf("%s failed to test data sample %d:%v", t.Name(), i, data)
		}

		ctrl.Release()
		ctx.Release()
	}
}

func TestFlowControl_Fail(t *testing.T) {
	for i, td := range templateData {
		ctx := NewContext()
		ctrl := NewControl(ctx)

		ctrl.Fail(td.data, fmt.Errorf("fail %d", i))
		result, data, err := ctrl.Result()
		if err == nil {
			t.Fatalf("%s expects error to be not nil", t.Name())
		} else if result.IsFailed() == false {
			t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
		} else if td.test(data) == false {
			t.Fatalf("%s failed to test data sample %d:%v", t.Name(), i, data)
		}

		ctrl.Release()
		ctx.Release()
	}
}

func TestFlowControl_IsFinished(t *testing.T) {
	var finishers = []func(ctrl Control){
		func(ctrl Control) { ctrl.Complete(nil) },
		func(ctrl Control) { ctrl.Cancel(nil) },
		func(ctrl Control) { ctrl.Fail(nil, nil) },
	}

	for _, finish := range finishers {
		ctx := NewContext()
		ctrl := NewControl(ctx)

		if ctrl.IsFinished() == true {
			t.Fatalf("%s expects Control to be not finished", t.Name())
		}

		finish(ctrl)
		if ctrl.IsFinished() == false {
			t.Fatalf("%s expects Control to be finished", t.Name())
		}

		ctrl.Release()
		ctx.Release()
	}
}
