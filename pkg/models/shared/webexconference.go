package shared

type WebExConferenceDataTelephonyCallInNumbers struct {
	CallInNumber string `json:"callInNumber"`
	Label        string `json:"label"`
	TollType     string `json:"tollType"`
}

type WebExConferenceDataTelephony struct {
	CallInNumbers []WebExConferenceDataTelephonyCallInNumbers `json:"callInNumbers"`
}

// WebExConferenceData
// The conference metadata supplied by GoToMeeting
type WebExConferenceData struct {
	ID        string                       `json:"id"`
	Password  string                       `json:"password"`
	Telephony WebExConferenceDataTelephony `json:"telephony"`
}

type WebExConferenceStatusEnum string

const (
	WebExConferenceStatusEnumInitiated  WebExConferenceStatusEnum = "initiated"
	WebExConferenceStatusEnumProcessing WebExConferenceStatusEnum = "processing"
	WebExConferenceStatusEnumPushed     WebExConferenceStatusEnum = "pushed"
	WebExConferenceStatusEnumFailed     WebExConferenceStatusEnum = "failed"
)

type WebExConferenceTypeEnum string

const (
	WebExConferenceTypeEnumWebexConference WebExConferenceTypeEnum = "webex_conference"
)

// WebExConference
// Details about an Event that will take place using a WebEx conference
type WebExConference struct {
	Data    WebExConferenceData       `json:"data"`
	JoinURL string                    `json:"join_url"`
	Status  WebExConferenceStatusEnum `json:"status"`
	Type    WebExConferenceTypeEnum   `json:"type"`
}
