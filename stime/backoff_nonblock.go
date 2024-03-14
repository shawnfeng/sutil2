package stime

import (
	"fmt"
	"sync/atomic"
	"time"
)

type BackOffNonblockCtrl struct {
	// 发生4xx 失败时候退控制
	failWaitUntil int64
	failIntv      int64

	failBackoffCeil int64
	failStep        int64
}

// 指定每次退避增加的时间最小步长，和上限，单位 毫秒
// 如果为0，默认failStep=65000 failBackoffCeil=60*1000*10 ，最大间隔到10分
func NewBackOffNonblockCtrl(failStep, failBackoffCeil int64) *BackOffNonblockCtrl {
	if failStep <= 0 {
		failStep = 65 * 1000
	}

	if failBackoffCeil <= 0 {
		failBackoffCeil = 60 * 1000 * 10
	}

	return &BackOffNonblockCtrl{
		failWaitUntil:   time.Now().UnixMilli(),
		failBackoffCeil: failBackoffCeil,
		failStep:        failStep,
	}
}

func (m *BackOffNonblockCtrl) String() string {
	return fmt.Sprintf("ceil:%d step:%d intv:%d until:%d", m.failBackoffCeil, m.failStep, m.failIntv, m.failWaitUntil)
}

// 判断是否在退避周期
func (m *BackOffNonblockCtrl) IsNeedWait() bool {
	unt := atomic.LoadInt64(&m.failWaitUntil)
	return unt > time.Now().UnixMilli()
}

// 延长退避，连续调用会让退避时间指数增加直到达到上限值
func (m *BackOffNonblockCtrl) IncreaseWait() {
	fc := atomic.LoadInt64(&m.failIntv) + m.failStep
	if fc > m.failBackoffCeil {
		fc = m.failBackoffCeil
	}

	atomic.StoreInt64(&m.failIntv, fc)
	atomic.StoreInt64(&m.failWaitUntil, time.Now().UnixMilli()+fc)
}

// 结束退避
func (m *BackOffNonblockCtrl) StopWait() {
	atomic.StoreInt64(&m.failWaitUntil, time.Now().UnixMilli())
	atomic.StoreInt64(&m.failIntv, 0)
}
