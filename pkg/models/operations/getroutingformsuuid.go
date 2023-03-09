package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetRoutingFormsUUIDSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetRoutingFormsUUIDPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetRoutingFormsUUIDRequest struct {
	PathParams GetRoutingFormsUUIDPathParams
	Security   GetRoutingFormsUUIDSecurity
}

type GetRoutingFormsUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetRoutingFormsUUIDErrorResponse
// Error Object
type GetRoutingFormsUUIDErrorResponse struct {
	Details []GetRoutingFormsUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                    `json:"message"`
	Title   string                                    `json:"title"`
}

type GetRoutingFormsUUID200ApplicationJSON struct {
	Resource shared.RoutingForm `json:"resource"`
}

type GetRoutingFormsUUIDResponse struct {
	ContentType                                 string
	ErrorResponse                               *GetRoutingFormsUUIDErrorResponse
	StatusCode                                  int
	RawResponse                                 *http.Response
	GetRoutingFormsUUID200ApplicationJSONObject *GetRoutingFormsUUID200ApplicationJSON
}
