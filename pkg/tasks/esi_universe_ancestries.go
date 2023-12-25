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
	TypeCronJobEsiUniverseAncestries = "cronjob:esi:universe:ancestries"
)

func NewCronJobUniverseAncestriesTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseAncestries, payload), nil
}

func HandleCronJobUniverseAncestriesTask(ctx context.Context, t *asynq.Task) error {
	resp, _, err := esi.EVE.ESI.UniverseApi.GetUniverseAncestries(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, data := range resp {
		record := model.UniverseAncestry{
			ID:               data.Id,
			Name:             data.Name,
			Description:      data.Description,
			ShortDescription: &data.ShortDescription,
			BloodlineID:      data.BloodlineId,
			IconID:           &data.IconId,
		}
		err = database.Use(db).UniverseAncestry.WithContext(ctx).Save(&record)
		if err != nil {
			return err
		}
	}

	return nil
}
