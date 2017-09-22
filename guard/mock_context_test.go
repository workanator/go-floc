package guard

import (
	"context"
	"testing"
	"time"

	"gopkg.in/workanator/go-floc.v2"
)

func TestMockContext_Done(t *testing.T) {
	oCtx := floc.NewContext()
	oCancelCtx, oCancel := context.WithCancel(oCtx.Ctx())
	oCtx.UpdateCtx(oCancelCtx)

	mCtx := floc.NewContext()
	mCancelCtx, mCancel := context.WithCancel(mCtx.Ctx())
	mCtx.UpdateCtx(mCancelCtx)

	defer mCancel()

	mock := mockContext{Context: oCtx, mock: mCtx}

	go func() {
		time.Sleep(time.Millisecond)
		oCancel()
	}()

	timer := time.NewTimer(5 * time.Millisecond)
	select {
	case <-oCtx.Done():
		// Ok
		timer.Stop()
	case <-mock.Done():
		// Not Ok
		t.Fatalf("%s expects original context to be canceled", t.Name())
	case <-timer.C:
		// Not Ok
		t.Fatalf("%s expects original context to be canceled in time", t.Name())
	}

	timer = time.NewTimer(time.Millisecond)
	select {
	case <-mock.Done():
		// Not Ok
		t.Fatalf("%s expects mock context to be not canceled", t.Name())
	case <-timer.C:
		// Ok
	}
}

func TestMockContext_Done2(t *testing.T) {
	oCtx := floc.NewContext()
	oCancelCtx, oCancel := context.WithCancel(oCtx.Ctx())
	oCtx.UpdateCtx(oCancelCtx)

	defer oCancel()

	mCtx := floc.NewContext()
	mCancelCtx, mCancel := context.WithCancel(mCtx.Ctx())
	mCtx.UpdateCtx(mCancelCtx)

	mock := mockContext{Context: oCtx, mock: mCtx}

	go func() {
		time.Sleep(time.Millisecond)
		mCancel()
	}()

	timer := time.NewTimer(5 * time.Millisecond)
	select {
	case <-oCtx.Done():
		// Not Ok
		t.Fatalf("%s expects mock context to be canceled", t.Name())
	case <-mock.Done():
		// Ok
	case <-timer.C:
		// Not Ok
		t.Fatalf("%s expects mock context to be canceled in time", t.Name())
	}

	timer = time.NewTimer(time.Millisecond)
	select {
	case <-oCtx.Done():
		// Not Ok
		t.Fatalf("%s expects original context to be not canceled", t.Name())
	case <-timer.C:
		// Ok
	}
}

func TestMockContext_Done3(t *testing.T) {
	oCtx := floc.NewContext()
	oCancelCtx, oCancel := context.WithCancel(oCtx.Ctx())
	oCtx.UpdateCtx(oCancelCtx)

	mCtx := floc.NewContext()
	mCancelCtx, mCancel := context.WithCancel(mCtx.Ctx())
	mCtx.UpdateCtx(mCancelCtx)

	mock := mockContext{Context: oCtx, mock: mCtx}

	go func() {
		time.Sleep(time.Millisecond)
		oCancel()
	}()

	timer := time.NewTimer(5 * time.Millisecond)
	select {
	case <-oCtx.Done():
		timer.Stop()
	case <-timer.C:
		t.Fatalf("%s expects original context to be canceled", t.Name())
	}

	go func() {
		time.Sleep(time.Millisecond)
		mCancel()
	}()

	timer = time.NewTimer(5 * time.Millisecond)
	select {
	case <-mock.Done():
		timer.Stop()
	case <-timer.C:
		t.Fatalf("%s expects mock context to be canceled", t.Name())
	}
}
