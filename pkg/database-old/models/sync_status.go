package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)

type SyncStatus struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string    `json:"name" bson:"name"`
	Group            string    `json:"group" bson:"group"`
	LastSync         time.Time `json:"lastSync" bson:"lastSync"`
}
