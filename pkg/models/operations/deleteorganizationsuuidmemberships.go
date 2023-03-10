package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteOrganizationsUUIDMembershipsPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type DeleteOrganizationsUUIDMembershipsRequest struct {
	PathParams DeleteOrganizationsUUIDMembershipsPathParams
}

type DeleteOrganizationsUUIDMembershipsErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// DeleteOrganizationsUUIDMembershipsErrorResponse
// Error Object
type DeleteOrganizationsUUIDMembershipsErrorResponse struct {
	Details []DeleteOrganizationsUUIDMembershipsErrorResponseDetails `json:"details,omitempty"`
	Message string                                                   `json:"message"`
	Title   string                                                   `json:"title"`
}

type DeleteOrganizationsUUIDMembershipsResponse struct {
	ContentType    string
	ErrorResponse  *DeleteOrganizationsUUIDMembershipsErrorResponse
	ErrorResponse1 *shared.ErrorResponse
	StatusCode     int
	RawResponse    *http.Response
}
