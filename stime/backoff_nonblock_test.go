package stime

import (
	"testing"
	"time"

	"github.com/shawnfeng/web3/slog"
)

func TestNonBlock(t *testing.T) {
	fun := "TestNonBlock -->"
	bc := NewBackOffNonblockCtrl(1*1000, 50*1000)
	slog.Infof("%s u:%d bc:%s", fun, time.Now().UnixMilli(), bc)

	for i := 0; i < 60; i++ {
		need := bc.IsNeedWait()
		slog.Infoln(fun, "F", i, time.Now().UnixMilli(), need, bc)
		if i < 10 {
			bc.IncreaseWait()
		}
		slog.Infoln(fun, "E", i, time.Now().UnixMilli(), need, bc)

		if i == 8 {
			bc.StopWait()
			slog.Infoln(fun, "STOP", i, time.Now().UnixMilli(), need, bc)

		}

		time.Sleep(time.Second * 1)
	}
}
