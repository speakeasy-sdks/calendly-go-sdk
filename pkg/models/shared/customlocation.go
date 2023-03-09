package shared

type CustomLocationTypeEnum string

const (
	CustomLocationTypeEnumCustom CustomLocationTypeEnum = "custom"
)

// CustomLocation
// Use this to describe an existing Calendly-supported event location.
type CustomLocation struct {
	Location string                 `json:"location"`
	Type     CustomLocationTypeEnum `json:"type"`
}
