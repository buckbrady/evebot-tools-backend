package tasks

import (
	"fmt"
	"github.com/hibiken/asynq"
)

const (
	QUEUE_PRIORITY_LOW    = 1
	QUEUE_PRIORITY_NORMAL = 4
	QUEUE_PRIORITY_HIGH   = 5
)

var (
	ESI_STATUS_QUEUE            = GetQueueOpts("esi", "status", QUEUE_PRIORITY_NORMAL)
	ESI_UNIVERSE_QUEUE          = GetQueueOpts("esi", "universe", QUEUE_PRIORITY_LOW)
	ESI_MARKET_QUEUE            = GetQueueOpts("esi", "market", QUEUE_PRIORITY_HIGH)
	ESI_UNIVERSE_REALTIME_QUEUE = GetQueueOpts("esi", "universe_realtime", QUEUE_PRIORITY_HIGH)
)

type QueueOpts struct {
	Name     string
	Priority int // 1 = low, 2 = normal, 3 = high
}

func GetQueueOpts(group, name string, priority int) QueueOpts {
	q := fmt.Sprintf("%s_%s", group, name)
	return QueueOpts{
		Name:     q,
		Priority: priority,
	}
}

func (q QueueOpts) GetName() string {
	return q.Name
}

func (q QueueOpts) GetPriority() int {
	return q.Priority
}

func (q QueueOpts) GetQueue() asynq.Option {
	return asynq.Queue(q.Name)
}
