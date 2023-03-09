package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetEventTypesUUIDSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetEventTypesUUIDPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetEventTypesUUIDRequest struct {
	PathParams GetEventTypesUUIDPathParams
	Security   GetEventTypesUUIDSecurity
}

type GetEventTypesUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetEventTypesUUIDErrorResponse
// Error Object
type GetEventTypesUUIDErrorResponse struct {
	Details []GetEventTypesUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                  `json:"message"`
	Title   string                                  `json:"title"`
}

type GetEventTypesUUID200ApplicationJSON struct {
	Resource shared.EventType `json:"resource"`
}

type GetEventTypesUUIDResponse struct {
	ContentType                               string
	ErrorResponse                             *GetEventTypesUUIDErrorResponse
	StatusCode                                int
	RawResponse                               *http.Response
	GetEventTypesUUID200ApplicationJSONObject *GetEventTypesUUID200ApplicationJSON
}
