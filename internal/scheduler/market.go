package scheduler

import (
	"context"
	"github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/tasks"
	"github.com/rs/zerolog/log"
)

func scheduleMarketGroupsJob() {
	data, _, err := esi.EVE.ESI.MarketApi.GetMarketsGroups(context.Background(), nil)
	if err != nil {
		log.Err(err).Msg("Failed to get market groups")
		return
	}
	for _, g := range data {
		task, err := tasks.NewCronJobMarketGroupsTask(g)
		if err != nil {
			log.Err(err).Msgf("Failed to create market group task: %d", g)
		}
		entryID, err := queueClient.Enqueue(task, tasks.BACKGROUND_QUEUE.GetQueue())
		if err != nil {
			log.Err(err).Msgf("Failed to register market group task: %d", g)
		}
		log.Info().Any("entryID", entryID).Msgf("Registered market group task: %d", g)
	}
}

func scheduleMarketPricesJob() {
	task, err := tasks.NewCronJobMarketPricesTask()
	if err != nil {
		log.Err(err).Msg("Failed to create market prices task")
	}
	entryID, err := scheduler.Register("0 * * * *", task, tasks.BACKGROUND_QUEUE.GetQueue())
	if err != nil {
		log.Err(err).Msg("Failed to register market prices task")
	}
	log.Info().Any("entryID", entryID).Msg("Registered market prices task")
}

func scheduleMarketHistoryTask() {
	regions, err := database.Use(db).UniverseRegion.WithContext(context.Background()).Find()
	if err != nil {
		log.Err(err).Msg("Failed to get universe regions")
		return
	}
	for _, r := range regions {
		types, _, err := esi.EVE.ESI.MarketApi.GetMarketsRegionIdTypes(context.Background(), r.ID, nil)
		if err != nil {
			log.Err(err).Msgf("Failed to get types for region: %d", r.ID)
			continue
		}
		for _, t := range types {
			task, err := tasks.NewCronJobMarketHistoryTask(t, r.ID)
			if err != nil {
				log.Err(err).Any("regionID", r.ID).Any("typeID", t).Msgf("Failed to create market history task: %d", t)
			}
			entryID, err := queueClient.Enqueue(task, tasks.STANDARD_QUEUE.GetQueue())
			if err != nil {
				log.Err(err).Any("regionID", r.ID).Any("typeID", t).Msgf("Failed to register market history task: %d", t)
			}
			log.Info().Any("regionID", r.ID).Any("typeID", t).Any("entryID", entryID).Msgf("Registered market history task: %d", t)
		}
	}
}

func scheduleMarketOrdersTask() {
	regions, err := database.Use(db).UniverseRegion.WithContext(context.Background()).Find()
	if err != nil {
		log.Err(err).Msg("Failed to get universe regions")
		return
	}
	for _, r := range regions {
		types, _, err := esi.EVE.ESI.MarketApi.GetMarketsRegionIdTypes(context.Background(), r.ID, nil)
		if err != nil {
			log.Err(err).Msgf("Failed to get types for region: %d", r.ID)
			continue
		}
		for _, t := range types {
			task, err := tasks.NewCronJobMarketOrdersTask(t, r.ID)
			if err != nil {
				log.Err(err).Any("regionID", r.ID).Any("typeID", t).Msgf("Failed to create market orders task: %d", t)
			}
			entryID, err := queueClient.Enqueue(task, tasks.CRITICAL_QUEUE.GetQueue())
			if err != nil {
				log.Err(err).Any("regionID", r.ID).Any("typeID", t).Msgf("Failed to register market orders task: %d", t)
			}
			log.Info().Any("regionID", r.ID).Any("typeID", t).Any("entryID", entryID).Msgf("Registered market orders task: %d", t)
		}
	}

}
