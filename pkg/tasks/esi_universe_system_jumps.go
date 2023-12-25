package tasks

import (
	"context"
	"encoding/json"
	"github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/model"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	TypeCronJobEsiUniverseSystemJumps = "cronjob:esi:universe:systemjumps"
)

func NewCronJobUniverseSystemJumpsTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       3600,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseSystemJumps, payload), nil
}

func HandleCronJobUniverseSystemJumpsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobUniverseTypesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	resp, _, err := esi.EVE.ESI.UniverseApi.GetUniverseSystemJumps(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, data := range resp {
		record := model.UniverseSystemJump{
			SystemID:  data.SystemId,
			ShipJumps: data.ShipJumps,
			Timestamp: time.Now().UTC(),
		}
		err = database.Use(db).UniverseSystemJump.WithContext(ctx).Save(&record)
		if err != nil {
			log.Err(err).Any("systemID", data.SystemId).Msg("failed to create universe system jump record")
		}
	}

	return nil
}
