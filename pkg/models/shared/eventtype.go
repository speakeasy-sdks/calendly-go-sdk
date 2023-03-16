package shared

import (
	"time"
)

type EventTypeBookingMethodEnum string

const (
	EventTypeBookingMethodEnumInstant EventTypeBookingMethodEnum = "instant"
	EventTypeBookingMethodEnumPoll    EventTypeBookingMethodEnum = "poll"
)

type EventTypeKindEnum string

const (
	EventTypeKindEnumSolo  EventTypeKindEnum = "solo"
	EventTypeKindEnumGroup EventTypeKindEnum = "group"
)

type EventTypeKindDescriptionEnum string

const (
	EventTypeKindDescriptionEnumCollective EventTypeKindDescriptionEnum = "Collective"
	EventTypeKindDescriptionEnumGroup      EventTypeKindDescriptionEnum = "Group"
	EventTypeKindDescriptionEnumOneOnOne   EventTypeKindDescriptionEnum = "One-on-One"
	EventTypeKindDescriptionEnumRoundRobin EventTypeKindDescriptionEnum = "Round Robin"
)

type EventTypePoolingTypeEnum string

const (
	EventTypePoolingTypeEnumRoundRobin EventTypePoolingTypeEnum = "round_robin"
	EventTypePoolingTypeEnumCollective EventTypePoolingTypeEnum = "collective"
)

type EventTypeTypeEnum string

const (
	EventTypeTypeEnumStandardEventType EventTypeTypeEnum = "StandardEventType"
	EventTypeTypeEnumAdhocEventType    EventTypeTypeEnum = "AdhocEventType"
)

// EventType
// A configuration for an Event
type EventType struct {
	Active           bool                         `json:"active"`
	AdminManaged     bool                         `json:"admin_managed"`
	BookingMethod    EventTypeBookingMethodEnum   `json:"booking_method"`
	Color            string                       `json:"color"`
	CreatedAt        time.Time                    `json:"created_at"`
	CustomQuestions  []EventTypeCustomQuestion    `json:"custom_questions"`
	DeletedAt        time.Time                    `json:"deleted_at"`
	DescriptionHTML  string                       `json:"description_html"`
	DescriptionPlain string                       `json:"description_plain"`
	Duration         float64                      `json:"duration"`
	InternalNote     string                       `json:"internal_note"`
	Kind             EventTypeKindEnum            `json:"kind"`
	KindDescription  EventTypeKindDescriptionEnum `json:"kind_description"`
	Name             string                       `json:"name"`
	PoolingType      EventTypePoolingTypeEnum     `json:"pooling_type"`
	Profile          Profile                      `json:"profile"`
	SchedulingURL    string                       `json:"scheduling_url"`
	Secret           bool                         `json:"secret"`
	Slug             string                       `json:"slug"`
	Type             EventTypeTypeEnum            `json:"type"`
	UpdatedAt        time.Time                    `json:"updated_at"`
	URI              string                       `json:"uri"`
}
