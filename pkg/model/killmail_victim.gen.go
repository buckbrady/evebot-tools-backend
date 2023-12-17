// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameKillmailVictim = "killmail_victim"

// KillmailVictim mapped from table <killmail_victim>
type KillmailVictim struct {
	ID            int64  `gorm:"column:id;primaryKey" json:"id"`
	CharacterID   int32  `gorm:"column:character_id;not null" json:"character_id"`
	CorporationID int32  `gorm:"column:corporation_id;not null" json:"corporation_id"`
	AllianceID    *int32 `gorm:"column:alliance_id" json:"alliance_id"`
	FactionID     *int32 `gorm:"column:faction_id" json:"faction_id"`
}

// TableName KillmailVictim's table name
func (*KillmailVictim) TableName() string {
	return TableNameKillmailVictim
}
