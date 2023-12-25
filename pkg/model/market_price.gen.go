// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMarketPrice = "market_price"

// MarketPrice mapped from table <market_price>
type MarketPrice struct {
	TypeID        int32     `gorm:"column:type_id;primaryKey" json:"type_id"`
	AveragePrice  *float64  `gorm:"column:average_price" json:"average_price"`
	AdjustedPrice *float64  `gorm:"column:adjusted_price" json:"adjusted_price"`
	Timestamp     time.Time `gorm:"column:timestamp;primaryKey;default:now()" json:"timestamp"`
}

// TableName MarketPrice's table name
func (*MarketPrice) TableName() string {
	return TableNameMarketPrice
}
