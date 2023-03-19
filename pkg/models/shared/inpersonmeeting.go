package shared

type InPersonMeetingTypeEnum string

const (
	InPersonMeetingTypeEnumPhysical InPersonMeetingTypeEnum = "physical"
)

// InPersonMeeting
// Information about the physical (in-person) event location
type InPersonMeeting struct {
	Location string                  `json:"location"`
	Type     InPersonMeetingTypeEnum `json:"type"`
}
