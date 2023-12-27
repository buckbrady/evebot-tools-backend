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
		task, err := queueClient.Enqueue(t, tasks.STANDARD_QUEUE.GetQueue())
		writeResponse(w, fmt.Sprintf("enqueued task %s", task.ID), err)

	})

	r.Get("/universe/races", func(w http.ResponseWriter, r *http.Request) {
		t, err := tasks.NewCronJubUniverseRacesTask()
		if err != nil {
			log.Err(err).Msg("failed to create races task")
			return
		}
		task, err := queueClient.Enqueue(t, tasks.BACKGROUND_QUEUE.GetQueue())
		writeResponse(w, fmt.Sprintf("enqueued task %s", task.ID), err)

	})

	r.Get("/universe/ancestries", func(w http.ResponseWriter, r *http.Request) {
		t, err := tasks.NewCronJobUniverseAncestriesTask()
		if err != nil {
			log.Err(err).Msg("failed to create ancestries task")
			return
		}
		task, err := queueClient.Enqueue(t, tasks.BACKGROUND_QUEUE.GetQueue())
		writeResponse(w, fmt.Sprintf("enqueued task %s", task.ID), err)

	})

	r.Get("/universe/factions", func(w http.ResponseWriter, r *http.Request) {
		t, err := tasks.NewCronJobUniverseFactionsTask()
		if err != nil {
			log.Err(err).Msg("failed to create factions task")
			return
		}
		task, err := queueClient.Enqueue(t, tasks.BACKGROUND_QUEUE.GetQueue())
		writeResponse(w, fmt.Sprintf("enqueued task %s", task.ID), err)

	})

	r.Get("/universe/bloodlines", func(w http.ResponseWriter, r *http.Request) {
		t, err := tasks.NewCronJobUniverseBloodlinesTask()
		if err != nil {
			log.Err(err).Msg("failed to create bloodlines task")
			return
		}
		task, err := queueClient.Enqueue(t, tasks.BACKGROUND_QUEUE.GetQueue())
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
	r.Get("/universe/graphics", func(w http.ResponseWriter, r *http.Request) {
		scheduleUniverseGraphicsJob()
		writeResponse(w, fmt.Sprintf("enqueued task"), nil)

	})
	r.Get("/universe/categories", func(w http.ResponseWriter, r *http.Request) {
		scheduleUniverseCategoriesJob()
		writeResponse(w, fmt.Sprintf("enqueued task"), nil)

	})

	// Market
	r.Get("/market/groups", func(w http.ResponseWriter, r *http.Request) {
		scheduleMarketGroupsJob()
		writeResponse(w, fmt.Sprintf("enqueued task"), nil)

	})
	r.Get("/market/history", func(w http.ResponseWriter, r *http.Request) {
		scheduleMarketHistoryTask()
		writeResponse(w, fmt.Sprintf("enqueued task"), nil)

	})
	r.Get("/market/orders", func(w http.ResponseWriter, r *http.Request) {
		scheduleMarketOrdersTask()
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
