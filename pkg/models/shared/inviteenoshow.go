package shared

import (
	"time"
)

// InviteeNoShow
// Information about an invitees no show.
type InviteeNoShow struct {
	CreatedAt time.Time `json:"created_at"`
	Invitee   string    `json:"invitee"`
	URI       string    `json:"uri"`
}
