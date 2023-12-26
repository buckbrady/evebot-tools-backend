package tasks

import "github.com/hibiken/asynq"

const (
	TypeCronJobEsiMarketOrders = "cronjob:esi:market:orders"
)

func NewCronJobMarketOrdersTask(typeID int32, regionID int32) (*asynq.Task, error) {
	return nil, nil
}
