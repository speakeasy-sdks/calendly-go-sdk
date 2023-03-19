package shared

type SubmissionCustomMessageResultTypeEnum string

const (
	SubmissionCustomMessageResultTypeEnumCustomMessage SubmissionCustomMessageResultTypeEnum = "custom_message"
)

// SubmissionCustomMessageResultValue
// Contains an object with custom message headline and body.
type SubmissionCustomMessageResultValue struct {
	Body     string `json:"body"`
	Headline string `json:"headline"`
}

// SubmissionCustomMessageResult
// Information about the custom message Routing Form Submission result.
type SubmissionCustomMessageResult struct {
	Type  SubmissionCustomMessageResultTypeEnum `json:"type"`
	Value SubmissionCustomMessageResultValue    `json:"value"`
}
