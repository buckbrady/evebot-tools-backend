package utils

import (
	"github.com/golang-module/carbon/v2"
	"time"
)

var (
	utc *time.Location
)

func init() {
	utc, _ = time.LoadLocation(carbon.UTC)
}

func CheckTTL(t time.Time, ttl int) bool {
	c := carbon.CreateFromStdTime(t).SetLocation(utc)
	c.AddSeconds(ttl)
	if c.IsPast() {
		return false
	}
	return true
}
