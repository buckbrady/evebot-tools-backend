// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUniversePlanet = "universe_planet"

// UniversePlanet mapped from table <universe_planet>
type UniversePlanet struct {
	ID       int32 `gorm:"column:id;primaryKey" json:"id"`
	SystemID int32 `gorm:"column:system_id;not null" json:"system_id"`
}

// TableName UniversePlanet's table name
func (*UniversePlanet) TableName() string {
	return TableNameUniversePlanet
}
