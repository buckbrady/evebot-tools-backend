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
	TypeCronJobEsiMarketPrices = "cronjob:esi:market:prices"
)

func NewCronJobMarketPricesTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobPayload{
		Timestamp: time.Now().UTC(),
		TTL:       3600,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiMarketPrices, payload), nil
}

func HandleCronJobMarketPricesTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	resp, _, err := esi.EVE.ESI.MarketApi.GetMarketsPrices(context.Background(), nil)
	if err != nil {
		return err
	}
	for _, price := range resp {
		record := model.MarketPrice{
			TypeID:        price.TypeId,
			AdjustedPrice: price.AdjustedPrice,
			AveragePrice:  price.AveragePrice,
			Timestamp:     time.Now().UTC(),
		}
		err = database.Use(db).MarketPrice.WithContext(ctx).Create(&record)
		if err != nil {
			log.Err(err).Any("typeID", price.TypeId).Msg("failed to create market price record")
		}

	}
	return nil
}
