package shared

type SubmissionExternalURLResultTypeEnum string

const (
	SubmissionExternalURLResultTypeEnumExternalURL SubmissionExternalURLResultTypeEnum = "external_url"
)

// SubmissionExternalURLResult
// Information about the external URL Routing Form Submission result.
type SubmissionExternalURLResult struct {
	Type  SubmissionExternalURLResultTypeEnum `json:"type"`
	Value string                              `json:"value"`
}
