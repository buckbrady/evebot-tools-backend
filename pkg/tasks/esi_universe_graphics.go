package tasks

import (
	"context"
	"encoding/json"
	"github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/buckbrady/evebot-tools-backend/pkg/esi"
	"github.com/buckbrady/evebot-tools-backend/pkg/model"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypeCronJobEsiUniverseGraphics = "cronjob:esi:universe:graphics"
)

func NewCronJobUniverseGraphicsTask(graphicID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobPayloadWithType{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    graphicID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseGraphics, payload), nil
}

func HandleCronJobUniverseGraphicsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobPayloadWithType
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	resp, _, err := esi.EVE.ESI.UniverseApi.GetUniverseGraphicsGraphicId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	record := model.UniverseGraphic{
		ID:            resp.GraphicId,
		CollisionFile: &resp.CollisionFile,
		GraphicFile:   &resp.GraphicFile,
		IconFolder:    &resp.IconFolder,
		SofDna:        &resp.SofDna,
		SofFationName: &resp.SofFationName,
		SofHullName:   &resp.SofHullName,
		SofRaceName:   &resp.SofRaceName,
	}

	err = database.Use(db).UniverseGraphic.WithContext(ctx).Save(&record)

	return nil
}
