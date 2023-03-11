package shared

import (
	"time"
)

// Guest
// An individual whom an invitee has invited to be a guest attendee to an event
type Guest struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}
