// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUniverseSystem = "universe_system"

// UniverseSystem mapped from table <universe_system>
type UniverseSystem struct {
	ID              int32    `gorm:"column:id;primaryKey" json:"id"`
	Name            string   `gorm:"column:name;not null" json:"name"`
	SecurityClass   string   `gorm:"column:security_class;not null" json:"security_class"`
	SecurityStatus  *float64 `gorm:"column:security_status" json:"security_status"`
	ConstellationID *int32   `gorm:"column:constellation_id" json:"constellation_id"`
	StarID          *int32   `gorm:"column:star_id" json:"star_id"`
}

// TableName UniverseSystem's table name
func (*UniverseSystem) TableName() string {
	return TableNameUniverseSystem
}
