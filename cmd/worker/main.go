package main

import (
	"fmt"
	"github.com/buckbrady/evebot-tools-backend/internal/worker"
	"github.com/buckbrady/evebot-tools-backend/pkg/utils"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"time"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              utils.GetEnv("SENTRY_DSN", ""),
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	defer sentry.Flush(5 * time.Second)
	worker.Run()
}
