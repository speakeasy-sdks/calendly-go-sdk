package shared

type OutboundCallTypeEnum string

const (
	OutboundCallTypeEnumOutboundCall OutboundCallTypeEnum = "outbound_call"
)

// OutboundCall
// Meeting publisher will call the Invitee
type OutboundCall struct {
	Location string               `json:"location"`
	Type     OutboundCallTypeEnum `json:"type"`
}
