package shared

// SubmissionQuestionAndAnswer
// All Routing Form Submission questions with answers.
type SubmissionQuestionAndAnswer struct {
	Answer       *string `json:"answer,omitempty"`
	Question     string  `json:"question"`
	QuestionUUID string  `json:"question_uuid"`
}
