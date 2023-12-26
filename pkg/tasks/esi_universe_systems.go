package tasks

import (
	"context"
	"encoding/json"
	"github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/model"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	TypeCronJobEsiUniverseSystems = "cronjob:esi:universe:systems"
)

func NewCronJobUniverseSystemsTask(id int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobPayloadWithType{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    id,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseSystems, payload), nil
}

func HandleCronJobUniverseSystemsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobPayloadWithType
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseSystemsSystemId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	record := model.UniverseSystem{
		ID:              data.SystemId,
		Name:            data.Name,
		SecurityClass:   data.SecurityClass,
		SecurityStatus:  float64(data.SecurityStatus),
		ConstellationID: data.ConstellationId,
		StarID:          data.StarId,
		PositionX:       data.Position.X,
		PositionY:       data.Position.Y,
		PositionZ:       data.Position.Z,
	}
	err = database.Use(db).UniverseSystem.WithContext(ctx).Save(&record)
	if err != nil {
		log.Err(err).Any("systemID", data.SystemId).Msg("failed to create universe system record")
		return err
	}

	taskStar, err := NewCronJobUniverseStarsTask(data.StarId)
	if err != nil {
		return err
	}
	entryIDStar, err := queueClient.Enqueue(taskStar, ESI_BACKGROUND.GetQueue())
	if err != nil {
		return err
	}
	log.Info().Any("entryID", entryIDStar).Any("starID", data.StarId).Msg("registered universe star task")

	for _, planet := range data.Planets {
		task, err := NewCronJobUniversePlanetsTask(planet.PlanetId)
		if err != nil {
			log.Err(err).Any("planetID", planet.PlanetId).Msg("failed to create universe planet task")
			return err
		}
		entryID, err := queueClient.Enqueue(task, ESI_BACKGROUND.GetQueue())
		if err != nil {
			log.Err(err).Any("planetID", planet.PlanetId).Msg("failed to register universe planet task")
			return err
		}
		log.Info().Any("entryID", entryID).Any("planetID", planet.PlanetId).Msg("registered universe planet task")

		for _, asteroidBelt := range planet.AsteroidBelts {
			taskBelt, err := NewCronJobUniverseAsteroidBeltsTask(asteroidBelt)
			if err != nil {
				log.Err(err).Any("asteroidBeltID", asteroidBelt).Msg("failed to create universe asteroid belt task")
				return err
			}
			entryIDBelt, err := queueClient.Enqueue(taskBelt, ESI_BACKGROUND.GetQueue())
			if err != nil {
				log.Err(err).Any("asteroidBeltID", asteroidBelt).Msg("failed to register universe asteroid belt task")
				return err
			}
			log.Info().Any("entryID", entryIDBelt).Any("asteroidBeltID", asteroidBelt).Msg("registered universe asteroid belt task")
		}

		for _, moon := range planet.Moons {
			taskMoon, err := NewCronJobUniverseMoonsTask(moon)
			if err != nil {
				log.Err(err).Any("moonID", moon).Msg("failed to create universe moon task")
				return err
			}
			entryIDMoon, err := queueClient.Enqueue(taskMoon, ESI_BACKGROUND.GetQueue())
			if err != nil {
				log.Err(err).Any("moonID", moon).Msg("failed to register universe moon task")
				return err
			}
			log.Info().Any("entryID", entryIDMoon).Any("moonID", moon).Msg("registered universe moon task")
		}
	}

	for _, stargate := range data.Stargates {
		taskStargate, err := NewCronJobUniverseStargatesTask(stargate)
		if err != nil {
			log.Err(err).Any("stargateID", stargate).Msg("failed to create universe stargate task")
			return err
		}
		entryIDStargate, err := queueClient.Enqueue(taskStargate, ESI_BACKGROUND.GetQueue())
		if err != nil {
			log.Err(err).Any("stargateID", stargate).Msg("failed to register universe stargate task")
			return err
		}
		log.Info().Any("entryID", entryIDStargate).Any("stargateID", stargate).Msg("registered universe stargate task")
	}

	for _, station := range data.Stations {
		taskStation, err := NewCronJobUniverseStationsTask(station)
		if err != nil {
			log.Err(err).Any("stationID", station).Msg("failed to create universe station task")
			return err
		}
		entryIDStation, err := queueClient.Enqueue(taskStation, ESI_BACKGROUND.GetQueue())
		if err != nil {
			log.Err(err).Any("stationID", station).Msg("failed to register universe station task")
			return err
		}
		log.Info().Any("entryID", entryIDStation).Any("stationID", station).Msg("registered universe station task")
	}

	return nil
}
