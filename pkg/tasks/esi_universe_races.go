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
	TypeCronJobEsiUniverseRaces = "cronjob:esi:universe:races"
)

func NewCronJubUniverseRacesTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobPayloadWithType{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseRaces, payload), nil
}

func HandleCronJobUniverseRacesTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobPayloadWithType
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	rasp, _, err := esi.EVE.ESI.UniverseApi.GetUniverseRaces(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, data := range rasp {
		record := model.UniverseRace{
			ID:          data.RaceId,
			Name:        data.Name,
			Description: data.Description,
			AllianceID:  data.AllianceId,
		}
		err = database.Use(db).UniverseRace.WithContext(ctx).Save(&record)
		if err != nil {
			log.Err(err).Any("raceID", data.RaceId).Msg("failed to save race data")
		}
	}

	return nil
}
