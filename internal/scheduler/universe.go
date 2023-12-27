package scheduler

import (
	"context"
	esi2 "github.com/antihax/goesi/esi"
	"github.com/antihax/goesi/optional"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/tasks"
	"github.com/rs/zerolog/log"
	"strconv"
)

func scheduleUniverseTypesJob() {
	_, resp, err := esi.EVE.ESI.UniverseApi.GetUniverseTypes(context.Background(), nil)
	if err != nil {
		log.Err(err).Msg("Failed to get universe types")
		return
	}
	pages, _ := strconv.Atoi(resp.Header.Get("x-pages"))
	var typeInts []int32
	for i := 1; i <= pages; i++ {
		data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseTypes(context.Background(), &esi2.GetUniverseTypesOpts{Page: optional.NewInt32(int32(i))})
		if err != nil {
			log.Err(err).Msg("Failed to get universe types")
			continue
		}
		typeInts = append(typeInts, data...)
	}
	for _, t := range typeInts {
		task, err := tasks.NewCronJobUniverseTypesTask(t)
		if err != nil {
			log.Err(err).Msgf("Failed to create universe type task: %d", t)
		}
		entryID, err := queueClient.Enqueue(task, tasks.BACKGROUND_QUEUE.GetQueue())
		if err != nil {
			log.Err(err).Msgf("Failed to register universe type task: %d", t)
		}
		log.Info().Any("entryID", entryID).Msgf("Registered universe type task: %d", t)
	}
}

func scheduleUniverseRegionsJob() {
	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseRegions(context.Background(), nil)
	if err != nil {
		log.Err(err).Msg("Failed to get universe regions")
		return
	}
	for _, r := range data {
		task, err := tasks.NewCronJobUniverseRegionsTask(r)
		if err != nil {
			log.Err(err).Msgf("Failed to create universe region task: %d", r)
		}
		entryID, err := queueClient.Enqueue(task, tasks.BACKGROUND_QUEUE.GetQueue())
		if err != nil {
			log.Err(err).Msgf("Failed to register universe region task: %d", r)
		}
		log.Info().Any("entryID", entryID).Msgf("Registered universe region task: %d", r)
	}
}

func scheduleUniverseAncestriesJob() {
	task, err := tasks.NewCronJobUniverseAncestriesTask()
	if err != nil {
		log.Err(err).Msg("Failed to create universe ancestries task")
		return
	}
	entryID, err := scheduler.Register("10 11 * * *", task, tasks.BACKGROUND_QUEUE.GetQueue())
	if err != nil {
		log.Err(err).Msg("Failed to register universe ancestries task")
		return
	}
	log.Info().Any("entryID", entryID).Msg("Registered universe ancestries task")
}

func scheduleUniverseBloodlinesJob() {
	task, err := tasks.NewCronJobUniverseBloodlinesTask()
	if err != nil {
		log.Err(err).Msg("Failed to create universe bloodlines task")
		return
	}
	entryID, err := scheduler.Register("10 11 * * *", task, tasks.BACKGROUND_QUEUE.GetQueue())
	if err != nil {
		log.Err(err).Msg("Failed to register universe bloodlines task")
		return
	}
	log.Info().Any("entryID", entryID).Msg("Registered universe bloodlines task")
}

func scheduleUniverseSystemJumpsJob() {
	task, err := tasks.NewCronJobUniverseSystemJumpsTask()
	if err != nil {
		log.Err(err).Msg("failed to create system jumps task")
	}
	statusID, err := scheduler.Register("0 * * * *", task, tasks.CRITICAL_QUEUE.GetQueue())
	if err != nil {
		log.Err(err).Msg("failed to register system jumps task")
	} else {
		log.Info().Str("entryID", statusID).Msg("registered system jumps task")
	}
}

func scheduleUniverseSystemKillsJob() {
	task, err := tasks.NewCronJobUniverseSystemKillsTask()
	if err != nil {
		log.Err(err).Msg("failed to create system kills task")
	}
	statusID, err := scheduler.Register("0 * * * *", task, tasks.CRITICAL_QUEUE.GetQueue())
	if err != nil {
		log.Err(err).Msg("failed to register system kills task")
	} else {
		log.Info().Str("entryID", statusID).Msg("registered system kills task")
	}
}

func scheduleUniverseGraphicsJob() {
	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseGraphics(context.Background(), nil)
	if err != nil {
		log.Err(err).Msg("Failed to get universe graphics")
		return
	}
	for _, r := range data {
		task, err := tasks.NewCronJobUniverseGraphicsTask(r)
		if err != nil {
			log.Err(err).Msgf("Failed to create universe graphics task: %d", r)
		}
		entryID, err := queueClient.Enqueue(task, tasks.BACKGROUND_QUEUE.GetQueue())
		if err != nil {
			log.Err(err).Msgf("Failed to register universe graphics task: %d", r)
		}
		log.Info().Any("entryID", entryID).Msgf("Registered universe graphics task: %d", r)
	}
}

func scheduleUniverseFactionsJob() {
	task, err := tasks.NewCronJobUniverseFactionsTask()
	if err != nil {
		log.Err(err).Msg("failed to create system kills task")
	}
	statusID, err := scheduler.Register("10 11 * * *", task, tasks.BACKGROUND_QUEUE.GetQueue())
	if err != nil {
		log.Err(err).Msg("failed to register factions task")
	} else {
		log.Info().Str("entryID", statusID).Msg("registered factions task")
	}
}

func scheduleUniverseRacesJob() {
	task, err := tasks.NewCronJubUniverseRacesTask()
	if err != nil {
		log.Err(err).Msg("Failed to create universe races task")
		return
	}
	entryID, err := scheduler.Register("10 11 * * *", task, tasks.BACKGROUND_QUEUE.GetQueue())
	if err != nil {
		log.Err(err).Msg("Failed to register universe races task")
		return
	}
	log.Info().Any("entryID", entryID).Msg("Registered universe races task")
}

func scheduleUniverseCategoriesJob() {
	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseCategories(context.Background(), nil)
	if err != nil {
		log.Err(err).Msg("Failed to get universe categories")
		return
	}
	for _, r := range data {
		task, err := tasks.NewCronJobUniverseCategoriesTask(r)
		if err != nil {
			log.Err(err).Msgf("Failed to create universe categories task: %d", r)
		}
		entryID, err := queueClient.Enqueue(task, tasks.BACKGROUND_QUEUE.GetQueue())
		if err != nil {
			log.Err(err).Msgf("Failed to register universe categories task: %d", r)
		}
		log.Info().Any("entryID", entryID).Msgf("Registered universe categories task: %d", r)
	}
}
