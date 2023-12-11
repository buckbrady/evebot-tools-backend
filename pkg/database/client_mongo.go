package database

import (
	"github.com/buckbrady/evebot-tools-backend/pkg/utils"
	"github.com/kamva/mgm/v3"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func init() {
	mongoOpts := options.Client()
	mongoOpts.SetReadPreference(readpref.SecondaryPreferred())
	mongoOpts.ApplyURI(utils.GetEnv("MONGODB_URI", ""))
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 60 * time.Second}, utils.GetEnv("MONGODB_DATABASE", "evetools"), mongoOpts)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to mongodb")
	}
}
