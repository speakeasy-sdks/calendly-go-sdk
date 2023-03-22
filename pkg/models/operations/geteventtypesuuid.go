// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetEventTypesUUIDPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetEventTypesUUIDRequest struct {
	PathParams GetEventTypesUUIDPathParams
}

type GetEventTypesUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetEventTypesUUIDErrorResponse - Error Object
type GetEventTypesUUIDErrorResponse struct {
	Details []GetEventTypesUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                  `json:"message"`
	Title   string                                  `json:"title"`
}

// GetEventTypesUUID200ApplicationJSON - OK
type GetEventTypesUUID200ApplicationJSON struct {
	// A configuration for an Event
	Resource shared.EventType `json:"resource"`
}

type GetEventTypesUUIDResponse struct {
	ContentType string
	// Request is not valid
	ErrorResponse *GetEventTypesUUIDErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// OK
	GetEventTypesUUID200ApplicationJSONObject *GetEventTypesUUID200ApplicationJSON
}
