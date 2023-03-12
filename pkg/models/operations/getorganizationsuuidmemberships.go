package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetOrganizationsUUIDMembershipsPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetOrganizationsUUIDMembershipsRequest struct {
	PathParams GetOrganizationsUUIDMembershipsPathParams
}

type GetOrganizationsUUIDMembershipsErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetOrganizationsUUIDMembershipsErrorResponse
// Error Object
type GetOrganizationsUUIDMembershipsErrorResponse struct {
	Details []GetOrganizationsUUIDMembershipsErrorResponseDetails `json:"details,omitempty"`
	Message string                                                `json:"message"`
	Title   string                                                `json:"title"`
}

type GetOrganizationsUUIDMemberships200ApplicationJSON struct {
	Resource shared.OrganizationMembership `json:"resource"`
}

type GetOrganizationsUUIDMembershipsResponse struct {
	ContentType                                             string
	ErrorResponse                                           *GetOrganizationsUUIDMembershipsErrorResponse
	ErrorResponse1                                          *shared.ErrorResponse
	StatusCode                                              int
	RawResponse                                             *http.Response
	GetOrganizationsUUIDMemberships200ApplicationJSONObject *GetOrganizationsUUIDMemberships200ApplicationJSON
}
