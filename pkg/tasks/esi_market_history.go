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
	TypeCronJobEsiMarketHistory = "cronjob:esi:market:history"
)

func NewCronJobMarketHistoryTask(typeID int32, regionID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobPayloadTypeWithRegion{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    typeID,
		RegionID:  regionID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiMarketHistory, payload), nil
}

func HandleCronJobMarketHistoryTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobPayloadTypeWithRegion
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	if p.TypeID < 1 || p.RegionID < 1 {
		log.Error().Any("typeID", p.TypeID).Any("regionID", p.RegionID).Msg("invalid typeID or regionID")
		return nil
	}
	data, _, err := esi.EVE.ESI.MarketApi.GetMarketsRegionIdHistory(context.Background(), p.RegionID, p.TypeID, nil)
	if err != nil {
		return err
	}
	for _, history := range data {
		date, err := time.Parse("2023-01-02", history.Date)
		if err != nil {
			log.Error().Err(err).Any("date", history.Date).Msg("failed to parse date")
			continue
		}
		record := model.MarketHistory{
			TypeID:     p.TypeID,
			RegionID:   p.RegionID,
			Date:       date,
			Average:    history.Average,
			Highest:    history.Highest,
			Lowest:     history.Lowest,
			OrderCount: history.OrderCount,
			Volume:     history.Volume,
		}
		err = database.Use(db).MarketHistory.WithContext(ctx).Save(&record)
		if err != nil {
			log.Err(err).Any("typeID", p.TypeID).Any("regionID", p.RegionID).Msg("failed to create market history record")
		}
	}

	return nil
}
