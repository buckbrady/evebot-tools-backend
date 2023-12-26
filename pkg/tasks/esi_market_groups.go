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
	TypeCronJobEsiMarketGroups = "cronjob:esi:market:groups"
)

func NewCronJobMarketGroupsTask(groupID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobPayloadWithType{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    groupID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiMarketGroups, payload), nil
}

func HandleCronJobMarketGroupsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobPayloadWithType
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	if p.TypeID < 1 {
		return nil
	}
	data, _, err := esi.EVE.ESI.MarketApi.GetMarketsGroupsMarketGroupId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}
	record := model.MarketGroup{
		ID:            data.MarketGroupId,
		Name:          data.Name,
		Description:   data.Description,
		ParentGroupID: &data.ParentGroupId,
	}
	err = database.Use(db).MarketGroup.WithContext(ctx).Save(&record)
	return err
}
