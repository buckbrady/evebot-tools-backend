package tasks

import (
	"context"
	"encoding/json"
	"github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/buckbrady/evebot-tools-backend/pkg/model"
	"github.com/buckbrady/evebot-tools-backend/pkg/utils"

	//"github.com/buckbrady/evebot-tools-backend/pkg/database/models"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	TypeCronJobEsiUniverseTypes = "cronjob:esi:universe:types"
)

type CronJobUniverseTypesPayload struct {
	Timestamp time.Time
	TTL       int
	TypeID    int32
}

// Tasks

func NewCronJobUniverseTypesTask(typeID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    typeID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseTypes, payload), nil
}

// Handlers

func HandleCronJobUniverseTypesTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobUniverseTypesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	if p.TypeID < 1 {
		log.Error().Msg("invalid typeID")
		return nil
	}
	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseTypesTypeId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	//log.Info().Any("typeID", data.TypeId).Msg("failed to find universe type record. creating new record")

	if data.TypeId < 1 {
		log.Error().Msg("invalid typeID")
		return nil
	}

	record := model.UniverseType{
		ID:             data.TypeId,
		Name:           data.Name,
		Description:    data.Description,
		Capacity:       float64(data.Capacity),
		GraphicID:      &data.GraphicId,
		GroupID:        data.GroupId,
		IconID:         &data.IconId,
		MarketGroupID:  &data.MarketGroupId,
		Mass:           float64(data.Mass),
		PackagedVolume: float64(data.PackagedVolume),
		PortionSize:    data.PortionSize,
		Published:      data.Published,
		Radius:         float64(data.Radius),
		Volume:         float64(data.Volume),
	}
	dbctx, cancel := utils.NewDBCtx(ctx, 60)
	defer cancel()
	err = database.Use(db).UniverseType.WithContext(dbctx).Save(&record)
	if err != nil {
		log.Err(err).Msg("failed to create universe type record")
		return err
	}

	for _, effect := range data.DogmaEffects {
		dogmaEffects := model.UniverseTypeDogmaEffect{
			TypeID:    data.TypeId,
			EffectID:  effect.EffectId,
			IsDefault: effect.IsDefault,
		}
		dbctx1, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		err = database.Use(db).UniverseTypeDogmaEffect.WithContext(dbctx1).Save(&dogmaEffects)
		if err != nil {
			log.Err(err).Any("typeID", data.TypeId).Any("effectID", effect.EffectId).Msg("failed to create universe type dogma effect record")
			return err
		}
	}

	for _, attribute := range data.DogmaAttributes {
		dogmaAttributes := model.UniverseTypeDogmaAttribute{
			TypeID:      data.TypeId,
			AttributeID: attribute.AttributeId,
			Value:       float64(attribute.Value),
		}
		dbctx1, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		err = database.Use(db).UniverseTypeDogmaAttribute.WithContext(dbctx1).Save(&dogmaAttributes)
		if err != nil {
			log.Err(err).Any("typeID", data.TypeId).Any("attributeID", attribute.AttributeId).Msg("failed to create universe type dogma attribute record")
			return err
		}
	}

	return nil
}
