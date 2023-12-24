package worker

import (
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	_ "github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/tasks"
	"github.com/buckbrady/evebot-tools-backend/pkg/utils"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"

	"strconv"
)

func Run() {
	redisDB, _ := strconv.Atoi(utils.GetEnv("QUEUE_REDIS_DB", "10"))
	redisOpts := asynq.RedisClientOpt{Addr: utils.GetEnv("REDIS_ADDR", "localhost:6379"), DB: redisDB}

	tasks.AddDB(utils.NewPGConn())

	srv := asynq.NewServer(
		redisOpts,
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 25,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				tasks.ESI_STATUS_QUEUE.GetName():   tasks.ESI_STATUS_QUEUE.GetPriority(),
				tasks.ESI_UNIVERSE_QUEUE.GetName(): tasks.ESI_UNIVERSE_QUEUE.GetPriority(),
			},
			// See the godoc for other configuration options
		},
	)
	mux := asynq.NewServeMux()

	mux.HandleFunc(tasks.TypeCronJobEsiStatus, tasks.HandleCronJobStatusTask)
	mux.HandleFunc(tasks.TypeCronJobEsiUniverseTypes, tasks.HandleCronJobUniverseTypesTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal().Msgf("could not run server: %v", err)
	}

	esi.EVECache.Close()

}
