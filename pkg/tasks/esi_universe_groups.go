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
	TypeCronJobEsiUniverseGroups = "cronjob:esi:universe:groups"
)

func NewCronJobUniverseGroupsTask(groupID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    groupID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseGroups, payload), nil
}

func HandleCronJobUniverseGroupsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobUniverseTypesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	if p.TypeID < 1 {
		return nil
	}

	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseGroupsGroupId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	record := model.UniverseGroup{
		ID:         data.GroupId,
		Name:       data.Name,
		Published:  data.Published,
		CategoryID: data.CategoryId,
	}
	err = database.Use(db).UniverseGroup.WithContext(ctx).Save(&record)
	if err != nil {
		return err
	}
	return nil
}
