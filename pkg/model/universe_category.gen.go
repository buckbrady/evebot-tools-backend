// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUniverseCategory = "universe_category"

// UniverseCategory mapped from table <universe_category>
type UniverseCategory struct {
	ID        int32  `gorm:"column:id;primaryKey" json:"id"`
	Name      string `gorm:"column:name;not null" json:"name"`
	Published bool   `gorm:"column:published;not null" json:"published"`
}

// TableName UniverseCategory's table name
func (*UniverseCategory) TableName() string {
	return TableNameUniverseCategory
}