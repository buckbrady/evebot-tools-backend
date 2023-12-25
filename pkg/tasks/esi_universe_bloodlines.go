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
	TypeCronJobEsiUniverseBloodlines = "cronjob:esi:universe:bloodlines"
)

func NewCronJobUniverseBloodlinesTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseBloodlines, payload), nil
}

func HandleCronJobUniverseBloodlinesTask(ctx context.Context, t *asynq.Task) error {
	resp, _, err := esi.EVE.ESI.UniverseApi.GetUniverseBloodlines(context.Background(), nil)
	if err != nil {
		return err
	}
	for _, data := range resp {
		record := model.UniverseBloodline{
			ID:            data.BloodlineId,
			Charisma:      data.Charisma,
			CorporationID: data.CorporationId,
			Description:   data.Description,
			Intelligence:  data.Intelligence,
			Memory:        data.Memory,
			Name:          data.Name,
			Perception:    data.Perception,
			RaceID:        data.RaceId,
			ShipTypeID:    data.ShipTypeId,
			Willpower:     data.Willpower,
		}
		err = database.Use(db).UniverseBloodline.WithContext(ctx).Save(&record)
		if err != nil {
			return err
		}
	}
	return nil
}
