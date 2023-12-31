// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameKillmailItem = "killmail_item"

// KillmailItem mapped from table <killmail_item>
type KillmailItem struct {
	ID                int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	KillmailID        int64  `gorm:"column:killmail_id;not null" json:"killmail_id"`
	TypeID            int32  `gorm:"column:type_id;not null" json:"type_id"`
	Flag              int32  `gorm:"column:flag;not null" json:"flag"`
	QuantityDestroyed *int32 `gorm:"column:quantity_destroyed" json:"quantity_destroyed"`
	QuantityDropped   *int32 `gorm:"column:quantity_dropped" json:"quantity_dropped"`
	Singleton         bool   `gorm:"column:singleton;not null" json:"singleton"`
}

// TableName KillmailItem's table name
func (*KillmailItem) TableName() string {
	return TableNameKillmailItem
}
