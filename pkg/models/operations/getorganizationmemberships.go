package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetOrganizationMembershipsQueryParams struct {
	Count        *float64 `queryParam:"style=form,explode=true,name=count"`
	Email        *string  `queryParam:"style=form,explode=true,name=email"`
	Organization *string  `queryParam:"style=form,explode=true,name=organization"`
	PageToken    *string  `queryParam:"style=form,explode=true,name=page_token"`
	User         *string  `queryParam:"style=form,explode=true,name=user"`
}

type GetOrganizationMembershipsRequest struct {
	QueryParams GetOrganizationMembershipsQueryParams
}

type GetOrganizationMembershipsErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetOrganizationMembershipsErrorResponse
// Error Object
type GetOrganizationMembershipsErrorResponse struct {
	Details []GetOrganizationMembershipsErrorResponseDetails `json:"details,omitempty"`
	Message string                                           `json:"message"`
	Title   string                                           `json:"title"`
}

// GetOrganizationMemberships200ApplicationJSON
// Service response
type GetOrganizationMemberships200ApplicationJSON struct {
	Collection []shared.OrganizationMembership `json:"collection"`
	Pagination shared.Pagination               `json:"pagination"`
}

type GetOrganizationMembershipsResponse struct {
	ContentType                                        string
	ErrorResponse                                      *GetOrganizationMembershipsErrorResponse
	ErrorResponse1                                     *shared.ErrorResponse
	StatusCode                                         int
	RawResponse                                        *http.Response
	GetOrganizationMemberships200ApplicationJSONObject *GetOrganizationMemberships200ApplicationJSON
}
