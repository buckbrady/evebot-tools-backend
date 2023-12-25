package tasks

import (
	"context"
	"encoding/json"
	"github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/model"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypeCronJobEsiStatus = "cronjob:esi:status"
)

type CronJobStatusPayload struct {
	Timestamp time.Time
	TTL       int
}

func NewCronJobStatusTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobStatusPayload{
		Timestamp: time.Now().UTC(),
		TTL:       30,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiStatus, payload), nil
}

func HandleCronJobStatusTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobStatusPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	resp, _, err := esi.EVE.ESI.StatusApi.GetStatus(context.Background(), nil)
	if err != nil {
		return err
	}
	record := model.ServerStatus{
		Players:       resp.Players,
		ServerVersion: resp.ServerVersion,
		StartTime:     resp.StartTime,
		Vip:           resp.Vip,
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = database.Use(db).ServerStatus.WithContext(ctx).Create(&record)
	return err
}
