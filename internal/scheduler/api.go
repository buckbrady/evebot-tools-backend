package scheduler

import (
	"fmt"
	"github.com/buckbrady/evebot-tools-backend/pkg/tasks"
	"github.com/buckbrady/evebot-tools-backend/pkg/utils"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	"net/http"
)

func startApi() {
	sentryMiddleware := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	r := chi.NewRouter()
	apiServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", utils.GetEnv("PORT", "3333")),
		Handler: r,
	}

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.StripSlashes)
	r.Use(sentryMiddleware.Handle)

	r.Get("/status/status", func(w http.ResponseWriter, r *http.Request) {
		t, err := tasks.NewCronJobStatusTask()
		if err != nil {
			log.Err(err).Msg("failed to create status task")
			return
		}
		task, err := queueClient.Enqueue(t, tasks.ESI_STATUS_QUEUE.GetQueue())
		writeResponse(w, fmt.Sprintf("enqueued task %s", task.ID), err)

	})

	r.Get("/universe/types", func(w http.ResponseWriter, r *http.Request) {
		scheduleUniverseTypesJob()
		writeResponse(w, fmt.Sprintf("enqueued task"), nil)

	})
	r.Get("/universe/regions", func(w http.ResponseWriter, r *http.Request) {
		scheduleUniverseRegionsJob()
		writeResponse(w, fmt.Sprintf("enqueued task"), nil)

	})

	log.Info().Msg("starting webserver")
	err := apiServer.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start webserver")
	}
}

func writeResponse(w http.ResponseWriter, msg string, err error) {
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
	} else {
		_, _ = w.Write([]byte(msg))
	}
}
