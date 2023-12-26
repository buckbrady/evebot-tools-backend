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
	TypeCronJobEsiUniverseSystemKills = "cronjob:esi:universe:systemkills"
)

func NewCronJobUniverseSystemKillsTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobPayloadWithType{
		Timestamp: time.Now().UTC(),
		TTL:       3600,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseSystemKills, payload), nil
}

func HandleCronJobUniverseSystemKillsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobPayloadWithType
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	resp, _, err := esi.EVE.ESI.UniverseApi.GetUniverseSystemKills(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, data := range resp {
		record := model.UniverseSystemKill{
			SystemID:  data.SystemId,
			NpcKills:  data.NpcKills,
			PodKills:  data.PodKills,
			ShipKills: data.ShipKills,
			Timestamp: time.Now().UTC(),
		}
		err = database.Use(db).UniverseSystemKill.WithContext(ctx).Save(&record)
		if err != nil {
			log.Err(err).Any("systemID", data.SystemId).Msg("failed to create universe system kill record")
		}
	}

	return nil
}
