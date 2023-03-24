// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetScheduledEventsUUIDPathParams struct {
	// The event's unique identifier
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetScheduledEventsUUIDRequest struct {
	PathParams GetScheduledEventsUUIDPathParams
}

type GetScheduledEventsUUID403ApplicationJSONMessageEnum string

const (
	GetScheduledEventsUUID403ApplicationJSONMessageEnumYouAreNotAllowedToViewThisEvent GetScheduledEventsUUID403ApplicationJSONMessageEnum = "You are not allowed to view this event"
)

type GetScheduledEventsUUID403ApplicationJSONTitleEnum string

const (
	GetScheduledEventsUUID403ApplicationJSONTitleEnumPermissionDenied GetScheduledEventsUUID403ApplicationJSONTitleEnum = "Permission Denied"
)

// GetScheduledEventsUUID403ApplicationJSON - Permission Denied
type GetScheduledEventsUUID403ApplicationJSON struct {
	Message *GetScheduledEventsUUID403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *GetScheduledEventsUUID403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type GetScheduledEventsUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetScheduledEventsUUIDErrorResponse - Error Object
type GetScheduledEventsUUIDErrorResponse struct {
	Details []GetScheduledEventsUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                       `json:"message"`
	Title   string                                       `json:"title"`
}

// GetScheduledEventsUUID200ApplicationJSON - OK
type GetScheduledEventsUUID200ApplicationJSON struct {
	// Information about a scheduled meeting
	Resource shared.Event `json:"resource"`
}

type GetScheduledEventsUUIDResponse struct {
	ContentType string
	// Cannot authenticate caller
	ErrorResponse *GetScheduledEventsUUIDErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// OK
	GetScheduledEventsUUID200ApplicationJSONObject *GetScheduledEventsUUID200ApplicationJSON
	// Permission Denied
	GetScheduledEventsUUID403ApplicationJSONObject *GetScheduledEventsUUID403ApplicationJSON
}
