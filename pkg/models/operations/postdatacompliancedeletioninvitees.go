package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type PostDataComplianceDeletionInviteesRequestBody struct {
	Emails []string `json:"emails"`
}

type PostDataComplianceDeletionInviteesRequest struct {
	Request PostDataComplianceDeletionInviteesRequestBody `request:"mediaType=application/json"`
}

type PostDataComplianceDeletionInviteesErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostDataComplianceDeletionInviteesErrorResponse
// Error Object
type PostDataComplianceDeletionInviteesErrorResponse struct {
	Details []PostDataComplianceDeletionInviteesErrorResponseDetails `json:"details,omitempty"`
	Message string                                                   `json:"message"`
	Title   string                                                   `json:"title"`
}

type PostDataComplianceDeletionInviteesResponse struct {
	ContentType                                                string
	ErrorResponse                                              *PostDataComplianceDeletionInviteesErrorResponse
	ErrorResponse1                                             *shared.ErrorResponse
	StatusCode                                                 int
	RawResponse                                                *http.Response
	PostDataComplianceDeletionInvitees202ApplicationJSONObject map[string]interface{}
}
