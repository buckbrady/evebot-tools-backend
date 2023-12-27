package scheduler

import (
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	_ "github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/tasks"
	"github.com/buckbrady/evebot-tools-backend/pkg/utils"
	"github.com/gomodule/redigo/redis"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

var (
	scheduler   *asynq.Scheduler
	queueClient *asynq.Client
	apiServer   *http.Server
	eveCache    redis.Conn = esi.EVECache
	db          *gorm.DB
)

func Run() {
	log.Info().Msg("Configuring scheduler")
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load time location")
	}
	redisDB, _ := strconv.Atoi(utils.GetEnv("QUEUE_REDIS_DB", "10"))
	redisOpts := asynq.RedisClientOpt{Addr: utils.GetEnv("REDIS_ADDR", "localhost:6379"), DB: redisDB}
	db = utils.NewPGConn()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	tasks.AddDB(db)

	scheduler = asynq.NewScheduler(redisOpts, &asynq.SchedulerOpts{Location: loc})
	queueClient = asynq.NewClient(redisOpts)

	go startApi()

	// Status Jobs
	scheduleStatusJobs()
	// Universe Jobs
	scheduleUniverseAncestriesJob()
	scheduleUniverseBloodlinesJob()
	scheduleUniverseSystemJumpsJob()
	scheduleUniverseSystemKillsJob()
	scheduleUniverseFactionsJob()
	scheduleUniverseRacesJob()
	// Market Jobs
	scheduleMarketPricesJob()

	if err := scheduler.Run(); err != nil {
		log.Fatal().Err(err)
	}

	//<-quit
	log.Info().Msg("SIGINT received, exiting")
	eveCache.Close()
	apiServer.Close()
}
