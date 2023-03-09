package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type PostInviteeNoShowSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type PostInviteeNoShowRequestBody struct {
	Invitee *string `json:"invitee,omitempty"`
}

type PostInviteeNoShowRequest struct {
	Request  PostInviteeNoShowRequestBody `request:"mediaType=application/json"`
	Security PostInviteeNoShowSecurity
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
