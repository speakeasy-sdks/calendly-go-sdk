package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type PostInviteeNoShowRequestBody struct {
	Invitee *string `json:"invitee,omitempty"`
}

type PostInviteeNoShowRequest struct {
	Request PostInviteeNoShowRequestBody `request:"mediaType=application/json"`
}

type PostInviteeNoShowErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostInviteeNoShowErrorResponse
// Error Object
type PostInviteeNoShowErrorResponse struct {
	Details []PostInviteeNoShowErrorResponseDetails `json:"details,omitempty"`
	Message string                                  `json:"message"`
	Title   string                                  `json:"title"`
}

type PostInviteeNoShow201ApplicationJSON struct {
	Resource shared.InviteeNoShow `json:"resource"`
}

type PostInviteeNoShowResponse struct {
	ContentType                               string
	ErrorResponse                             *PostInviteeNoShowErrorResponse
	StatusCode                                int
	RawResponse                               *http.Response
	PostInviteeNoShow201ApplicationJSONObject *PostInviteeNoShow201ApplicationJSON
}
