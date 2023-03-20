package shared

type MicrosoftTeamsConferenceDataAudioConferencing struct {
	ConferenceID *string `json:"conferenceId,omitempty"`
	DialinURL    *string `json:"dialinUrl,omitempty"`
	TollNumber   *string `json:"tollNumber,omitempty"`
}

// MicrosoftTeamsConferenceData
// The conference metadata supplied by Microsoft Teams
type MicrosoftTeamsConferenceData struct {
	AudioConferencing *MicrosoftTeamsConferenceDataAudioConferencing `json:"audioConferencing,omitempty"`
	ID                *string                                        `json:"id,omitempty"`
}

type MicrosoftTeamsConferenceStatusEnum string

const (
	MicrosoftTeamsConferenceStatusEnumInitiated  MicrosoftTeamsConferenceStatusEnum = "initiated"
	MicrosoftTeamsConferenceStatusEnumProcessing MicrosoftTeamsConferenceStatusEnum = "processing"
	MicrosoftTeamsConferenceStatusEnumPushed     MicrosoftTeamsConferenceStatusEnum = "pushed"
	MicrosoftTeamsConferenceStatusEnumFailed     MicrosoftTeamsConferenceStatusEnum = "failed"
)

type MicrosoftTeamsConferenceTypeEnum string

const (
	MicrosoftTeamsConferenceTypeEnumMicrosoftTeamsConference MicrosoftTeamsConferenceTypeEnum = "microsoft_teams_conference"
)

// MicrosoftTeamsConference
// Meeting will take place in a Microsoft Teams conference
type MicrosoftTeamsConference struct {
	Data    MicrosoftTeamsConferenceData       `json:"data"`
	JoinURL string                             `json:"join_url"`
	Status  MicrosoftTeamsConferenceStatusEnum `json:"status"`
	Type    MicrosoftTeamsConferenceTypeEnum   `json:"type"`
}
