package tasks

import (
	"context"
	"encoding/json"
	"github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/buckbrady/evebot-tools-backend/pkg/model"
	//"github.com/buckbrady/evebot-tools-backend/pkg/database/models"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	TypeCronJobEsiUniverseConstellations = "cronjob:esi:universe:constellations"
)

// Tasks

func NewCronJobUnvierseConstellationsTask(id int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobPayloadWithType{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    id,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseConstellations, payload), nil
}

// Handlers

func HandleCronJobUniverseConstellationsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobPayloadWithType
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseConstellationsConstellationId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	record := model.UniverseConstellation{
		ID:        data.ConstellationId,
		Name:      data.Name,
		RegionID:  data.RegionId,
		PositionX: data.Position.X,
		PositionY: data.Position.Y,
		PositionZ: data.Position.Z,
	}
	err = database.Use(db).UniverseConstellation.WithContext(ctx).Save(&record)
	if err != nil {
		log.Err(err).Msg("failed to create universe constellation record")
		return err
	}

	for _, s := range data.Systems {
		task, err := NewCronJobUniverseSystemsTask(s)
		if err != nil {
			log.Err(err).Msgf("Failed to create universe system task: %d", s)
		}
		entryID, err := queueClient.Enqueue(task, ESI_BACKGROUND.GetQueue())
		if err != nil {
			log.Err(err).Msgf("Failed to register universe system task: %d", s)
		}
		log.Info().Any("entryID", entryID).Msgf("Registered universe system task: %d", s)
	}

	return nil
}
