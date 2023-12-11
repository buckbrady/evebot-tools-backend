package models

import (
	"github.com/kamva/mgm/v3"
	"time"
)

type ServerStatus struct {
	mgm.DefaultModel `bson:",inline"`
	Players          int       `json:"players" bson:"players"`
	ServerVersion    string    `json:"serverVersion" bson:"serverVersion"`
	StartTime        time.Time `json:"startTime" bson:"startTime"`
	Vip              bool      `json:"vip" bson:"vip"`
}
