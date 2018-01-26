package run

import (
  "time"
  "testing"

  "gopkg.in/devishot/go-floc.v2"
)

func TestFirst_AlreadyFinished(t *testing.T) {
  ctx := floc.NewContext()
  defer ctx.Release()

  ctrl := floc.NewControl(ctx)
  defer ctrl.Release()

  flow := First(cancel(nil))

  ctrl.Complete(nil)

  result, _, _ := floc.RunWith(ctx, ctrl, flow)
  if !result.IsCompleted() {
    t.Fatalf("%s expects flow to be Completed but has %s", t.Name(), result.String())
  }
}

func TestFirst_Simple(t *testing.T) {
  flow := First(cancel(nil))
  result, data, err := floc.Run(flow)
  if !result.IsCanceled() {
    t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
  } else if data != nil {
    t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
  } else if err != nil {
    t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
  }
}

func TestFirst_None(t *testing.T) {
  flow := First(Delay(2*time.Millisecond, complete(nil)), Delay(3*time.Millisecond, cancel(nil)))
  result, data, err := floc.Run(flow)
  if !result.IsNone() {
    t.Fatalf("%s expects result to be None but has %s", t.Name(), result.String())
  } else if data != nil {
    t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
  } else if err != nil {
    t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
  }
}

func TestFirst_Canceled(t *testing.T) {
  flow := First(Delay(4*time.Millisecond, complete(nil)), Delay(3*time.Millisecond, cancel(nil)))
  result, data, err := floc.Run(flow)
  if !result.IsCanceled() {
    t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
  } else if data != nil {
    t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
  } else if err != nil {
    t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
  }
}

func TestFirst_Sequence(t *testing.T) {
  flow := Sequence(
    First(Delay(2*time.Millisecond, complete(nil)), Delay(1*time.Millisecond, cancel(nil))),
    complete(nil),
  )
  result, data, err := floc.Run(flow)
  if !result.IsCanceled() {
    t.Fatalf("%s expects result to be Canceled but has %s", t.Name(), result.String())
  } else if data != nil {
    t.Fatalf("%s expects data to be nil but has %v", t.Name(), data)
  } else if err != nil {
    t.Fatalf("%s expects error to be nil but has %v", t.Name(), err)
  }
}