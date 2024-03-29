// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetRoutingFormSubmissionsUUIDRequest struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetRoutingFormSubmissionsUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetRoutingFormSubmissionsUUIDErrorResponse - Error Object
type GetRoutingFormSubmissionsUUIDErrorResponse struct {
	Details []GetRoutingFormSubmissionsUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                              `json:"message"`
	Title   string                                              `json:"title"`
}

// GetRoutingFormSubmissionsUUID200ApplicationJSON - OK
type GetRoutingFormSubmissionsUUID200ApplicationJSON struct {
	// Information about a Routing Form Submission.
	Resource shared.RoutingFormSubmission `json:"resource"`
}

type GetRoutingFormSubmissionsUUIDResponse struct {
	ContentType string
	// Request is not valid
	ErrorResponse *GetRoutingFormSubmissionsUUIDErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// OK
	GetRoutingFormSubmissionsUUID200ApplicationJSONObject *GetRoutingFormSubmissionsUUID200ApplicationJSON
}
