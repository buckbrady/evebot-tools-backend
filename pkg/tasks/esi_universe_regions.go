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
	payload, err := json.Marshal(CronJobPayloadWithType{
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
	var p CronJobPayloadWithType
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseRegionsRegionId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
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
		var taskCon *asynq.Task
		var taskErr error
		taskCon, taskErr = NewCronJobUnvierseConstellationsTask(cID)
		if taskErr != nil {
			log.Err(taskErr).Msgf("Failed to create universe constellation task: %d", cID)
			return taskErr
		}
		entryIDCon, taskErr := queueClient.Enqueue(taskCon, ESI_BACKGROUND.GetQueue())
		if taskErr != nil {
			log.Err(taskErr).Msgf("Failed to register universe constellation task: %d", cID)
			return taskErr
		}
		log.Info().Any("entryID", entryIDCon).Msgf("Registered universe constellation task: %d", cID)
	}
	return nil
}
