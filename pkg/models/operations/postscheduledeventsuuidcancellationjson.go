package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type PostScheduledEventsUUIDCancellationSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type PostScheduledEventsUUIDCancellationPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type PostScheduledEventsUUIDCancellationApplicationJSON struct {
	Reason *string `json:"reason,omitempty"`
}

type PostScheduledEventsUUIDCancellationJSONRequest struct {
	PathParams PostScheduledEventsUUIDCancellationPathParams
	Request    *PostScheduledEventsUUIDCancellationApplicationJSON `request:"mediaType=application/json"`
	Security   PostScheduledEventsUUIDCancellationSecurity
}

type PostScheduledEventsUUIDCancellationJSONErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostScheduledEventsUUIDCancellationJSONErrorResponse
// Error Object
type PostScheduledEventsUUIDCancellationJSONErrorResponse struct {
	Details []PostScheduledEventsUUIDCancellationJSONErrorResponseDetails `json:"details,omitempty"`
	Message string                                                        `json:"message"`
	Title   string                                                        `json:"title"`
}

type PostScheduledEventsUUIDCancellationJSON403ApplicationJSONMessageEnum string

const (
	PostScheduledEventsUUIDCancellationJSON403ApplicationJSONMessageEnumYouAreNotAllowedToCancelThisEvent PostScheduledEventsUUIDCancellationJSON403ApplicationJSONMessageEnum = "You are not allowed to cancel this event"
	PostScheduledEventsUUIDCancellationJSON403ApplicationJSONMessageEnumEventInThePast                    PostScheduledEventsUUIDCancellationJSON403ApplicationJSONMessageEnum = "Event in the past"
	PostScheduledEventsUUIDCancellationJSON403ApplicationJSONMessageEnumEventIsAlreadyCanceled            PostScheduledEventsUUIDCancellationJSON403ApplicationJSONMessageEnum = "Event is already canceled"
)

type PostScheduledEventsUUIDCancellationJSON403ApplicationJSONTitleEnum string

const (
	PostScheduledEventsUUIDCancellationJSON403ApplicationJSONTitleEnumPermissionDenied PostScheduledEventsUUIDCancellationJSON403ApplicationJSONTitleEnum = "Permission Denied"
)

type PostScheduledEventsUUIDCancellationJSON403ApplicationJSON struct {
	Message *PostScheduledEventsUUIDCancellationJSON403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *PostScheduledEventsUUIDCancellationJSON403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type PostScheduledEventsUUIDCancellationJSON201ApplicationJSON struct {
	Resource shared.Cancellation `json:"resource"`
}

type PostScheduledEventsUUIDCancellationJSONResponse struct {
	ContentType                                                     string
	ErrorResponse                                                   *PostScheduledEventsUUIDCancellationJSONErrorResponse
	StatusCode                                                      int
	RawResponse                                                     *http.Response
	PostScheduledEventsUUIDCancellationJSON201ApplicationJSONObject *PostScheduledEventsUUIDCancellationJSON201ApplicationJSON
	PostScheduledEventsUUIDCancellationJSON403ApplicationJSONObject *PostScheduledEventsUUIDCancellationJSON403ApplicationJSON
}
