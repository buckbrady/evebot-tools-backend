package tasks

import (
	"fmt"
	"github.com/hibiken/asynq"
)

var (
	ESI_STANDARD   = GetQueueOpts("esi", "standard", 5)
	ESI_BACKGROUND = GetQueueOpts("esi", "background", 1)
	ESI_HIGH       = GetQueueOpts("esi", "high", 10)
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
