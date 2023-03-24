// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type PostScheduledEventsUUIDCancellationPathParams struct {
	// The event's unique indentifier
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type PostScheduledEventsUUIDCancellationRawRequest struct {
	PathParams PostScheduledEventsUUIDCancellationPathParams
	// Optional cancellation reason.
	Request []byte `request:"mediaType=application/xml"`
}

type PostScheduledEventsUUIDCancellationRaw403ApplicationJSONMessageEnum string

const (
	PostScheduledEventsUUIDCancellationRaw403ApplicationJSONMessageEnumYouAreNotAllowedToCancelThisEvent PostScheduledEventsUUIDCancellationRaw403ApplicationJSONMessageEnum = "You are not allowed to cancel this event"
	PostScheduledEventsUUIDCancellationRaw403ApplicationJSONMessageEnumEventInThePast                    PostScheduledEventsUUIDCancellationRaw403ApplicationJSONMessageEnum = "Event in the past"
	PostScheduledEventsUUIDCancellationRaw403ApplicationJSONMessageEnumEventIsAlreadyCanceled            PostScheduledEventsUUIDCancellationRaw403ApplicationJSONMessageEnum = "Event is already canceled"
)

type PostScheduledEventsUUIDCancellationRaw403ApplicationJSONTitleEnum string

const (
	PostScheduledEventsUUIDCancellationRaw403ApplicationJSONTitleEnumPermissionDenied PostScheduledEventsUUIDCancellationRaw403ApplicationJSONTitleEnum = "Permission Denied"
)

// PostScheduledEventsUUIDCancellationRaw403ApplicationJSON - Caller not authorized to perform this action
type PostScheduledEventsUUIDCancellationRaw403ApplicationJSON struct {
	Message *PostScheduledEventsUUIDCancellationRaw403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *PostScheduledEventsUUIDCancellationRaw403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type PostScheduledEventsUUIDCancellationRawErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostScheduledEventsUUIDCancellationRawErrorResponse - Error Object
type PostScheduledEventsUUIDCancellationRawErrorResponse struct {
	Details []PostScheduledEventsUUIDCancellationRawErrorResponseDetails `json:"details,omitempty"`
	Message string                                                       `json:"message"`
	Title   string                                                       `json:"title"`
}

// PostScheduledEventsUUIDCancellationRaw201ApplicationJSON - Created
type PostScheduledEventsUUIDCancellationRaw201ApplicationJSON struct {
	// Provides data pertaining to the cancellation of the Event
	Resource shared.Cancellation `json:"resource"`
}

type PostScheduledEventsUUIDCancellationRawResponse struct {
	ContentType string
	// Request is not valid
	ErrorResponse *PostScheduledEventsUUIDCancellationRawErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// Created
	PostScheduledEventsUUIDCancellationRaw201ApplicationJSONObject *PostScheduledEventsUUIDCancellationRaw201ApplicationJSON
	// Caller not authorized to perform this action
	PostScheduledEventsUUIDCancellationRaw403ApplicationJSONObject *PostScheduledEventsUUIDCancellationRaw403ApplicationJSON
}
