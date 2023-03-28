// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SubmissionExternalURLResultTypeEnum - Indicates that the routing form submission resulted in a redirect to an external URL.
type SubmissionExternalURLResultTypeEnum string

const (
	SubmissionExternalURLResultTypeEnumExternalURL SubmissionExternalURLResultTypeEnum = "external_url"
)

func (e *SubmissionExternalURLResultTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "external_url":
		*e = SubmissionExternalURLResultTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SubmissionExternalURLResultTypeEnum: %s", s)
	}
}

// SubmissionExternalURLResult - Information about the external URL Routing Form Submission result.
type SubmissionExternalURLResult struct {
	// Indicates that the routing form submission resulted in a redirect to an external URL.
	Type SubmissionExternalURLResultTypeEnum `json:"type"`
	// The external URL the respondent were redirected to.
	Value string `json:"value"`
}
