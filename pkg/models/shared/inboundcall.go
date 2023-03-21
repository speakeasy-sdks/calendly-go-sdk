package shared

type InboundCallTypeEnum string

const (
	InboundCallTypeEnumInboundCall InboundCallTypeEnum = "inbound_call"
)

// InboundCall
// Invitee will call meeting publisher at the specified phone number
type InboundCall struct {
	Location string              `json:"location"`
	Type     InboundCallTypeEnum `json:"type"`
}
