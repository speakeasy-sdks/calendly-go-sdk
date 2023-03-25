// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserBusyTimeEvent struct {
	// The uri associated with the calendly scheduled event
	URI string `json:"uri"`
}

// UserBusyTimeTypeEnum - Indicates whether the scheduled event is internal or external
type UserBusyTimeTypeEnum string

const (
	UserBusyTimeTypeEnumCalendly UserBusyTimeTypeEnum = "calendly"
	UserBusyTimeTypeEnumExternal UserBusyTimeTypeEnum = "external"
)

func (e *UserBusyTimeTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "calendly":
		fallthrough
	case "external":
		*e = UserBusyTimeTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for UserBusyTimeTypeEnum: %s", s)
	}
}

// UserBusyTime - An internal or external scheduled event for a given user
type UserBusyTime struct {
	// The end time of the calendly event, as calculated by any "after" buffer set by the user
	BufferedEndTime *time.Time `json:"buffered_end_time,omitempty"`
	// The start time of the calendly event, as calculated by any "before" buffer set by the user
	BufferedStartTime *time.Time `json:"buffered_start_time,omitempty"`
	// The end time of the scheduled event in UTC time
	EndTime time.Time          `json:"end_time"`
	Event   *UserBusyTimeEvent `json:"event,omitempty"`
	// The start time of the scheduled event in UTC time
	StartTime time.Time `json:"start_time"`
	// Indicates whether the scheduled event is internal or external
	Type UserBusyTimeTypeEnum `json:"type"`
}
