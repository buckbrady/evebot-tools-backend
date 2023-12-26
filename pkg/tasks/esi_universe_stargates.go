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
	TypeCronJobEsiUniverseStargates = "cronjob:esi:universe:stargates"
)

func NewCronJobUniverseStargatesTask(typeID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobPayloadWithType{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    typeID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseStargates, payload), nil
}

func HandleCronJobUniverseStargatesTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobPayloadWithType
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	if p.TypeID < 1 {
		log.Error().Msg("invalid typeID")
		return nil
	}

	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseStargatesStargateId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	if data.StargateId < 1 {
		log.Error().Msg("invalid stargateID")
		return nil
	}

	record := model.UniverseStargate{
		ID:                    data.StargateId,
		Name:                  data.Name,
		SystemID:              data.SystemId,
		TypeID:                data.TypeId,
		DestinationStargateID: data.Destination.StargateId,
		DestinationSystemID:   data.Destination.SystemId,
		PositionX:             data.Position.X,
		PositionY:             data.Position.Y,
		PositionZ:             data.Position.Z,
	}
	err = database.Use(db).UniverseStargate.WithContext(ctx).Save(&record)
	if err != nil {
		log.Err(err).Msg("failed to create universe stargate record")
		return err
	}

	return nil
}
