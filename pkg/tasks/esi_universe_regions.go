package tasks

import (
	"context"
	"encoding/json"
	"github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/model"
	"github.com/rs/zerolog/log"

	"github.com/hibiken/asynq"
	"time"
)

const (
	TypeCronJobEsiUniverseRegions = "cronjob:esi:universe:regions"
)

// Tasks

func NewCronJobUniverseRegionsTask(typeID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    typeID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseRegions, payload), nil
}

// Handlers

func HandleCronJobUniverseRegionsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobUniverseTypesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	if p.TypeID < 1 {
		log.Error().Msg("invalid typeID")
		return nil
	}
	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseRegionsRegionId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	if data.RegionId < 1 {
		log.Error().Msg("invalid regionID")
		return nil
	}

	record := model.UniverseRegion{
		ID:          data.RegionId,
		Name:        data.Name,
		Description: data.Description,
	}

	err = database.Use(db).UniverseRegion.WithContext(ctx).Save(&record)
	if err != nil {
		return err
	}

	for _, cID := range data.Constellations {
		task, err := NewCronJobUnvierseConstellationsTask(cID)
		if err != nil {
			log.Err(err).Msgf("Failed to create universe constellation task: %d", cID)
			return err
		}
		entryID, err := queueClient.Enqueue(task, ESI_UNIVERSE_QUEUE.GetQueue())
		if err != nil {
			log.Err(err).Msgf("Failed to register universe constellation task: %d", cID)
			return err
		}
		log.Info().Any("entryID", entryID).Msgf("Registered universe constellation task: %d", cID)
	}
	return nil
}
