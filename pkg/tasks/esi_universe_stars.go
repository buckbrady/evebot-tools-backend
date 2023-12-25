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
	TypeCronJobEsiUniverseStars = "cronjob:esi:universe:stars"
)

func NewCronJobUniverseStarsTask(typeID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    typeID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseStars, payload), nil
}

func HandleCronJobUniverseStarsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobUniverseTypesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	if p.TypeID < 1 {
		log.Error().Msg("invalid typeID")
		return nil
	}

	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseStarsStarId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	if data.TypeId < 1 {
		log.Error().Msg("invalid typeID")
		return nil
	}

	record := model.UniverseStar{
		ID:            p.TypeID,
		Name:          data.Name,
		SolarSystemID: data.SolarSystemId,
		TypeID:        data.TypeId,
		Age:           data.Age,
		Luminosity:    float64(data.Luminosity),
		Radius:        data.Radius,
		SpectralClass: data.SpectralClass,
		Temperature:   data.Temperature,
	}
	dbctx, cancel := utils.NewDBCtx(ctx, 60)
	defer cancel()
	err = database.Use(db).UniverseStar.WithContext(dbctx).Save(&record)
	if err != nil {
		return err
	}

	return nil
}
