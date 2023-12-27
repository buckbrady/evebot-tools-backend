package tasks

import (
	"context"
	"encoding/json"
	esi2 "github.com/antihax/goesi/esi"
	"github.com/antihax/goesi/optional"
	"github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/model"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

const (
	TypeCronJobEsiMarketOrders = "cronjob:esi:market:orders"
)

func NewCronJobMarketOrdersTask(typeID int32, regionID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobPayloadTypeWithRegion{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    typeID,
		RegionID:  regionID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiMarketOrders, payload), nil
}

func HandleCronJobMarketOrdersTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobPayloadTypeWithRegion
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		log.Error().
			Err(err).
			Any("regionID", p.RegionID).
			Any("typeID", p.TypeID).
			Any("payload", string(t.Payload())).
			Any("task", t.Type()).
			Msg("failed to unmarshal payload")
		return nil
	}

	if p.TypeID < 1 || p.RegionID < 1 {
		log.Error().
			Any("typeID", p.TypeID).
			Any("regionID", p.RegionID).
			Any("task", t.Type()).
			Msg("invalid typeID or regionID")
		return nil
	}

	_, req, err := esi.EVE.ESI.MarketApi.GetMarketsRegionIdOrders(context.Background(), "all", p.RegionID, &esi2.GetMarketsRegionIdOrdersOpts{TypeId: optional.NewInt32(p.TypeID)})
	if err != nil {
		log.Error().
			Err(err).
			Any("typeID", p.TypeID).
			Any("regionID", p.RegionID).
			Any("task", t.Type()).
			Msg("failed to get market orders")
		return err
	}
	pages, err := strconv.Atoi(req.Header.Get("X-Pages"))

	for i := 1; i <= pages; i++ {
		data, _, err := esi.EVE.ESI.MarketApi.GetMarketsRegionIdOrders(context.Background(), "all", p.RegionID, &esi2.GetMarketsRegionIdOrdersOpts{TypeId: optional.NewInt32(p.TypeID), Page: optional.NewInt32(int32(i))})
		if err != nil {
			log.Error().
				Err(err).
				Any("typeID", p.TypeID).
				Any("regionID", p.RegionID).
				Any("task", t.Type()).
				Msg("failed to get market orders")
			return err
		}
		for _, order := range data {
			record := model.MarketOrder{
				TypeID:       p.TypeID,
				RegionID:     p.RegionID,
				ID:           order.OrderId,
				IsBuyOrder:   order.IsBuyOrder,
				Price:        order.Price,
				VolumeRemain: order.VolumeRemain,
				VolumeTotal:  order.VolumeTotal,
				Issued:       order.Issued,
				Range:        order.Range_,
				MinVolume:    order.MinVolume,
				Duration:     order.Duration,
				LocationID:   order.LocationId,
				LastUpdated:  time.Now().UTC(),
			}
			err = database.Use(db).MarketOrder.WithContext(ctx).Save(&record)
			if err != nil {
				log.Error().
					Err(err).
					Any("typeID", p.TypeID).
					Any("regionID", p.RegionID).
					Any("task", t.Type()).
					Msg("failed to create market order record")
			}
		}
	}

	return nil
}
