// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCorporationIcon = "corporation_icon"

// CorporationIcon mapped from table <corporation_icon>
type CorporationIcon struct {
	ID        int32   `gorm:"column:id;primaryKey" json:"id"`
	Px128x128 *string `gorm:"column:px128x128" json:"px128x128"`
	Px256x256 *string `gorm:"column:px256x256" json:"px256x256"`
	Px64x64   *string `gorm:"column:px64x64" json:"px64x64"`
}

// TableName CorporationIcon's table name
func (*CorporationIcon) TableName() string {
	return TableNameCorporationIcon
}
