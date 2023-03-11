package shared

// GotoMeetingConferenceData
// The conference metadata supplied by GoToMeeting
type GotoMeetingConferenceData struct {
	ConferenceCallInfo *string  `json:"conferenceCallInfo,omitempty"`
	UniqueMeetingID    *float64 `json:"uniqueMeetingId,omitempty"`
}

type GotoMeetingConferenceStatusEnum string

const (
	GotoMeetingConferenceStatusEnumInitiated  GotoMeetingConferenceStatusEnum = "initiated"
	GotoMeetingConferenceStatusEnumProcessing GotoMeetingConferenceStatusEnum = "processing"
	GotoMeetingConferenceStatusEnumPushed     GotoMeetingConferenceStatusEnum = "pushed"
	GotoMeetingConferenceStatusEnumFailed     GotoMeetingConferenceStatusEnum = "failed"
)

type GotoMeetingConferenceTypeEnum string

const (
	GotoMeetingConferenceTypeEnumGotomeeting GotoMeetingConferenceTypeEnum = "gotomeeting"
)

// GotoMeetingConference
// Details about an Event that will take place using a GotoMeeting conference
type GotoMeetingConference struct {
	Data    GotoMeetingConferenceData       `json:"data"`
	JoinURL string                          `json:"join_url"`
	Status  GotoMeetingConferenceStatusEnum `json:"status"`
	Type    GotoMeetingConferenceTypeEnum   `json:"type"`
}
