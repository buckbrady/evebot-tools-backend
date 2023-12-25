package worker

import (
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	_ "github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/tasks"
	"github.com/buckbrady/evebot-tools-backend/pkg/utils"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"strconv"
)

func Run() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	redisDB, _ := strconv.Atoi(utils.GetEnv("QUEUE_REDIS_DB", "10"))
	redisOpts := asynq.RedisClientOpt{Addr: utils.GetEnv("REDIS_ADDR", "localhost:6379"), DB: redisDB}

	db := utils.NewPGConn()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	tasks.AddDB(db)
	tasks.AddQueueClient(asynq.NewClient(redisOpts))

	currencyCount, err := strconv.Atoi(utils.GetEnv("WORKER_PROCS", "25"))
	if err != nil {
		log.Fatal().Msgf("could not convert WORKER_PROCS to int: %v", err)
	}

	srv := asynq.NewServer(
		redisOpts,
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: currencyCount,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				tasks.ESI_STATUS_QUEUE.GetName():   tasks.ESI_STATUS_QUEUE.GetPriority(),
				tasks.ESI_UNIVERSE_QUEUE.GetName(): tasks.ESI_UNIVERSE_QUEUE.GetPriority(),
			},
			// See the godoc for other configuration options
		},
	)
	mux := asynq.NewServeMux()

	// ESI Status
	mux.HandleFunc(tasks.TypeCronJobEsiStatus, tasks.HandleCronJobStatusTask)
	// ESI Universe
	mux.HandleFunc(tasks.TypeCronJobEsiUniverseTypes, tasks.HandleCronJobUniverseTypesTask)
	mux.HandleFunc(tasks.TypeCronJobEsiUniverseRegions, tasks.HandleCronJobUniverseRegionsTask)
	mux.HandleFunc(tasks.TypeCronJobEsiUniverseConstellations, tasks.HandleCronJobUniverseConstellationsTask)
	mux.HandleFunc(tasks.TypeCronJobEsiUniverseSystems, tasks.HandleCronJobUniverseSystemsTask)
	mux.HandleFunc(tasks.TypeCronJobEsiUniverseStargates, tasks.HandleCronJobUniverseStargatesTask)
	mux.HandleFunc(tasks.TypeCronJobEsiUniversePlanets, tasks.HandleCronJobUniversePlanetsTask)
	mux.HandleFunc(tasks.TypeCronJobEsiUniverseMoons, tasks.HandleCronJobUniverseMoonsTask)
	mux.HandleFunc(tasks.TypeCronJobEsiUniverseAsteroidBelts, tasks.HandleCronJobUniverseAsteroidBeltsTask)
	mux.HandleFunc(tasks.TypeCronJobEsiUniverseStations, tasks.HandleCronJobUniverseStationsTask)
	mux.HandleFunc(tasks.TypeCronJobEsiUniverseStars, tasks.HandleCronJobUniverseStarsTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal().Msgf("could not run server: %v", err)
	}

	esi.EVECache.Close()

}
