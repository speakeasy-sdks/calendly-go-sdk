// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetEventTypeAvailableTimesRequest struct {
	// End time of the requested availability range.
	EndTime *string `queryParam:"style=form,explode=true,name=end_time"`
	// The uri associated with the event type
	EventType *string `queryParam:"style=form,explode=true,name=event_type"`
	// Start time of the requested availability range.
	StartTime *string `queryParam:"style=form,explode=true,name=start_time"`
}

type GetEventTypeAvailableTimesErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetEventTypeAvailableTimesErrorResponse - Error Object
type GetEventTypeAvailableTimesErrorResponse struct {
	Details []GetEventTypeAvailableTimesErrorResponseDetails `json:"details,omitempty"`
	Message string                                           `json:"message"`
	Title   string                                           `json:"title"`
}

// GetEventTypeAvailableTimes200ApplicationJSON - Service Response
type GetEventTypeAvailableTimes200ApplicationJSON struct {
	// The set of available times for the event type matching the criteria
	Collection []shared.EventTypeAvailableTime `json:"collection"`
}

type GetEventTypeAvailableTimesResponse struct {
	ContentType string
	// Request is not valid
	ErrorResponse *GetEventTypeAvailableTimesErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// OK
	GetEventTypeAvailableTimes200ApplicationJSONObject *GetEventTypeAvailableTimes200ApplicationJSON
}
