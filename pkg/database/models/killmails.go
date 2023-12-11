package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)

type Killmail struct {
	mgm.DefaultModel `bson:",inline"`
	KillmailID       int                 `json:"killmailID" bson:"killmailID"`
	KillmailTime     time.Time           `json:"killmailTime" bson:"killmailTime"`
	SolarSystemID    int                 `json:"solarSystemID" bson:"solarSystemID"`
	WarID            *int                `json:"warID,omitempty" bson:"warID,omitempty"`
	Attackers        []KillmailAttackers `json:"attackers" bson:"attackers"`
	Victim           KillmailVictim      `json:"victim" bson:"victim"`
	Items            []KillmailItems     `json:"items" bson:"items"`
	Position         KillmailPosition    `json:"position" bson:"position"`
	DamageTaken      int                 `json:"damageTaken" bson:"damageTaken"`
	TotalValue       float64             `json:"totalValue" bson:"totalValue"`
	ZkillboardData   KillmailZkb         `json:"zkillboardData" bson:"zkillboardData"`
}

type KillmailAttackers struct {
	AllianceID     *int    `json:"allianceID,omitempty" bson:"allianceID,omitempty"`
	CharacterID    *int    `json:"characterID,omitempty" bson:"characterID,omitempty"`
	CorporationID  *int    `json:"corporationID,omitempty" bson:"corporationID,omitempty"`
	DamageDone     int     `json:"damageDone" bson:"damageDone"`
	FactionID      *int    `json:"factionID,omitempty" bson:"factionID,omitempty"`
	FinalBlow      bool    `json:"finalBlow" bson:"finalBlow"`
	SecurityStatus float64 `json:"securityStatus" bson:"securityStatus"`
	ShipTypeID     *int    `json:"shipTypeID,omitempty" bson:"shipTypeID,omitempty"`
	WeaponTypeID   *int    `json:"weaponTypeID,omitempty" bson:"weaponTypeID,omitempty"`
}

type KillmailVictim struct {
	AllianceID    *int `json:"allianceID,omitempty" bson:"allianceID,omitempty"`
	CharacterID   int  `json:"characterID" bson:"characterID"`
	CorporationID int  `json:"corporationID" bson:"corporationID"`
	FactionID     *int `json:"factionID,omitempty" bson:"factionID,omitempty"`
}

type KillmailItems struct {
	Flag       int `json:"flag" bson:"flag"`
	ItemTypeID int `json:"itemTypeID" bson:"itemTypeID"`
	Items      []struct {
		Flag            int `json:"flag" bson:"flag"`
		ItemTypeID      int `json:"itemTypeID" bson:"itemTypeID"`
		QuantityDropped int `json:"quantityDropped" bson:"quantityDropped"`
		Singleton       int `json:"singleton" bson:"singleton"`
	} `json:"items,omitempty" bson:"items,omitempty"`
	QuantityDestroyed *int `json:"quantityDestroyed,omitempty" bson:"quantityDestroyed"`
	QuantityDropped   *int `json:"quantityDropped,omitempty" bson:"quantityDropped"`
	Singleton         int  `json:"singleton" bson:"singleton"`
}

type KillmailPosition struct {
	X float64 `json:"x" bson:"x"`
	Y float64 `json:"y" bson:"y"`
	Z float64 `json:"z" bson:"z"`
}

type KillmailZkb struct {
	Awox           bool     `json:"awox" bson:"awox"`
	DestroyedValue float64  `json:"destroyedValue" bson:"destroyedValue"`
	DroppedValue   float64  `json:"droppedValue" bson:"droppedValue"`
	FittedValue    float64  `json:"fittedValue" bson:"fittedValue"`
	Hash           string   `json:"hash" bson:"hash"`
	Href           string   `json:"href" bson:"href"`
	Labels         []string `json:"labels" bson:"labels"`
	LocationID     int      `json:"locationID" bson:"locationID"`
	Npc            bool     `json:"npc" bson:"npc"`
	Points         int      `json:"points" bson:"points"`
	Solo           bool     `json:"solo" bson:"solo"`
}
