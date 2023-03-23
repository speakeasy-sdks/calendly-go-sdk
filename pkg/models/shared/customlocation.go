// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CustomLocationTypeEnum - The event location doesn't fall into a standard category defined by the event host (publisher)
type CustomLocationTypeEnum string

const (
	CustomLocationTypeEnumCustom CustomLocationTypeEnum = "custom"
)

// CustomLocation - Use this to describe an existing Calendly-supported event location.
type CustomLocation struct {
	// The location description provided by the event host (publisher)
	Location string `json:"location"`
	// The event location doesn't fall into a standard category defined by the event host (publisher)
	Type CustomLocationTypeEnum `json:"type"`
}
