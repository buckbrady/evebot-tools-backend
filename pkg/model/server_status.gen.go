// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameServerStatus = "server_status"

// ServerStatus mapped from table <server_status>
type ServerStatus struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Players       int32     `gorm:"column:players;not null" json:"players"`
	ServerVersion string    `gorm:"column:server_version;not null" json:"server_version"`
	StartTime     time.Time `gorm:"column:start_time;not null" json:"start_time"`
	Vip           bool      `gorm:"column:vip;not null" json:"vip"`
}

// TableName ServerStatus's table name
func (*ServerStatus) TableName() string {
	return TableNameServerStatus
}
