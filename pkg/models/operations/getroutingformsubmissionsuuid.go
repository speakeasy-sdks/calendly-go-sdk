package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetRoutingFormSubmissionsUUIDSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetRoutingFormSubmissionsUUIDPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetRoutingFormSubmissionsUUIDRequest struct {
	PathParams GetRoutingFormSubmissionsUUIDPathParams
	Security   GetRoutingFormSubmissionsUUIDSecurity
}

type GetRoutingFormSubmissionsUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetRoutingFormSubmissionsUUIDErrorResponse
// Error Object
type GetRoutingFormSubmissionsUUIDErrorResponse struct {
	Details []GetRoutingFormSubmissionsUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                              `json:"message"`
	Title   string                                              `json:"title"`
}

type GetRoutingFormSubmissionsUUID200ApplicationJSON struct {
	Resource shared.RoutingFormSubmission `json:"resource"`
}

type GetRoutingFormSubmissionsUUIDResponse struct {
	ContentType                                           string
	ErrorResponse                                         *GetRoutingFormSubmissionsUUIDErrorResponse
	StatusCode                                            int
	RawResponse                                           *http.Response
	GetRoutingFormSubmissionsUUID200ApplicationJSONObject *GetRoutingFormSubmissionsUUID200ApplicationJSON
}
