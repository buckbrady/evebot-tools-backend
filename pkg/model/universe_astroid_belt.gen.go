// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUniverseAstroidBelt = "universe_astroid_belt"

// UniverseAstroidBelt mapped from table <universe_astroid_belt>
type UniverseAstroidBelt struct {
	ID        int32   `gorm:"column:id;primaryKey" json:"id"`
	SystemID  int32   `gorm:"column:system_id;not null" json:"system_id"`
	PositionX float64 `gorm:"column:position_x;not null" json:"position_x"`
	PositionY float64 `gorm:"column:position_y;not null" json:"position_y"`
	PositionZ float64 `gorm:"column:position_z;not null" json:"position_z"`
	Name      string  `gorm:"column:name;not null" json:"name"`
}

// TableName UniverseAstroidBelt's table name
func (*UniverseAstroidBelt) TableName() string {
	return TableNameUniverseAstroidBelt
}
