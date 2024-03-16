package strace

import (
	"testing"
	"time"

	"github.com/shawnfeng/sutil2/slog"
)

func show(i int, ctx ContextTrace) {

	slog.Infof("i=%d tm:%s %d", i, ctx.Duration(), ctx.Duration())
}

func TestCtx(t *testing.T) {
	ctx := NewContextTrace()

	show(1, ctx)

	time.Sleep(time.Nanosecond)
	show(2, ctx)

	time.Sleep(time.Microsecond)
	show(3, ctx)

	time.Sleep(time.Millisecond)
	show(4, ctx)

	time.Sleep(time.Second)
	show(5, ctx)

}
