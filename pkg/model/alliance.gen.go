// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAlliance = "alliance"

// Alliance mapped from table <alliance>
type Alliance struct {
	ID                    int32     `gorm:"column:id;primaryKey" json:"id"`
	CreatorCorporationID  int32     `gorm:"column:creator_corporation_id;not null" json:"creator_corporation_id"`
	CreatorID             int32     `gorm:"column:creator_id;not null" json:"creator_id"`
	DateFounded           time.Time `gorm:"column:date_founded;not null" json:"date_founded"`
	ExecutorCorporationID *int32    `gorm:"column:executor_corporation_id" json:"executor_corporation_id"`
	FactionID             *int32    `gorm:"column:faction_id" json:"faction_id"`
	Name                  string    `gorm:"column:name;not null" json:"name"`
	Ticker                string    `gorm:"column:ticker;not null" json:"ticker"`
}

// TableName Alliance's table name
func (*Alliance) TableName() string {
	return TableNameAlliance
}
