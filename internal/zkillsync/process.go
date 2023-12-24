package zkillsync

import (
	"context"
	"github.com/buckbrady/evebot-tools-backend/pkg/database"
	"github.com/buckbrady/evebot-tools-backend/pkg/model"
	"github.com/buckbrady/evebot-tools-backend/pkg/utils"
	"github.com/rs/zerolog/log"
)

func ProcessKillmail(resp *KillmailResponse) {
	ctx := context.Background()
	r := resp.Package
	var attackers []*model.KillmailAttacker
	var items []*model.KillmailItem
	km := model.Killmail{
		ID:            r.Killmail.KillmailID,
		KillmailTime:  r.Killmail.KillmailTime,
		SolarSystemID: r.Killmail.SolarSystemID,
		WarID:         r.Killmail.WarID,
		DamageTaken:   r.Killmail.Victim.DamageTaken,
		TotalValue:    r.Zkb.TotalValue,
	}
	for _, a := range r.Killmail.Attackers {
		a := model.KillmailAttacker{
			KillmailID:     r.Killmail.KillmailID,
			CharacterID:    a.CharacterID,
			CorporationID:  a.CorporationID,
			AllianceID:     a.AllianceID,
			FactionID:      a.FactionID,
			DamageDone:     a.DamageDone,
			FinalBlow:      a.FinalBlow,
			SecurityStatus: a.SecurityStatus,
			WeaponTypeID:   a.WeaponTypeID,
			ShipTypeID:     a.ShipTypeID,
		}
		attackers = append(attackers, &a)
	}
	victim := model.KillmailVictim{
		ID:            r.Killmail.KillmailID,
		CharacterID:   r.Killmail.Victim.CharacterID,
		CorporationID: r.Killmail.Victim.CorporationID,
		AllianceID:    r.Killmail.Victim.AllianceID,
		FactionID:     r.Killmail.Victim.FactionID,
	}
	for _, i := range r.Killmail.Victim.Items {
		item := model.KillmailItem{
			KillmailID:        r.Killmail.KillmailID,
			TypeID:            i.ItemTypeID,
			Flag:              i.Flag,
			QuantityDestroyed: i.QuantityDestroyed,
			QuantityDropped:   i.QuantityDropped,
			Singleton:         utils.IntToBool(i.Singleton),
		}
		items = append(items, &item)
	}
	err := database.Use(db).Killmail.WithContext(ctx).Create(&km)
	if err != nil {
		log.Err(err).Msg("failed to create killmail")
		return
	}
	err = database.Use(db).KillmailVictim.WithContext(ctx).Create(&victim)
	if err != nil {
		log.Err(err).Msg("failed to create victim")
	}
	err = database.Use(db).KillmailAttacker.WithContext(ctx).CreateInBatches(attackers, 100)
	if err != nil {
		log.Err(err).Msg("failed to create attackers")
	}
	err = database.Use(db).KillmailItem.WithContext(ctx).CreateInBatches(items, 100)
	if err != nil {
		log.Err(err).Msg("failed to create items")
	}
}
