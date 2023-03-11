package shared

import (
	"time"
)

// EventTypeAvailableTime
// An available meeting time slot for the given event type
type EventTypeAvailableTime struct {
	InviteesRemaining float64   `json:"invitees_remaining"`
	SchedulingURL     string    `json:"scheduling_url"`
	StartTime         time.Time `json:"start_time"`
	Status            string    `json:"status"`
}
