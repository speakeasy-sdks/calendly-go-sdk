// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CancellationCancelerTypeEnum string

const (
	CancellationCancelerTypeEnumHost    CancellationCancelerTypeEnum = "host"
	CancellationCancelerTypeEnumInvitee CancellationCancelerTypeEnum = "invitee"
)

func (e *CancellationCancelerTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "host":
		fallthrough
	case "invitee":
		*e = CancellationCancelerTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CancellationCancelerTypeEnum: %s", s)
	}
}

// Cancellation - Provides data pertaining to the cancellation of the Event
type Cancellation struct {
	// Name of the person whom canceled
	CanceledBy   string                       `json:"canceled_by"`
	CancelerType CancellationCancelerTypeEnum `json:"canceler_type"`
	// Reason that the cancellation occurred
	Reason string `json:"reason"`
}
