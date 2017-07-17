package guard

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
	"github.com/workanator/go-floc/run"
	"github.com/workanator/go-floc/state"
)

func TestResume(t *testing.T) {
	f := flow.New()
	s := state.New(nil)
	job := run.Sequence(
		Resume(floc.NewResultSet(), Cancel(nil)),
		Complete(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed, result)
	}
}

func TestResumeCancelFiltered(t *testing.T) {
	f := flow.New()
	s := state.New(nil)
	job := run.Sequence(
		Resume(floc.NewResultSet(floc.Canceled), Cancel(nil)),
		Complete(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed, result)
	}
}

func TestResumeWithCancelNotFiltered(t *testing.T) {
	f := flow.New()
	s := state.New(nil)
	job := run.Sequence(
		Resume(floc.NewResultSet(floc.None), Cancel(nil)),
		Complete(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled, result)
	}
}

func TestResumeCompleteFiltered(t *testing.T) {
	f := flow.New()
	s := state.New(nil)
	job := run.Sequence(
		Resume(floc.NewResultSet(floc.Completed), Complete(nil)),
		Cancel(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled, result)
	}
}

func TestResumeWithCompleteNotFiltered(t *testing.T) {
	f := flow.New()
	s := state.New(nil)
	job := run.Sequence(
		Resume(floc.NewResultSet(floc.None), Complete(nil)),
		Cancel(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed, result)
	}
}
