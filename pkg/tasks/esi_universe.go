package tasks

import (
	"context"
	"encoding/json"
	"github.com/buckbrady/evebot-tools-backend/pkg/database/models"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/hibiken/asynq"
	"github.com/kamva/mgm/v3"
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

func HandleCronJobUniverseTypesTask(t *asynq.Task) error {
	var p CronJobUniverseTypesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseTypesTypeId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}
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
		return err
	}
	return nil
}
