package shared

import (
	"time"
)

type OrganizationInvitationStatusEnum string

const (
	OrganizationInvitationStatusEnumPending  OrganizationInvitationStatusEnum = "pending"
	OrganizationInvitationStatusEnumAccepted OrganizationInvitationStatusEnum = "accepted"
	OrganizationInvitationStatusEnumDeclined OrganizationInvitationStatusEnum = "declined"
)

// OrganizationInvitation
// Organization Invitation object
type OrganizationInvitation struct {
	CreatedAt    time.Time                        `json:"created_at"`
	Email        string                           `json:"email"`
	LastSentAt   time.Time                        `json:"last_sent_at"`
	Organization string                           `json:"organization"`
	Status       OrganizationInvitationStatusEnum `json:"status"`
	UpdatedAt    time.Time                        `json:"updated_at"`
	URI          string                           `json:"uri"`
	User         *string                          `json:"user,omitempty"`
}
