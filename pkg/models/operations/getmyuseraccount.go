package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetMyUserAccountSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetMyUserAccountRequest struct {
	Security GetMyUserAccountSecurity
}

type GetMyUserAccountErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetMyUserAccountErrorResponse
// Error Object
type GetMyUserAccountErrorResponse struct {
	Details []GetMyUserAccountErrorResponseDetails `json:"details,omitempty"`
	Message string                                 `json:"message"`
	Title   string                                 `json:"title"`
}

// GetMyUserAccount200ApplicationJSON
// Service response
type GetMyUserAccount200ApplicationJSON struct {
	Resource shared.User `json:"resource"`
}

type GetMyUserAccountResponse struct {
	ContentType                              string
	ErrorResponse                            *GetMyUserAccountErrorResponse
	StatusCode                               int
	RawResponse                              *http.Response
	GetMyUserAccount200ApplicationJSONObject *GetMyUserAccount200ApplicationJSON
}
