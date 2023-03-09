package shared

type SubmissionEventTypeResultTypeEnum string

const (
	SubmissionEventTypeResultTypeEnumEventType SubmissionEventTypeResultTypeEnum = "event_type"
)

// SubmissionEventTypeResult
// Information about the event type Routing Form Submission result.
type SubmissionEventTypeResult struct {
	Type  SubmissionEventTypeResultTypeEnum `json:"type"`
	Value string                            `json:"value"`
}
