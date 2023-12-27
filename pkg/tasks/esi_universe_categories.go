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
	TypeCronJobEsiUniverseCategories = "cronjob:esi:universe:categories"
)

func NewCronJobUniverseCategoriesTask(categoryID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobPayloadWithType{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    categoryID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseCategories, payload), nil
}

func HandleCronJobUniverseCategoriesTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobPayloadWithType
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	if p.TypeID < 1 {
		return nil
	}

	resp, _, err := esi.EVE.ESI.UniverseApi.GetUniverseCategoriesCategoryId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	record := model.UniverseCategory{
		ID:        resp.CategoryId,
		Name:      resp.Name,
		Published: resp.Published,
	}
	err = database.Use(db).UniverseCategory.WithContext(ctx).Save(&record)
	if err != nil {
		return err
	}

	for _, group := range resp.Groups {
		var taskGroup *asynq.Task
		var taskErr error
		taskGroup, taskErr = NewCronJobUniverseGroupsTask(group)
		if taskErr != nil {
			return taskErr
		}
		entityID, err := queueClient.Enqueue(taskGroup, BACKGROUND_QUEUE.GetQueue())
		if err != nil {
			log.Err(err).Any("groupID", group).Msg("failed to enqueue group task")
		}
		log.Info().Any("groupID", group).Any("entityID", entityID).Msg("enqueued group task")
	}

	return nil
}
