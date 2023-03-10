package shared

type InviteeSpecifiedLocationTypeEnum string

const (
	InviteeSpecifiedLocationTypeEnumAskInvitee InviteeSpecifiedLocationTypeEnum = "ask_invitee"
)

// InviteeSpecifiedLocation
// Information about an event location that’s specified by the invitee.
type InviteeSpecifiedLocation struct {
	Location string                           `json:"location"`
	Type     InviteeSpecifiedLocationTypeEnum `json:"type"`
}
