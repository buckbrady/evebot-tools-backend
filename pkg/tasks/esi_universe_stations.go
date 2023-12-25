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
	TypeCronJobEsiUniverseStations = "cronjob:esi:universe:stations"
)

func NewCronJobUniverseStationsTask(typeID int32) (*asynq.Task, error) {
	payload, err := json.Marshal(CronJobUniverseTypesPayload{
		Timestamp: time.Now().UTC(),
		TTL:       86400,
		TypeID:    typeID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronJobEsiUniverseStations, payload), nil
}

func HandleCronJobUniverseStationsTask(ctx context.Context, t *asynq.Task) error {
	var p CronJobUniverseTypesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	if p.TypeID <= 1 {
		log.Error().Msg("invalid typeID")
		return nil
	}
	data, _, err := esi.EVE.ESI.UniverseApi.GetUniverseStationsStationId(context.Background(), p.TypeID, nil)
	if err != nil {
		return err
	}

	if data.StationId <= 1 {
		log.Error().Msg("invalid stationID")
		return nil
	}

	record := model.UniverseStation{
		ID:                       data.StationId,
		Name:                     data.Name,
		Owner:                    data.Owner,
		MaxDockableShipVolume:    float64(data.MaxDockableShipVolume),
		OfficeRentalCost:         float64(data.OfficeRentalCost),
		PositionX:                data.Position.X,
		PositionY:                data.Position.Y,
		PositionZ:                data.Position.Z,
		RaceID:                   data.RaceId,
		ReprocessingEfficiency:   float64(data.ReprocessingEfficiency),
		ReprocessingStationsTake: float64(data.ReprocessingStationsTake),
		SystemID:                 data.SystemId,
		TypeID:                   data.TypeId,
	}

	dbctx, cancel := utils.NewDBCtx(ctx, 60)
	defer cancel()
	err = database.Use(db).UniverseStation.WithContext(dbctx).Save(&record)
	if err != nil {
		return err
	}

	for _, service := range data.Services {
		record := model.UniverseStationService{
			StationID: data.StationId,
			Service:   service,
		}
		dbctx, cancel := utils.NewDBCtx(ctx, 60)
		defer cancel()
		err = database.Use(db).UniverseStationService.WithContext(dbctx).Save(&record)
		if err != nil {
			return err
		}
	}

	return nil
}
