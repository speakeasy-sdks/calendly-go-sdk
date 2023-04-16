// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// RoutingFormSubmissionSubmitterTypeEnum - Type of the respondent resource that submitted the form and scheduled a meeting.
type RoutingFormSubmissionSubmitterTypeEnum string

const (
	RoutingFormSubmissionSubmitterTypeEnumInvitee RoutingFormSubmissionSubmitterTypeEnum = "Invitee"
	RoutingFormSubmissionSubmitterTypeEnumNull    RoutingFormSubmissionSubmitterTypeEnum = "null"
)

func (e *RoutingFormSubmissionSubmitterTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Invitee":
		fallthrough
	case "null":
		*e = RoutingFormSubmissionSubmitterTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for RoutingFormSubmissionSubmitterTypeEnum: %s", s)
	}
}

// RoutingFormSubmission - Information about a Routing Form Submission.
type RoutingFormSubmission struct {
	// The moment the routing form was submitted.
	CreatedAt time.Time `json:"created_at"`
	// All Routing Form Submission questions with answers.
	QuestionsAndAnswers []SubmissionQuestionAndAnswer `json:"questions_and_answers"`
	// The polymorphic base type for a Routing Form Submission result.
	Result *SubmissionResult `json:"result,omitempty"`
	// The URI of the routing form that's associated with the submission.
	RoutingForm string `json:"routing_form"`
	// The reference to the Invitee resource when routing form submission results in a scheduled meeting.
	Submitter string `json:"submitter"`
	// Type of the respondent resource that submitted the form and scheduled a meeting.
	SubmitterType RoutingFormSubmissionSubmitterTypeEnum `json:"submitter_type"`
	// The UTM and Salesforce tracking parameters associated with a Routing Form Submission.
	Tracking SubmissionTracking `json:"tracking"`
	// The moment when the routing form submission was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// Canonical reference (unique identifier) for the routing form submission.
	URI string `json:"uri"`
}
