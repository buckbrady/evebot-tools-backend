// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUniverseMoon = "universe_moon"

// UniverseMoon mapped from table <universe_moon>
type UniverseMoon struct {
	ID       int32 `gorm:"column:id;primaryKey" json:"id"`
	SystemID int32 `gorm:"column:system_id;not null" json:"system_id"`
}

// TableName UniverseMoon's table name
func (*UniverseMoon) TableName() string {
	return TableNameUniverseMoon
}
