package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type PostScheduledEventsUUIDCancellationPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type PostScheduledEventsUUIDCancellationRawRequest struct {
	PathParams PostScheduledEventsUUIDCancellationPathParams
	Request    []byte `request:"mediaType=application/xml"`
}

type PostScheduledEventsUUIDCancellationRawErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostScheduledEventsUUIDCancellationRawErrorResponse
// Error Object
type PostScheduledEventsUUIDCancellationRawErrorResponse struct {
	Details []PostScheduledEventsUUIDCancellationRawErrorResponseDetails `json:"details,omitempty"`
	Message string                                                       `json:"message"`
	Title   string                                                       `json:"title"`
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

type PostScheduledEventsUUIDCancellationRaw403ApplicationJSON struct {
	Message *PostScheduledEventsUUIDCancellationRaw403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *PostScheduledEventsUUIDCancellationRaw403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type PostScheduledEventsUUIDCancellationRaw201ApplicationJSON struct {
	Resource shared.Cancellation `json:"resource"`
}

type PostScheduledEventsUUIDCancellationRawResponse struct {
	ContentType                                                    string
	ErrorResponse                                                  *PostScheduledEventsUUIDCancellationRawErrorResponse
	StatusCode                                                     int
	RawResponse                                                    *http.Response
	PostScheduledEventsUUIDCancellationRaw201ApplicationJSONObject *PostScheduledEventsUUIDCancellationRaw201ApplicationJSON
	PostScheduledEventsUUIDCancellationRaw403ApplicationJSONObject *PostScheduledEventsUUIDCancellationRaw403ApplicationJSON
}
