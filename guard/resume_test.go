package guard

import (
	"testing"

	floc "gopkg.in/workanator/go-floc.v1"
	"gopkg.in/workanator/go-floc.v1/run"
)

func TestResume(t *testing.T) {
	f := floc.NewFlow()
	s := floc.NewState(nil)
	job := run.Sequence(
		Resume(floc.NewResultSet(), Cancel(nil)),
		Complete(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed.String(), result)
	}
}

func TestResumeCancelFiltered(t *testing.T) {
	f := floc.NewFlow()
	s := floc.NewState(nil)
	job := run.Sequence(
		Resume(floc.NewResultSet(floc.Canceled), Cancel(nil)),
		Complete(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed.String(), result)
	}
}

func TestResumeWithCancelNotFiltered(t *testing.T) {
	f := floc.NewFlow()
	s := floc.NewState(nil)
	job := run.Sequence(
		Resume(floc.NewResultSet(floc.None), Cancel(nil)),
		Complete(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled.String(), result)
	}
}

func TestResumeCompleteFiltered(t *testing.T) {
	f := floc.NewFlow()
	s := floc.NewState(nil)
	job := run.Sequence(
		Resume(floc.NewResultSet(floc.Completed), Complete(nil)),
		Cancel(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCanceled() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Canceled.String(), result)
	}
}

func TestResumeWithCompleteNotFiltered(t *testing.T) {
	f := floc.NewFlow()
	s := floc.NewState(nil)
	job := run.Sequence(
		Resume(floc.NewResultSet(floc.None), Complete(nil)),
		Cancel(nil),
	)

	floc.Run(f, s, nil, job)

	result, _ := f.Result()
	if !result.IsCompleted() {
		t.Fatalf("%s expects result to be %s but has %s", t.Name(), floc.Completed.String(), result)
	}
}
