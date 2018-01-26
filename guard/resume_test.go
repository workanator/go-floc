package guard

import (
	"testing"

	"fmt"

	"gopkg.in/devishot/go-floc.v2"
	"gopkg.in/devishot/go-floc.v2/run"
)

func TestResume(t *testing.T) {
	flow := run.Sequence(
		Resume(floc.EmptyResultMask(), Cancel(nil)),
		Complete(nil),
	)

	result, _, _ := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	}
}

func TestResume_CanceledFiltered(t *testing.T) {
	flow := run.Sequence(
		Resume(floc.NewResultMask(floc.Canceled), Cancel(nil)),
		Complete(nil),
	)

	result, _, _ := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	}
}

func TestResume_CanceledNotFiltered(t *testing.T) {
	flow := run.Sequence(
		Resume(floc.NewResultMask(floc.None), Cancel(nil)),
		Complete(nil),
	)

	result, _, _ := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	}
}

func TestResume_CompletedFiltered(t *testing.T) {
	flow := run.Sequence(
		Resume(floc.NewResultMask(floc.Completed), Complete(nil)),
		Cancel(nil),
	)

	result, _, _ := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	}
}

func TestResume_CompletedNotFiltered(t *testing.T) {
	flow := run.Sequence(
		Resume(floc.NewResultMask(floc.None), Complete(nil)),
		Cancel(nil),
	)

	result, _, _ := floc.Run(flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be Completed but has %s", t.Name(), result.String())
	}
}

func TestResume_FailedFiltered(t *testing.T) {
	errTpl := fmt.Errorf("err %s", t.Name())

	flow := run.Sequence(
		Resume(floc.NewResultMask(floc.Failed), Fail(nil, errTpl)),
		Cancel(nil),
	)

	result, _, _ := floc.Run(flow)
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
	}
}

func TestResume_CompleteNotFiltered(t *testing.T) {
	dataTpl := "failed"
	errTpl := fmt.Errorf("err %s", t.Name())

	flow := run.Sequence(
		Resume(floc.NewResultMask(floc.None), Fail(dataTpl, errTpl)),
		Cancel(nil),
	)

	result, data, err := floc.Run(flow)
	if !result.IsFailed() {
		t.Fatalf("%s expects result to be Failed but has %s", t.Name(), result.String())
	} else if data == nil {
		t.Fatalf("%s expects data to be not nil", t.Name())
	} else if s, ok := data.(string); !ok {
		t.Fatalf("%s expects data to be of type string but has %T", t.Name(), data)
	} else if s != dataTpl {
		t.Fatalf("%s expects data to be %s but has %s", t.Name(), dataTpl, s)
	} else if err == nil {
		t.Fatalf("%s expects error to be not nil", t.Name())
	} else if err.Error() != errTpl.Error() {
		t.Fatalf("%s expects error to be %s but has %s", t.Name(), errTpl.Error(), err.Error())
	}
}
