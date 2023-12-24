// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUniverseStation = "universe_station"

// UniverseStation mapped from table <universe_station>
type UniverseStation struct {
	ID                       int32   `gorm:"column:id;primaryKey" json:"id"`
	Name                     string  `gorm:"column:name;not null" json:"name"`
	Owner                    *int32  `gorm:"column:owner" json:"owner"`
	MaxDockableShipVolume    float64 `gorm:"column:max_dockable_ship_volume;not null" json:"max_dockable_ship_volume"`
	OfficeRentalCost         float64 `gorm:"column:office_rental_cost;not null" json:"office_rental_cost"`
	PositionX                float64 `gorm:"column:position_x;not null" json:"position_x"`
	PositionY                float64 `gorm:"column:position_y;not null" json:"position_y"`
	PositionZ                float64 `gorm:"column:position_z;not null" json:"position_z"`
	RaceID                   *int32  `gorm:"column:race_id" json:"race_id"`
	ReprocessingEfficiency   float64 `gorm:"column:reprocessing_efficiency;not null" json:"reprocessing_efficiency"`
	ReprocessingStationsTake float64 `gorm:"column:reprocessing_stations_take;not null" json:"reprocessing_stations_take"`
	SystemID                 int32   `gorm:"column:system_id;not null" json:"system_id"`
	TypeID                   int32   `gorm:"column:type_id;not null" json:"type_id"`
}

// TableName UniverseStation's table name
func (*UniverseStation) TableName() string {
	return TableNameUniverseStation
}
