package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetRoutingFormSubmissionsSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetRoutingFormSubmissionsQueryParams struct {
	Count     *float64 `queryParam:"style=form,explode=true,name=count"`
	Form      string   `queryParam:"style=form,explode=true,name=form"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
	Sort      *string  `queryParam:"style=form,explode=true,name=sort"`
}

type GetRoutingFormSubmissionsRequest struct {
	QueryParams GetRoutingFormSubmissionsQueryParams
	Security    GetRoutingFormSubmissionsSecurity
}

type GetRoutingFormSubmissionsErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetRoutingFormSubmissionsErrorResponse
// Error Object
type GetRoutingFormSubmissionsErrorResponse struct {
	Details []GetRoutingFormSubmissionsErrorResponseDetails `json:"details,omitempty"`
	Message string                                          `json:"message"`
	Title   string                                          `json:"title"`
}

type GetRoutingFormSubmissions200ApplicationJSON struct {
	Collection []shared.RoutingFormSubmission `json:"collection"`
	Pagination shared.Pagination              `json:"pagination"`
}

type GetRoutingFormSubmissionsResponse struct {
	ContentType                                       string
	ErrorResponse                                     *GetRoutingFormSubmissionsErrorResponse
	StatusCode                                        int
	RawResponse                                       *http.Response
	GetRoutingFormSubmissions200ApplicationJSONObject *GetRoutingFormSubmissions200ApplicationJSON
}
