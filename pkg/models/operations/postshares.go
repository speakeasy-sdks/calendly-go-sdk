package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/types"
	"net/http"
)

type PostSharesSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type PostSharesRequestBodyAvailabilityRuleRulesIntervals struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

type PostSharesRequestBodyAvailabilityRuleRulesTypeEnum string

const (
	PostSharesRequestBodyAvailabilityRuleRulesTypeEnumWday PostSharesRequestBodyAvailabilityRuleRulesTypeEnum = "wday"
	PostSharesRequestBodyAvailabilityRuleRulesTypeEnumDate PostSharesRequestBodyAvailabilityRuleRulesTypeEnum = "date"
)

type PostSharesRequestBodyAvailabilityRuleRulesWdayEnum string

const (
	PostSharesRequestBodyAvailabilityRuleRulesWdayEnumSunday    PostSharesRequestBodyAvailabilityRuleRulesWdayEnum = "sunday"
	PostSharesRequestBodyAvailabilityRuleRulesWdayEnumMonday    PostSharesRequestBodyAvailabilityRuleRulesWdayEnum = "monday"
	PostSharesRequestBodyAvailabilityRuleRulesWdayEnumTuesday   PostSharesRequestBodyAvailabilityRuleRulesWdayEnum = "tuesday"
	PostSharesRequestBodyAvailabilityRuleRulesWdayEnumWednesday PostSharesRequestBodyAvailabilityRuleRulesWdayEnum = "wednesday"
	PostSharesRequestBodyAvailabilityRuleRulesWdayEnumThursday  PostSharesRequestBodyAvailabilityRuleRulesWdayEnum = "thursday"
	PostSharesRequestBodyAvailabilityRuleRulesWdayEnumFriday    PostSharesRequestBodyAvailabilityRuleRulesWdayEnum = "friday"
	PostSharesRequestBodyAvailabilityRuleRulesWdayEnumSaturday  PostSharesRequestBodyAvailabilityRuleRulesWdayEnum = "saturday"
)

type PostSharesRequestBodyAvailabilityRuleRules struct {
	Date      *types.Date                                           `json:"date,omitempty"`
	Intervals []PostSharesRequestBodyAvailabilityRuleRulesIntervals `json:"intervals,omitempty"`
	Type      *PostSharesRequestBodyAvailabilityRuleRulesTypeEnum   `json:"type,omitempty"`
	Wday      *PostSharesRequestBodyAvailabilityRuleRulesWdayEnum   `json:"wday,omitempty"`
}

type PostSharesRequestBodyAvailabilityRule struct {
	Rules    []PostSharesRequestBodyAvailabilityRuleRules `json:"rules,omitempty"`
	Timezone *string                                      `json:"timezone,omitempty"`
}

type PostSharesRequestBodyLocationConfigurationsKindEnum string

const (
	PostSharesRequestBodyLocationConfigurationsKindEnumPhysical                 PostSharesRequestBodyLocationConfigurationsKindEnum = "physical"
	PostSharesRequestBodyLocationConfigurationsKindEnumAskInvitee               PostSharesRequestBodyLocationConfigurationsKindEnum = "ask_invitee"
	PostSharesRequestBodyLocationConfigurationsKindEnumCustom                   PostSharesRequestBodyLocationConfigurationsKindEnum = "custom"
	PostSharesRequestBodyLocationConfigurationsKindEnumOutboundCall             PostSharesRequestBodyLocationConfigurationsKindEnum = "outbound_call"
	PostSharesRequestBodyLocationConfigurationsKindEnumInboundCall              PostSharesRequestBodyLocationConfigurationsKindEnum = "inbound_call"
	PostSharesRequestBodyLocationConfigurationsKindEnumGoogleConference         PostSharesRequestBodyLocationConfigurationsKindEnum = "google_conference"
	PostSharesRequestBodyLocationConfigurationsKindEnumGotomeetingConference    PostSharesRequestBodyLocationConfigurationsKindEnum = "gotomeeting_conference"
	PostSharesRequestBodyLocationConfigurationsKindEnumMicrosoftTeamsConference PostSharesRequestBodyLocationConfigurationsKindEnum = "microsoft_teams_conference"
	PostSharesRequestBodyLocationConfigurationsKindEnumWebexConference          PostSharesRequestBodyLocationConfigurationsKindEnum = "webex_conference"
	PostSharesRequestBodyLocationConfigurationsKindEnumZoomConference           PostSharesRequestBodyLocationConfigurationsKindEnum = "zoom_conference"
)

type PostSharesRequestBodyLocationConfigurations struct {
	AdditionalInfo *string                                              `json:"additional_info,omitempty"`
	Kind           *PostSharesRequestBodyLocationConfigurationsKindEnum `json:"kind,omitempty"`
	Location       *string                                              `json:"location,omitempty"`
	PhoneNumber    *string                                              `json:"phone_number,omitempty"`
	Position       *int64                                               `json:"position,omitempty"`
}

type PostSharesRequestBodyPeriodTypeEnum string

const (
	PostSharesRequestBodyPeriodTypeEnumAvailableMoving PostSharesRequestBodyPeriodTypeEnum = "available_moving"
	PostSharesRequestBodyPeriodTypeEnumMoving          PostSharesRequestBodyPeriodTypeEnum = "moving"
	PostSharesRequestBodyPeriodTypeEnumFixed           PostSharesRequestBodyPeriodTypeEnum = "fixed"
	PostSharesRequestBodyPeriodTypeEnumUnlimited       PostSharesRequestBodyPeriodTypeEnum = "unlimited"
)

type PostSharesRequestBody struct {
	AvailabilityRule       *PostSharesRequestBodyAvailabilityRule        `json:"availability_rule,omitempty"`
	Duration               *int64                                        `json:"duration,omitempty"`
	EndDate                *types.Date                                   `json:"end_date,omitempty"`
	EventType              string                                        `json:"event_type"`
	HideLocation           *bool                                         `json:"hide_location,omitempty"`
	LocationConfigurations []PostSharesRequestBodyLocationConfigurations `json:"location_configurations,omitempty"`
	MaxBookingTime         *int64                                        `json:"max_booking_time,omitempty"`
	Name                   *string                                       `json:"name,omitempty"`
	PeriodType             *PostSharesRequestBodyPeriodTypeEnum          `json:"period_type,omitempty"`
	StartDate              *types.Date                                   `json:"start_date,omitempty"`
}

type PostSharesRequest struct {
	Request  PostSharesRequestBody `request:"mediaType=application/json"`
	Security PostSharesSecurity
}

type PostSharesErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostSharesErrorResponse
// Error Object
type PostSharesErrorResponse struct {
	Details []PostSharesErrorResponseDetails `json:"details,omitempty"`
	Message string                           `json:"message"`
	Title   string                           `json:"title"`
}

type PostShares201ApplicationJSON struct {
	Resource shared.Share `json:"resource"`
}

type PostSharesResponse struct {
	ContentType                        string
	ErrorResponse                      *PostSharesErrorResponse
	StatusCode                         int
	RawResponse                        *http.Response
	PostShares201ApplicationJSONObject *PostShares201ApplicationJSON
}
