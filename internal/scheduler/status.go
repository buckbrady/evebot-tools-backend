package scheduler

import (
	"github.com/buckbrady/evebot-tools-backend/pkg/tasks"
	"github.com/rs/zerolog/log"
)

func scheduleStatusJobs() {
	statusTask, err := tasks.NewCronJobStatusTask()
	if err != nil {
		log.Err(err).Msg("failed to create status task")
	}
	statusID, err := scheduler.Register("@every 30s", statusTask, tasks.STANDARD_QUEUE.GetQueue())
	if err != nil {
		log.Err(err).Msg("failed to register status task")
	} else {
		log.Info().Str("entryID", statusID).Msg("registered status task")
	}
}
