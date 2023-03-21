package shared

type GoogleConferenceStatusEnum string

const (
	GoogleConferenceStatusEnumInitiated  GoogleConferenceStatusEnum = "initiated"
	GoogleConferenceStatusEnumProcessing GoogleConferenceStatusEnum = "processing"
	GoogleConferenceStatusEnumPushed     GoogleConferenceStatusEnum = "pushed"
	GoogleConferenceStatusEnumFailed     GoogleConferenceStatusEnum = "failed"
)

type GoogleConferenceTypeEnum string

const (
	GoogleConferenceTypeEnumGoogleConference GoogleConferenceTypeEnum = "google_conference"
)

// GoogleConference
// Details about an Event that will take place using a Google Meet / Hangout conference
type GoogleConference struct {
	JoinURL string                     `json:"join_url"`
	Status  GoogleConferenceStatusEnum `json:"status"`
	Type    GoogleConferenceTypeEnum   `json:"type"`
}
