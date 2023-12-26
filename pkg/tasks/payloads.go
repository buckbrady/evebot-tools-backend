package tasks

import "time"

type CronJobPayload struct {
	Timestamp time.Time
	TTL       int
}

type CronJobPayloadWithType struct {
	Timestamp time.Time
	TTL       int
	TypeID    int32
}

type CronJobPayloadWithType64 struct {
	Timestamp time.Time
	TTL       int
	TypeID    int64
}

type CronJobPayloadTypeWithRegion struct {
	Timestamp time.Time
	TTL       int
	TypeID    int32
	RegionID  int32
}
