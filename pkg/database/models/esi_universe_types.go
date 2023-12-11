package models

import "github.com/kamva/mgm/v3"

type UniverseType struct {
	mgm.DefaultModel `bson:",inline"`
	TypeID           int32    `json:"typeID" bson:"typeID"`
	Name             string   `json:"name" bson:"name"`
	Description      string   `json:"description" bson:"description"`
	Capacity         *float32 `json:"capacity,omitempty" bson:"capacity,omitempty"`
	//DogmaAttributes  *[]struct {
	//	AttributeID int32   `json:"attributeID" bson:"attributeID"`
	//	Value       float32 `json:"value" bson:"value"`
	//} `json:"dogmaAttributes,omitempty" bson:"dogmaAttributes,omitempty"`
	//DogmaEffects *[]struct {
	//	EffectID  int32 `json:"effectID" bson:"effectID"`
	//	IsDefault bool  `json:"isDefault" bson:"isDefault"`
	//} `json:"dogmaEffects,omitempty" bson:"dogmaEffects,omitempty"`
	GraphicID      *int32   `json:"graphicID,omitempty" bson:"graphicID,omitempty"`
	GroupID        int32    `json:"groupID" bson:"groupID"`
	IconID         *int32   `json:"iconID,omitempty" bson:"iconID,omitempty"`
	MarketGroupID  *int32   `json:"marketGroupID,omitempty" bson:"marketGroupID,omitempty"`
	Mass           *float32 `json:"mass,omitempty" bson:"mass,omitempty"`
	PackagedVolume *float32 `json:"packagedVolume,omitempty" bson:"packagedVolume,omitempty"`
	PortionSize    *int32   `json:"portionSize,omitempty" bson:"portionSize,omitempty"`
	Published      bool     `json:"published" bson:"published"`
	Radius         *float32 `json:"radius,omitempty" bson:"radius,omitempty"`
	Volume         *float32 `json:"volume,omitempty" bson:"volume,omitempty"`
}
