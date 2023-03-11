package shared

import (
	"time"
)

type EventEventMemberships struct {
	User string `json:"user"`
}

type EventInviteesCounter struct {
	Active float64 `json:"active"`
	Limit  float64 `json:"limit"`
	Total  float64 `json:"total"`
}

type EventStatusEnum string

const (
	EventStatusEnumActive   EventStatusEnum = "active"
	EventStatusEnumCanceled EventStatusEnum = "canceled"
)

// Event
// Information about a scheduled meeting
type Event struct {
	CalendarEvent    LegacyCalendarEvent     `json:"calendar_event"`
	Cancellation     *Cancellation           `json:"cancellation,omitempty"`
	CreatedAt        time.Time               `json:"created_at"`
	EndTime          time.Time               `json:"end_time"`
	EventGuests      []Guest                 `json:"event_guests"`
	EventMemberships []EventEventMemberships `json:"event_memberships"`
	EventType        string                  `json:"event_type"`
	InviteesCounter  EventInviteesCounter    `json:"invitees_counter"`
	Location         interface{}             `json:"location"`
	Name             string                  `json:"name"`
	StartTime        time.Time               `json:"start_time"`
	Status           EventStatusEnum         `json:"status"`
	UpdatedAt        time.Time               `json:"updated_at"`
	URI              string                  `json:"uri"`
}
