// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// InviteeSpecifiedLocationTypeEnum - The event location selected by the invitee
type InviteeSpecifiedLocationTypeEnum string

const (
	InviteeSpecifiedLocationTypeEnumAskInvitee InviteeSpecifiedLocationTypeEnum = "ask_invitee"
)

// InviteeSpecifiedLocation - Information about an event location that’s specified by the invitee.
type InviteeSpecifiedLocation struct {
	// The event location description provided by the invitee
	Location string `json:"location"`
	// The event location selected by the invitee
	Type InviteeSpecifiedLocationTypeEnum `json:"type"`
}
