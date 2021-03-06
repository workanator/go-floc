package run

import (
	"testing"
	"time"

	"github.com/workanator/go-floc/v3"
)

func TestWait_AlreadyFinished(t *testing.T) {
	ctx := floc.NewContext()
	defer ctx.Release()

	ctrl := floc.NewControl(ctx)
	defer ctrl.Release()

	flow := Sequence(Wait(yes(), time.Millisecond), cancel(nil))

	ctrl.Complete(nil)

	result, _, _ := floc.RunWith(ctx, ctrl, flow)
	if !result.IsCompleted() {
		t.Fatalf("%s expects flow to be Completed but has %s", t.Name(), result.String())
	}
}
