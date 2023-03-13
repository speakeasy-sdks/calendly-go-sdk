package shared

type AvailabilityRuleIntervals struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

type AvailabilityRuleTypeEnum string

const (
	AvailabilityRuleTypeEnumWday AvailabilityRuleTypeEnum = "wday"
	AvailabilityRuleTypeEnumDate AvailabilityRuleTypeEnum = "date"
)

type AvailabilityRuleWdayEnum string

const (
	AvailabilityRuleWdayEnumSunday    AvailabilityRuleWdayEnum = "sunday"
	AvailabilityRuleWdayEnumMonday    AvailabilityRuleWdayEnum = "monday"
	AvailabilityRuleWdayEnumTuesday   AvailabilityRuleWdayEnum = "tuesday"
	AvailabilityRuleWdayEnumWednesday AvailabilityRuleWdayEnum = "wednesday"
	AvailabilityRuleWdayEnumThursday  AvailabilityRuleWdayEnum = "thursday"
	AvailabilityRuleWdayEnumFriday    AvailabilityRuleWdayEnum = "friday"
	AvailabilityRuleWdayEnumSaturday  AvailabilityRuleWdayEnum = "saturday"
)

// AvailabilityRule
// The rules for an availability schedule.
type AvailabilityRule struct {
	Date      *string                     `json:"date,omitempty"`
	Intervals []AvailabilityRuleIntervals `json:"intervals"`
	Type      AvailabilityRuleTypeEnum    `json:"type"`
	Wday      *AvailabilityRuleWdayEnum   `json:"wday,omitempty"`
}
