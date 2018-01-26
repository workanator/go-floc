package run

import (
	"time"

	"gopkg.in/devishot/go-floc.v2"
)

func noop() floc.Job {
	return func(floc.Context, floc.Control) error {
		return nil
	}
}

func complete(data interface{}) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		ctrl.Complete(data)
		return nil
	}
}

func cancel(data interface{}) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		ctrl.Cancel(data)
		return nil
	}
}

func fail(data interface{}, err error) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		ctrl.Fail(data, err)
		return nil
	}
}

func throw(err error) floc.Job {
	return func(ctx floc.Context, ctrl floc.Control) error {
		return err
	}
}

func waitUntilFinished() floc.Job {
	return Wait(func(ctx floc.Context) bool { return false }, time.Millisecond)
}

func yes() floc.Predicate {
	return func(ctx floc.Context) bool {
		return true
	}
}

func no() floc.Predicate {
	return func(ctx floc.Context) bool {
		return false
	}
}

func countdown(start int) floc.Predicate {
	return func(ctx floc.Context) bool {
		if start > 0 {
			start--
			return true
		}

		return false
	}
}
