// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/types"
	"net/http"
)

type PostSharesRequestBodyAvailabilityRuleRulesIntervals struct {
	// Format: `"hh:mm"`
	From *string `json:"from,omitempty"`
	// Format: `"hh:mm"`
	To *string `json:"to,omitempty"`
}

type PostSharesRequestBodyAvailabilityRuleRulesTypeEnum string

const (
	PostSharesRequestBodyAvailabilityRuleRulesTypeEnumWday PostSharesRequestBodyAvailabilityRuleRulesTypeEnum = "wday"
	PostSharesRequestBodyAvailabilityRuleRulesTypeEnumDate PostSharesRequestBodyAvailabilityRuleRulesTypeEnum = "date"
)

func (e *PostSharesRequestBodyAvailabilityRuleRulesTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "wday":
		fallthrough
	case "date":
		*e = PostSharesRequestBodyAvailabilityRuleRulesTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSharesRequestBodyAvailabilityRuleRulesTypeEnum: %s", s)
	}
}

// PostSharesRequestBodyAvailabilityRuleRulesWdayEnum - is required when `type` is 'wday'
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

func (e *PostSharesRequestBodyAvailabilityRuleRulesWdayEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "sunday":
		fallthrough
	case "monday":
		fallthrough
	case "tuesday":
		fallthrough
	case "wednesday":
		fallthrough
	case "thursday":
		fallthrough
	case "friday":
		fallthrough
	case "saturday":
		*e = PostSharesRequestBodyAvailabilityRuleRulesWdayEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSharesRequestBodyAvailabilityRuleRulesWdayEnum: %s", s)
	}
}

type PostSharesRequestBodyAvailabilityRuleRules struct {
	// is required when `type` is 'date'
	// Format: `YYYY-MM-DD`
	Date      *types.Date                                           `json:"date,omitempty"`
	Intervals []PostSharesRequestBodyAvailabilityRuleRulesIntervals `json:"intervals,omitempty"`
	Type      *PostSharesRequestBodyAvailabilityRuleRulesTypeEnum   `json:"type,omitempty"`
	// is required when `type` is 'wday'
	Wday *PostSharesRequestBodyAvailabilityRuleRulesWdayEnum `json:"wday,omitempty"`
}

type PostSharesRequestBodyAvailabilityRule struct {
	// are required when an availability rule is provided
	Rules []PostSharesRequestBodyAvailabilityRuleRules `json:"rules,omitempty"`
	// is required when an availability rule is provided
	Timezone *string `json:"timezone,omitempty"`
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

func (e *PostSharesRequestBodyLocationConfigurationsKindEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "physical":
		fallthrough
	case "ask_invitee":
		fallthrough
	case "custom":
		fallthrough
	case "outbound_call":
		fallthrough
	case "inbound_call":
		fallthrough
	case "google_conference":
		fallthrough
	case "gotomeeting_conference":
		fallthrough
	case "microsoft_teams_conference":
		fallthrough
	case "webex_conference":
		fallthrough
	case "zoom_conference":
		*e = PostSharesRequestBodyLocationConfigurationsKindEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSharesRequestBodyLocationConfigurationsKindEnum: %s", s)
	}
}

type PostSharesRequestBodyLocationConfigurations struct {
	// is only supported when `kind` is 'physical' or 'inbound_call'
	AdditionalInfo *string                                              `json:"additional_info,omitempty"`
	Kind           *PostSharesRequestBodyLocationConfigurationsKindEnum `json:"kind,omitempty"`
	// is only supported when `kind` is 'physical', 'custom' or 'ask_invitee'
	Location *string `json:"location,omitempty"`
	// is required when `kind` is 'inbound_call'
	PhoneNumber *string `json:"phone_number,omitempty"`
	Position    *int64  `json:"position,omitempty"`
}

type PostSharesRequestBodyPeriodTypeEnum string

const (
	PostSharesRequestBodyPeriodTypeEnumAvailableMoving PostSharesRequestBodyPeriodTypeEnum = "available_moving"
	PostSharesRequestBodyPeriodTypeEnumMoving          PostSharesRequestBodyPeriodTypeEnum = "moving"
	PostSharesRequestBodyPeriodTypeEnumFixed           PostSharesRequestBodyPeriodTypeEnum = "fixed"
	PostSharesRequestBodyPeriodTypeEnumUnlimited       PostSharesRequestBodyPeriodTypeEnum = "unlimited"
)

func (e *PostSharesRequestBodyPeriodTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "available_moving":
		fallthrough
	case "moving":
		fallthrough
	case "fixed":
		fallthrough
	case "unlimited":
		*e = PostSharesRequestBodyPeriodTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PostSharesRequestBodyPeriodTypeEnum: %s", s)
	}
}

type PostSharesRequestBody struct {
	AvailabilityRule *PostSharesRequestBodyAvailabilityRule `json:"availability_rule,omitempty"`
	Duration         *int64                                 `json:"duration,omitempty"`
	// is required when `period_type` is 'fixed'
	// Format: `YYYY-MM-DD`
	EndDate   *types.Date `json:"end_date,omitempty"`
	EventType string      `json:"event_type"`
	// determines if a location is hidden until invitee books a spot, only respected when there is a single custom location configured
	HideLocation           *bool                                         `json:"hide_location,omitempty"`
	LocationConfigurations []PostSharesRequestBodyLocationConfigurations `json:"location_configurations,omitempty"`
	// is required when `period_type` is 'moving' or 'available_moving'
	MaxBookingTime *int64                               `json:"max_booking_time,omitempty"`
	Name           *string                              `json:"name,omitempty"`
	PeriodType     *PostSharesRequestBodyPeriodTypeEnum `json:"period_type,omitempty"`
	// is required when `period_type` is 'fixed'
	// Format: `YYYY-MM-DD`
	StartDate *types.Date `json:"start_date,omitempty"`
}

type PostSharesErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostSharesErrorResponse - Error Object
type PostSharesErrorResponse struct {
	Details []PostSharesErrorResponseDetails `json:"details,omitempty"`
	Message string                           `json:"message"`
	Title   string                           `json:"title"`
}

// PostShares201ApplicationJSON - Created
type PostShares201ApplicationJSON struct {
	Resource shared.Share `json:"resource"`
}

type PostSharesResponse struct {
	ContentType string
	// Request is not valid
	ErrorResponse *PostSharesErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// Created
	PostShares201ApplicationJSONObject *PostShares201ApplicationJSON
}
