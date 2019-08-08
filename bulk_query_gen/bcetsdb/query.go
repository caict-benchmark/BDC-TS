package bcetsdb

import (
	"fmt"
	"sync"
)

type KairosDbQuery struct {
    Body               string
}

var KairosDbQueryPool sync.Pool = sync.Pool{
    New: func() interface{} {
        return &KairosDbQuery{
            "",
        }
    },
}

func NewKairosDbQuery() *KairosDbQuery {
    return KairosDbQueryPool.Get().(*KairosDbQuery)
}

func (q *KairosDbQuery) String() string {
    return fmt.Sprintf("%s", q.Body)
}

func (q *KairosDbQuery) HumanLabelName() []byte {
    return nil
}
func (q *KairosDbQuery) HumanDescriptionName() []byte {
    return nil
}

func (q *KairosDbQuery) Release() {
    q.Body = q.Body[:0]
    KairosDbQueryPool.Put(q)
}