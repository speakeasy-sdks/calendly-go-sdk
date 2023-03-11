package shared

import (
	"time"
)

type RoutingFormSubmissionSubmitterTypeEnum string

const (
	RoutingFormSubmissionSubmitterTypeEnumInvitee RoutingFormSubmissionSubmitterTypeEnum = "Invitee"
)

// RoutingFormSubmission
// Information about a Routing Form Submission.
type RoutingFormSubmission struct {
	CreatedAt           time.Time                              `json:"created_at"`
	QuestionsAndAnswers []SubmissionQuestionAndAnswer          `json:"questions_and_answers"`
	Result              interface{}                            `json:"result,omitempty"`
	RoutingForm         string                                 `json:"routing_form"`
	Submitter           string                                 `json:"submitter"`
	SubmitterType       RoutingFormSubmissionSubmitterTypeEnum `json:"submitter_type"`
	Tracking            SubmissionTracking                     `json:"tracking"`
	UpdatedAt           time.Time                              `json:"updated_at"`
	URI                 string                                 `json:"uri"`
}
