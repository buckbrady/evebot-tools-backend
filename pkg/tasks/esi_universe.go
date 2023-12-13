package tasks

import (
	"context"
	"encoding/json"
	"github.com/buckbrady/evebot-tools-backend/pkg/database/models"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/utils"
	"github.com/hibiken/asynq"
	"github.com/kamva/mgm/v3"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		TTL:       30,
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
	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseTypesTypeId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}
	// find record or create
	existingRecord := &models.UniverseType{}
	err = mgm.Coll(&models.UniverseType{}).First(bson.M{"typeID": data.TypeId}, existingRecord)
	if err != nil {
		log.Info().Any("typeID", data.TypeId).Msg("failed to find universe type record. creating new record")
		record := models.UniverseType{
			TypeID:         data.TypeId,
			Name:           data.Name,
			Description:    data.Description,
			Capacity:       &data.Capacity,
			GraphicID:      &data.GraphicId,
			GroupID:        data.GroupId,
			IconID:         &data.IconId,
			MarketGroupID:  &data.MarketGroupId,
			Mass:           &data.Mass,
			PackagedVolume: &data.PackagedVolume,
			PortionSize:    &data.PortionSize,
			Published:      data.Published,
			Radius:         &data.Radius,
			Volume:         &data.Volume,
		}
		err = mgm.Coll(&models.UniverseType{}).Create(&record)
		if err != nil {
			log.Err(err).Msg("failed to create universe type record")
			return err
		}
	} else {
		log.Info().Any("typeID", data.TypeId).Msg("found existing universe type record. updating")
		existingRecord.IconID = &data.IconId
		existingRecord.GraphicID = &data.GraphicId
		existingRecord.Capacity = &data.Capacity
		existingRecord.Description = data.Description
		existingRecord.GroupID = data.GroupId
		existingRecord.MarketGroupID = &data.MarketGroupId
		existingRecord.Mass = &data.Mass
		existingRecord.Name = data.Name
		existingRecord.PackagedVolume = &data.PackagedVolume
		existingRecord.PortionSize = &data.PortionSize
		existingRecord.Published = data.Published
		existingRecord.Radius = &data.Radius
		existingRecord.Volume = &data.Volume
		err = mgm.Coll(&models.UniverseType{}).Update(existingRecord, &options.UpdateOptions{Upsert: utils.NewBool(true)})
		if err != nil {
			log.Err(err).Msg("failed to update universe type record")
			return err
		}
	}

	return nil
}
