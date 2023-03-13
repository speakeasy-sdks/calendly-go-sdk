package shared

type ZoomConferenceDataExtra struct {
	IntlNumbersURL *string `json:"intl_numbers_url,omitempty"`
}

type ZoomConferenceDataSettingsGlobalDialInNumbers struct {
	City        *string `json:"city,omitempty"`
	Country     *string `json:"country,omitempty"`
	CountryName *string `json:"country_name,omitempty"`
	Number      *string `json:"number,omitempty"`
	Type        *string `json:"type,omitempty"`
}

type ZoomConferenceDataSettings struct {
	GlobalDialInNumbers []ZoomConferenceDataSettingsGlobalDialInNumbers `json:"global_dial_in_numbers,omitempty"`
}

// ZoomConferenceData
// The conference metadata supplied by Zoom
type ZoomConferenceData struct {
	Extra    *ZoomConferenceDataExtra    `json:"extra,omitempty"`
	ID       *string                     `json:"id,omitempty"`
	Password *string                     `json:"password,omitempty"`
	Settings *ZoomConferenceDataSettings `json:"settings,omitempty"`
}

type ZoomConferenceStatusEnum string

const (
	ZoomConferenceStatusEnumInitiated  ZoomConferenceStatusEnum = "initiated"
	ZoomConferenceStatusEnumProcessing ZoomConferenceStatusEnum = "processing"
	ZoomConferenceStatusEnumPushed     ZoomConferenceStatusEnum = "pushed"
	ZoomConferenceStatusEnumFailed     ZoomConferenceStatusEnum = "failed"
)

type ZoomConferenceTypeEnum string

const (
	ZoomConferenceTypeEnumZoomConference ZoomConferenceTypeEnum = "zoom_conference"
)

// ZoomConference
// Meeting will take place in a Zoom conference
type ZoomConference struct {
	Data    ZoomConferenceData       `json:"data"`
	JoinURL string                   `json:"join_url"`
	Status  ZoomConferenceStatusEnum `json:"status"`
	Type    ZoomConferenceTypeEnum   `json:"type"`
}
