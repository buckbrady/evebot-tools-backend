package esi

import (
	"fmt"
	"github.com/buckbrady/evebot-tools-backend/pkg/utils"
	"net/http"
	"time"

	"github.com/antihax/goesi"
	"github.com/gomodule/redigo/redis"
	"github.com/gregjones/httpcache"
	httpredis "github.com/gregjones/httpcache/redis"
	"github.com/rs/zerolog/log"
)

var (
	EVE      *goesi.APIClient
	EVECache redis.Conn
)

func init() {
	EVE, EVECache = initEsiClient()
}

func initEsiClient() (*goesi.APIClient, redis.Conn) {
	cacheConn, err := redis.Dial("tcp", utils.GetEnv("REDIS_ADDR", "localhost:6379"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to redis")
	}

	transport := httpcache.NewTransport(httpredis.NewWithClient(cacheConn))
	transport.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
	httpClient := &http.Client{Transport: transport, Timeout: time.Second * 30}
	appName := utils.GetEnv("ESI_APP_NAME", "evebot-tools")
	appContact := utils.GetEnv("ESI_APP_CONTACT", "")
	esi := goesi.NewAPIClient(httpClient, fmt.Sprintf("%s <%s>", appName, appContact))

	return esi, cacheConn
}
