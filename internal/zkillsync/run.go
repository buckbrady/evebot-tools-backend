package zkillsync

import (
	_ "github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

const (
	zkillboardURL = "https://redisq.zkillboard.com/listen.php?queueID=evebot-tools-prod"
)

func init() {
	log.Info().Msg("Initializing zkillsync")
}

func Run() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	client := req.C()
	for {
		select {
		case <-quit:
			log.Fatal().Msg("SIGINT received, exiting")
		default:
			resp, err := client.R().Get(zkillboardURL)
			if err != nil {
				log.Err(err).Msg("failed to get killmail api response")
			}
			var data KillmailResponse
			err = resp.UnmarshalJson(&data)
			if err != nil {
				log.Err(err).Msg("failed to unmarshal json")
			} else {
				if data == (KillmailResponse{}) {
					log.Info().Msg("Timeout reached. No killmails to process.")
				} else {
					go ProcessKillmail(&data)
				}
			}
		}
	}
}
