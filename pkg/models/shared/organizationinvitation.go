// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// OrganizationInvitationStatusEnum - The status of the invitation ("pending", "accepted", or "declined")
type OrganizationInvitationStatusEnum string

const (
	OrganizationInvitationStatusEnumPending  OrganizationInvitationStatusEnum = "pending"
	OrganizationInvitationStatusEnumAccepted OrganizationInvitationStatusEnum = "accepted"
	OrganizationInvitationStatusEnumDeclined OrganizationInvitationStatusEnum = "declined"
)

func (e *OrganizationInvitationStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "pending":
		fallthrough
	case "accepted":
		fallthrough
	case "declined":
		*e = OrganizationInvitationStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for OrganizationInvitationStatusEnum: %s", s)
	}
}

// OrganizationInvitation - Organization Invitation object
type OrganizationInvitation struct {
	// The moment the invitation was created (e.g. “2020-01-02T03:04:05.678123Z")
	CreatedAt time.Time `json:"created_at"`
	// The email address of the person who was invited to join the organization
	Email string `json:"email"`
	// The moment the invitation was last sent (e.g. "2020-01-02T03:04:05.678123Z")
	LastSentAt time.Time `json:"last_sent_at"`
	// Canonical reference (unique identifier) for the organization
	Organization string `json:"organization"`
	// The status of the invitation ("pending", "accepted", or "declined")
	Status OrganizationInvitationStatusEnum `json:"status"`
	// The moment the invitation was last updated (e.g. "2020-01-02T03:04:05.678123Z")
	UpdatedAt time.Time `json:"updated_at"`
	// Canonical reference (unique identifier) for the organization invitation
	URI string `json:"uri"`
	// When the invitation is accepted, a reference to the user who accepted the invitation
	User *string `json:"user,omitempty"`
}
