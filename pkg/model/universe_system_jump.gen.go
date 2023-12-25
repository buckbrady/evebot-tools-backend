// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUniverseSystemJump = "universe_system_jump"

// UniverseSystemJump mapped from table <universe_system_jump>
type UniverseSystemJump struct {
	SystemID  int32     `gorm:"column:system_id;primaryKey" json:"system_id"`
	Timestamp time.Time `gorm:"column:timestamp;primaryKey;default:now()" json:"timestamp"`
	ShipJumps int32     `gorm:"column:ship_jumps;not null" json:"ship_jumps"`
}

// TableName UniverseSystemJump's table name
func (*UniverseSystemJump) TableName() string {
	return TableNameUniverseSystemJump
}
