package shared

type CancellationCancelerTypeEnum string

const (
	CancellationCancelerTypeEnumHost    CancellationCancelerTypeEnum = "host"
	CancellationCancelerTypeEnumInvitee CancellationCancelerTypeEnum = "invitee"
)

// Cancellation
// Provides data pertaining to the cancellation of the Event
type Cancellation struct {
	CanceledBy   string                       `json:"canceled_by"`
	CancelerType CancellationCancelerTypeEnum `json:"canceler_type"`
	Reason       string                       `json:"reason"`
}
