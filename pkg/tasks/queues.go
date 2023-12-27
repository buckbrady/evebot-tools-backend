package tasks

import (
	"github.com/hibiken/asynq"
)

var (
	BACKGROUND_QUEUE = GetQueueOpts("background", 1)
	STANDARD_QUEUE   = GetQueueOpts("standard", 2)
	CRITICAL_QUEUE   = GetQueueOpts("high", 3)
	REALTIME_QUEUE   = GetQueueOpts("realtime", 4)
)

type QueueOpts struct {
	Name     string
	Priority int // 1 = low, 2 = normal, 3 = high
}

func GetQueueOpts(name string, priority int) QueueOpts {
	return QueueOpts{
		Name:     name,
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
