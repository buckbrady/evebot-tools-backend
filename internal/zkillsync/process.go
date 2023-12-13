package zkillsync

import (
	"github.com/buckbrady/evebot-tools-backend/pkg/database/models"
	"github.com/kamva/mgm/v3"
	"github.com/rs/zerolog/log"
)

func ProcessKillmail(resp *KillmailResponse) {
	r := resp.Package
	var attackers []models.KillmailAttackers
	var items []models.KillmailItems
	for _, a := range r.Killmail.Attackers {
		attackers = append(attackers, models.KillmailAttackers{
			AllianceID:     a.AllianceID,
			CharacterID:    a.CharacterID,
			CorporationID:  a.CorporationID,
			DamageDone:     a.DamageDone,
			FactionID:      a.FactionID,
			FinalBlow:      a.FinalBlow,
			SecurityStatus: a.SecurityStatus,
			ShipTypeID:     a.ShipTypeID,
			WeaponTypeID:   a.WeaponTypeID,
		})
	}
	for _, i := range r.Killmail.Victim.Items {
		items = append(items, models.KillmailItems{
			Flag:       i.Flag,
			ItemTypeID: i.ItemTypeID,
			Items: []struct {
				Flag            int `json:"flag" bson:"flag"`
				ItemTypeID      int `json:"itemTypeID" bson:"itemTypeID"`
				QuantityDropped int `json:"quantityDropped" bson:"quantityDropped"`
				Singleton       int `json:"singleton" bson:"singleton"`
			}(i.Items),
			QuantityDestroyed: i.QuantityDestroyed,
			QuantityDropped:   i.QuantityDropped,
			Singleton:         i.Singleton,
		})
	}
	record := models.Killmail{
		KillmailID:    r.Killmail.KillmailID,
		KillmailTime:  r.Killmail.KillmailTime,
		SolarSystemID: r.Killmail.SolarSystemID,
		WarID:         r.Killmail.WarID,
		DamageTaken:   r.Killmail.Victim.DamageTaken,
		TotalValue:    r.Zkb.TotalValue,
		Victim: models.KillmailVictim{
			AllianceID:    r.Killmail.Victim.AllianceID,
			CharacterID:   r.Killmail.Victim.CharacterID,
			CorporationID: r.Killmail.Victim.CorporationID,
			FactionID:     r.Killmail.Victim.FactionID,
		},
		Items: items,
		Position: models.KillmailPosition{
			X: r.Killmail.Victim.Position.X,
			Y: r.Killmail.Victim.Position.Y,
			Z: r.Killmail.Victim.Position.Z,
		},
		Attackers: attackers,
		ZkillboardData: models.KillmailZkb{
			Awox:           r.Zkb.Awox,
			DestroyedValue: r.Zkb.DestroyedValue,
			DroppedValue:   r.Zkb.DroppedValue,
			FittedValue:    r.Zkb.FittedValue,
			Hash:           r.Zkb.Hash,
			Href:           r.Zkb.Href,
			Labels:         r.Zkb.Labels,
			LocationID:     r.Zkb.LocationID,
			Npc:            r.Zkb.Npc,
			Points:         r.Zkb.Points,
			Solo:           r.Zkb.Solo,
		},
	}
	err := mgm.Coll(&models.Killmail{}).Create(&record)
	if err != nil {
		log.Err(err).Msg("failed to create record")
	}
}
