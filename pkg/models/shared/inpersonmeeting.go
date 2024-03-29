// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InPersonMeetingTypeEnum - Indicates that the event will be an in-person meeting.
type InPersonMeetingTypeEnum string

const (
	InPersonMeetingTypeEnumPhysical InPersonMeetingTypeEnum = "physical"
)

func (e *InPersonMeetingTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "physical":
		*e = InPersonMeetingTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InPersonMeetingTypeEnum: %s", s)
	}
}

// InPersonMeeting - Information about the physical (in-person) event location
type InPersonMeeting struct {
	// The physical location specified by the event host (publisher)
	Location string `json:"location"`
	// Indicates that the event will be an in-person meeting.
	Type InPersonMeetingTypeEnum `json:"type"`
}
