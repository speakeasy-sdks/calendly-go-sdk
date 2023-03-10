package shared

type LegacyCalendarEventKindEnum string

const (
	LegacyCalendarEventKindEnumExchange       LegacyCalendarEventKindEnum = "exchange"
	LegacyCalendarEventKindEnumGoogle         LegacyCalendarEventKindEnum = "google"
	LegacyCalendarEventKindEnumIcloud         LegacyCalendarEventKindEnum = "icloud"
	LegacyCalendarEventKindEnumOutlook        LegacyCalendarEventKindEnum = "outlook"
	LegacyCalendarEventKindEnumOutlookDesktop LegacyCalendarEventKindEnum = "outlook_desktop"
)

// LegacyCalendarEvent
// Information about the calendar event from the calendar provider.
type LegacyCalendarEvent struct {
	ExternalID string                      `json:"external_id"`
	Kind       LegacyCalendarEventKindEnum `json:"kind"`
}
