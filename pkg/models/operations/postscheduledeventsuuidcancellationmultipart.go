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

type PostScheduledEventsUUIDCancellationMultipartRequest struct {
	PathParams PostScheduledEventsUUIDCancellationPathParams
	Request    map[string]interface{} `request:"mediaType=multipart/form-data"`
	Security   PostScheduledEventsUUIDCancellationSecurity
}

type PostScheduledEventsUUIDCancellationMultipartErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostScheduledEventsUUIDCancellationMultipartErrorResponse
// Error Object
type PostScheduledEventsUUIDCancellationMultipartErrorResponse struct {
	Details []PostScheduledEventsUUIDCancellationMultipartErrorResponseDetails `json:"details,omitempty"`
	Message string                                                             `json:"message"`
	Title   string                                                             `json:"title"`
}

type PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONMessageEnum string

const (
	PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONMessageEnumYouAreNotAllowedToCancelThisEvent PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONMessageEnum = "You are not allowed to cancel this event"
	PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONMessageEnumEventInThePast                    PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONMessageEnum = "Event in the past"
	PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONMessageEnumEventIsAlreadyCanceled            PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONMessageEnum = "Event is already canceled"
)

type PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONTitleEnum string

const (
	PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONTitleEnumPermissionDenied PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONTitleEnum = "Permission Denied"
)

type PostScheduledEventsUUIDCancellationMultipart403ApplicationJSON struct {
	Message *PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type PostScheduledEventsUUIDCancellationMultipart201ApplicationJSON struct {
	Resource shared.Cancellation `json:"resource"`
}

type PostScheduledEventsUUIDCancellationMultipartResponse struct {
	ContentType                                                          string
	ErrorResponse                                                        *PostScheduledEventsUUIDCancellationMultipartErrorResponse
	StatusCode                                                           int
	RawResponse                                                          *http.Response
	PostScheduledEventsUUIDCancellationMultipart201ApplicationJSONObject *PostScheduledEventsUUIDCancellationMultipart201ApplicationJSON
	PostScheduledEventsUUIDCancellationMultipart403ApplicationJSONObject *PostScheduledEventsUUIDCancellationMultipart403ApplicationJSON
}
