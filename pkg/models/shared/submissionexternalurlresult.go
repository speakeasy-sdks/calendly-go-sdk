// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SubmissionExternalURLResultTypeEnum - Indicates that the routing form submission resulted in a redirect to an external URL.
type SubmissionExternalURLResultTypeEnum string

const (
	SubmissionExternalURLResultTypeEnumExternalURL SubmissionExternalURLResultTypeEnum = "external_url"
)

// SubmissionExternalURLResult - Information about the external URL Routing Form Submission result.
type SubmissionExternalURLResult struct {
	// Indicates that the routing form submission resulted in a redirect to an external URL.
	Type SubmissionExternalURLResultTypeEnum `json:"type"`
	// The external URL the respondent were redirected to.
	Value string `json:"value"`
}
