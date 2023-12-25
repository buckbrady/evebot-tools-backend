package tasks

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypeCronJobEsiUniverseStations = "cronjob:esi:universe:stations"
)

func NewCronJobUniverseStationsTask(typeID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    typeID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseStations, payload), nil
}

func HandleCronJobUniverseStationsTask(ctx context.Context, t *asynq.Task) error {
	return nil
}
