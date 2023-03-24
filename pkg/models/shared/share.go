// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/types"
)

// ShareSchedulingLinksOwnerTypeEnum - Resource type (currently, this is always EventType)
type ShareSchedulingLinksOwnerTypeEnum string

const (
	ShareSchedulingLinksOwnerTypeEnumEventType ShareSchedulingLinksOwnerTypeEnum = "EventType"
)

type ShareSchedulingLinks struct {
	// Scheduling link url
	BookingURL string `json:"booking_url"`
	// A link to the resource that owns this Scheduling Link (currently, this is always an Event Type)
	Owner *string `json:"owner,omitempty"`
	// Resource type (currently, this is always EventType)
	OwnerType *ShareSchedulingLinksOwnerTypeEnum `json:"owner_type,omitempty"`
}

type ShareShareOverrideAvailabilityRuleRulesIntervals struct {
	// Format: `"hh:mm"`
	From *string `json:"from,omitempty"`
	// Format: `"hh:mm"`
	To *string `json:"to,omitempty"`
}

type ShareShareOverrideAvailabilityRuleRulesTypeEnum string

const (
	ShareShareOverrideAvailabilityRuleRulesTypeEnumWday ShareShareOverrideAvailabilityRuleRulesTypeEnum = "wday"
	ShareShareOverrideAvailabilityRuleRulesTypeEnumDate ShareShareOverrideAvailabilityRuleRulesTypeEnum = "date"
)

type ShareShareOverrideAvailabilityRuleRulesWdayEnum string

const (
	ShareShareOverrideAvailabilityRuleRulesWdayEnumSunday    ShareShareOverrideAvailabilityRuleRulesWdayEnum = "sunday"
	ShareShareOverrideAvailabilityRuleRulesWdayEnumMonday    ShareShareOverrideAvailabilityRuleRulesWdayEnum = "monday"
	ShareShareOverrideAvailabilityRuleRulesWdayEnumTuesday   ShareShareOverrideAvailabilityRuleRulesWdayEnum = "tuesday"
	ShareShareOverrideAvailabilityRuleRulesWdayEnumWednesday ShareShareOverrideAvailabilityRuleRulesWdayEnum = "wednesday"
	ShareShareOverrideAvailabilityRuleRulesWdayEnumThursday  ShareShareOverrideAvailabilityRuleRulesWdayEnum = "thursday"
	ShareShareOverrideAvailabilityRuleRulesWdayEnumFriday    ShareShareOverrideAvailabilityRuleRulesWdayEnum = "friday"
	ShareShareOverrideAvailabilityRuleRulesWdayEnumSaturday  ShareShareOverrideAvailabilityRuleRulesWdayEnum = "saturday"
)

type ShareShareOverrideAvailabilityRuleRules struct {
	// Format: `YYYY-MM-DD`
	Date      *types.Date                                        `json:"date,omitempty"`
	Intervals []ShareShareOverrideAvailabilityRuleRulesIntervals `json:"intervals,omitempty"`
	Type      ShareShareOverrideAvailabilityRuleRulesTypeEnum    `json:"type"`
	Wday      *ShareShareOverrideAvailabilityRuleRulesWdayEnum   `json:"wday,omitempty"`
}

type ShareShareOverrideAvailabilityRule struct {
	Rules    []ShareShareOverrideAvailabilityRuleRules `json:"rules"`
	Timezone string                                    `json:"timezone"`
}

type ShareShareOverrideLocationConfigurationsKindEnum string

const (
	ShareShareOverrideLocationConfigurationsKindEnumPhysical                 ShareShareOverrideLocationConfigurationsKindEnum = "physical"
	ShareShareOverrideLocationConfigurationsKindEnumAskInvitee               ShareShareOverrideLocationConfigurationsKindEnum = "ask_invitee"
	ShareShareOverrideLocationConfigurationsKindEnumCustom                   ShareShareOverrideLocationConfigurationsKindEnum = "custom"
	ShareShareOverrideLocationConfigurationsKindEnumOutboundCall             ShareShareOverrideLocationConfigurationsKindEnum = "outbound_call"
	ShareShareOverrideLocationConfigurationsKindEnumInboundCall              ShareShareOverrideLocationConfigurationsKindEnum = "inbound_call"
	ShareShareOverrideLocationConfigurationsKindEnumGoogleConference         ShareShareOverrideLocationConfigurationsKindEnum = "google_conference"
	ShareShareOverrideLocationConfigurationsKindEnumGotomeetingConference    ShareShareOverrideLocationConfigurationsKindEnum = "gotomeeting_conference"
	ShareShareOverrideLocationConfigurationsKindEnumMicrosoftTeamsConference ShareShareOverrideLocationConfigurationsKindEnum = "microsoft_teams_conference"
	ShareShareOverrideLocationConfigurationsKindEnumZoomConference           ShareShareOverrideLocationConfigurationsKindEnum = "zoom_conference"
)

type ShareShareOverrideLocationConfigurations struct {
	AdditionalInfo *string                                           `json:"additional_info,omitempty"`
	Kind           *ShareShareOverrideLocationConfigurationsKindEnum `json:"kind,omitempty"`
	Location       *string                                           `json:"location,omitempty"`
	PhoneNumber    *string                                           `json:"phone_number,omitempty"`
	Position       *int64                                            `json:"position,omitempty"`
}

type ShareShareOverridePeriodTypeEnum string

const (
	ShareShareOverridePeriodTypeEnumAvailableMoving ShareShareOverridePeriodTypeEnum = "available_moving"
	ShareShareOverridePeriodTypeEnumMoving          ShareShareOverridePeriodTypeEnum = "moving"
	ShareShareOverridePeriodTypeEnumFixed           ShareShareOverridePeriodTypeEnum = "fixed"
	ShareShareOverridePeriodTypeEnumUnlimited       ShareShareOverridePeriodTypeEnum = "unlimited"
)

type ShareShareOverride struct {
	AvailabilityRule *ShareShareOverrideAvailabilityRule `json:"availability_rule,omitempty"`
	Duration         *int64                              `json:"duration,omitempty"`
	// Format: `YYYY-MM-DD`
	EndDate                *types.Date                                `json:"end_date,omitempty"`
	HideLocation           *bool                                      `json:"hide_location,omitempty"`
	LocationConfigurations []ShareShareOverrideLocationConfigurations `json:"location_configurations,omitempty"`
	MaxBookingTime         *int64                                     `json:"max_booking_time,omitempty"`
	Name                   *string                                    `json:"name,omitempty"`
	PeriodType             *ShareShareOverridePeriodTypeEnum          `json:"period_type,omitempty"`
	// Format: `YYYY-MM-DD`
	StartDate *types.Date `json:"start_date,omitempty"`
}

type Share struct {
	SchedulingLinks []ShareSchedulingLinks `json:"scheduling_links"`
	ShareOverride   ShareShareOverride     `json:"share_override"`
}
