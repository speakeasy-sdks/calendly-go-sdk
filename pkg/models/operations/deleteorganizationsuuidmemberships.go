// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteOrganizationsUUIDMembershipsRequest struct {
	// The organization membership's unique identifier
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type DeleteOrganizationsUUIDMembershipsErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// DeleteOrganizationsUUIDMembershipsErrorResponse - Error Object
type DeleteOrganizationsUUIDMembershipsErrorResponse struct {
	Details []DeleteOrganizationsUUIDMembershipsErrorResponseDetails `json:"details,omitempty"`
	Message string                                                   `json:"message"`
	Title   string                                                   `json:"title"`
}

type DeleteOrganizationsUUIDMembershipsResponse struct {
	ContentType string
	// Request is not valid
	ErrorResponse *DeleteOrganizationsUUIDMembershipsErrorResponse
	// Caller not authorized to perform this action
	ErrorResponse1 *shared.ErrorResponse
	StatusCode     int
	RawResponse    *http.Response
}
