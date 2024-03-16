package strace

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
)

var snowFlakeNode *snowflake.Node

func init() {
	var err error
	snowFlakeNode, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatal("fatal init snowflake ", err)
	}
}

func GenId() snowflake.ID {
	return snowFlakeNode.Generate()
}

type ContextTrace struct {
	// event 起始id号，方便channel追踪
	id snowflake.ID
	// begin time，方便追踪每一步骤的效率
	since time.Time
	// 通用上下文
	//ctx map[string]interface{}
}

func NewContextTrace() ContextTrace {

	return ContextTrace{
		id:    GenId(),
		since: time.Now(),
		//ctx:   make(map[string]interface{}),
	}
}

func (m ContextTrace) String() string {
	//return fmt.Sprintf("[CT:%d sc:%d du:%s]", m.id, m.since.UnixNano(), time.Since(m.since))
	//return fmt.Sprintf("[CT:%d du:%s]", m.id, time.Since(m.since))
	//strconv.FormatInt(int64(m.tagSnow), 36)
	//return fmt.Sprintf("[CT:%s]", m.id)
	return fmt.Sprintf("[CT:%s]", strconv.FormatInt(int64(m.id), 36))

}

func (m ContextTrace) Id() snowflake.ID {
	return m.id
}

func (m ContextTrace) Duration() time.Duration {
	return time.Since(m.since)
}

func (m ContextTrace) Since() time.Time {
	return m.since
}

func (m ContextTrace) UnixMilli() int64 {
	return m.since.UnixMilli()
}
