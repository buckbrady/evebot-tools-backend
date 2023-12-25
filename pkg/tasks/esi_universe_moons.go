package tasks

import (
	"context"
	"encoding/json"
	"github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/model"
	"github.com/buckbrady/evebot-tools-backend/pkg/utils"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	TypeCronJobEsiUniverseMoons = "cronjob:esi:universe:moons"
)

func NewCronJobUniverseMoonsTask(typeID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    typeID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseMoons, payload), nil
}

func HandleCronJobUniverseMoonsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobUniverseTypesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	if p.TypeID < 1 {
		log.Error().Msg("invalid typeID")
		return nil
	}
	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseMoonsMoonId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	if data.MoonId < 1 {
		log.Error().Msg("invalid moonID")
		return nil
	}

	record := model.UniverseMoon{
		ID:        data.MoonId,
		Name:      data.Name,
		SystemID:  data.SystemId,
		PositionX: data.Position.X,
		PositionY: data.Position.Y,
		PositionZ: data.Position.Z,
	}
	dbctx, cancel := utils.NewDBCtx(ctx, 60)
	defer cancel()
	err = database.Use(db).UniverseMoon.WithContext(dbctx).Save(&record)
	if err != nil {
		return err
	}
	return nil
}
