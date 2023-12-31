// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameKillmail = "killmail"

// Killmail mapped from table <killmail>
type Killmail struct {
	ID            int64     `gorm:"column:id;primaryKey" json:"id"`
	KillmailTime  time.Time `gorm:"column:killmail_time;not null" json:"killmail_time"`
	SolarSystemID int32     `gorm:"column:solar_system_id;not null" json:"solar_system_id"`
	WarID         *int64    `gorm:"column:war_id" json:"war_id"`
	DamageTaken   int32     `gorm:"column:damage_taken;not null" json:"damage_taken"`
	TotalValue    float64   `gorm:"column:total_value;not null" json:"total_value"`
}

// TableName Killmail's table name
func (*Killmail) TableName() string {
	return TableNameKillmail
}
