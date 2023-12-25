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
	TypeCronJobEsiUniverseFactions = "cronjob:esi:universe:factions"
)

func NewCronJobUniverseFactionsTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseFactions, payload), nil
}

func HandleCronJobUniverseFactionsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobUniverseTypesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseFactions(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, d := range data {
		record := model.UniverseFaction{
			ID:                   d.FactionId,
			Name:                 d.Name,
			Description:          d.Description,
			MilitiaCorporationID: &d.MilitiaCorporationId,
			SizeFactor:           float64(d.SizeFactor),
			StationCount:         d.StationCount,
			StationSystemCount:   d.StationSystemCount,
			SolarSystemID:        &d.SolarSystemId,
			CorporationID:        &d.CorporationId,
		}
		err = database.Use(db).UniverseFaction.WithContext(ctx).Save(&record)
		if err != nil {
			log.Err(err).Any("factionID", d.FactionId).Msg("failed to save faction data")
		}
	}
	return nil
}
