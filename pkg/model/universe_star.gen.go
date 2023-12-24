// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUniverseStar = "universe_star"

// UniverseStar mapped from table <universe_star>
type UniverseStar struct {
	ID            int32   `gorm:"column:id;primaryKey" json:"id"`
	Name          string  `gorm:"column:name;not null" json:"name"`
	Age           int64   `gorm:"column:age;not null" json:"age"`
	Luminosity    float64 `gorm:"column:luminosity;not null" json:"luminosity"`
	Radius        int64   `gorm:"column:radius;not null" json:"radius"`
	SolarSystemID int32   `gorm:"column:solar_system_id;not null" json:"solar_system_id"`
	SpectralClass string  `gorm:"column:spectral_class;not null" json:"spectral_class"`
	Temperature   int32   `gorm:"column:temperature;not null" json:"temperature"`
	TypeID        int32   `gorm:"column:type_id;not null" json:"type_id"`
}

// TableName UniverseStar's table name
func (*UniverseStar) TableName() string {
	return TableNameUniverseStar
}