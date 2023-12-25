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
	TypeCronJobEsiUniversePlanets = "cronjob:esi:universe:planets"
)

func NewCronJobUniversePlanetsTask(typeID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    typeID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniversePlanets, payload), nil
}

func HandleCronJobUniversePlanetsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobUniverseTypesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	if p.TypeID < 1 {
		log.Error().Msg("invalid typeID")
		return nil
	}

	data, _, err := esi.EVE.ESI.UniverseApi.GetUniversePlanetsPlanetId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	if data.PlanetId < 1 {
		log.Error().Msg("invalid planetID")
		return nil
	}

	record := model.UniversePlanet{
		ID:        data.PlanetId,
		Name:      data.Name,
		SystemID:  data.SystemId,
		TypeID:    data.TypeId,
		PositionX: data.Position.X,
		PositionY: data.Position.Y,
		PositionZ: data.Position.Z,
	}
	err = database.Use(db).UniversePlanet.WithContext(ctx).Save(&record)
	if err != nil {
		return err
	}

	return nil
}
