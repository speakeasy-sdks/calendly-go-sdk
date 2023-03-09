package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetRoutingFormsSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetRoutingFormsQueryParams struct {
	Count        *float64 `queryParam:"style=form,explode=true,name=count"`
	Organization string   `queryParam:"style=form,explode=true,name=organization"`
	PageToken    *string  `queryParam:"style=form,explode=true,name=page_token"`
	Sort         *string  `queryParam:"style=form,explode=true,name=sort"`
}

type GetRoutingFormsRequest struct {
	QueryParams GetRoutingFormsQueryParams
	Security    GetRoutingFormsSecurity
}

type GetRoutingFormsErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetRoutingFormsErrorResponse
// Error Object
type GetRoutingFormsErrorResponse struct {
	Details []GetRoutingFormsErrorResponseDetails `json:"details,omitempty"`
	Message string                                `json:"message"`
	Title   string                                `json:"title"`
}

type GetRoutingForms200ApplicationJSON struct {
	Collection []shared.RoutingForm `json:"collection"`
	Pagination shared.Pagination    `json:"pagination"`
}

type GetRoutingFormsResponse struct {
	ContentType                             string
	ErrorResponse                           *GetRoutingFormsErrorResponse
	StatusCode                              int
	RawResponse                             *http.Response
	GetRoutingForms200ApplicationJSONObject *GetRoutingForms200ApplicationJSON
}
