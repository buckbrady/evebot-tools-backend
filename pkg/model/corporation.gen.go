// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCorporation = "corporation"

// Corporation mapped from table <corporation>
type Corporation struct {
	ID            int32     `gorm:"column:id;primaryKey" json:"id"`
	AllianceID    *int32    `gorm:"column:alliance_id" json:"alliance_id"`
	CeoID         int32     `gorm:"column:ceo_id;not null" json:"ceo_id"`
	CreatorID     int32     `gorm:"column:creator_id;not null" json:"creator_id"`
	DateFounded   time.Time `gorm:"column:date_founded;not null" json:"date_founded"`
	Description   string    `gorm:"column:description;not null" json:"description"`
	FactionID     *int32    `gorm:"column:faction_id" json:"faction_id"`
	HomeStationID *int32    `gorm:"column:home_station_id" json:"home_station_id"`
	MemberCount   int32     `gorm:"column:member_count;not null" json:"member_count"`
	Name          string    `gorm:"column:name;not null" json:"name"`
	Shares        int32     `gorm:"column:shares;not null" json:"shares"`
	TaxRate       float64   `gorm:"column:tax_rate;not null" json:"tax_rate"`
	Ticker        string    `gorm:"column:ticker;not null" json:"ticker"`
	URL           *string   `gorm:"column:url" json:"url"`
	WarEligible   bool      `gorm:"column:war_eligible;not null" json:"war_eligible"`
}

// TableName Corporation's table name
func (*Corporation) TableName() string {
	return TableNameCorporation
}
