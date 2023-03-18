package shared

import (
	"time"
)

type UserBusyTimeEvent struct {
	URI string `json:"uri"`
}

type UserBusyTimeTypeEnum string

const (
	UserBusyTimeTypeEnumCalendly UserBusyTimeTypeEnum = "calendly"
	UserBusyTimeTypeEnumExternal UserBusyTimeTypeEnum = "external"
)

// UserBusyTime
// An internal or external scheduled event for a given user
type UserBusyTime struct {
	BufferedEndTime   *time.Time           `json:"buffered_end_time,omitempty"`
	BufferedStartTime *time.Time           `json:"buffered_start_time,omitempty"`
	EndTime           time.Time            `json:"end_time"`
	Event             *UserBusyTimeEvent   `json:"event,omitempty"`
	StartTime         time.Time            `json:"start_time"`
	Type              UserBusyTimeTypeEnum `json:"type"`
}
