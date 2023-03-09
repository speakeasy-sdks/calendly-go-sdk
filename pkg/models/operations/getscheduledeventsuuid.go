package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetScheduledEventsUUIDSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetScheduledEventsUUIDPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetScheduledEventsUUIDRequest struct {
	PathParams GetScheduledEventsUUIDPathParams
	Security   GetScheduledEventsUUIDSecurity
}

type GetScheduledEventsUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetScheduledEventsUUIDErrorResponse
// Error Object
type GetScheduledEventsUUIDErrorResponse struct {
	Details []GetScheduledEventsUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                       `json:"message"`
	Title   string                                       `json:"title"`
}

type GetScheduledEventsUUID403ApplicationJSONMessageEnum string

const (
	GetScheduledEventsUUID403ApplicationJSONMessageEnumYouAreNotAllowedToViewThisEvent GetScheduledEventsUUID403ApplicationJSONMessageEnum = "You are not allowed to view this event"
)

type GetScheduledEventsUUID403ApplicationJSONTitleEnum string

const (
	GetScheduledEventsUUID403ApplicationJSONTitleEnumPermissionDenied GetScheduledEventsUUID403ApplicationJSONTitleEnum = "Permission Denied"
)

type GetScheduledEventsUUID403ApplicationJSON struct {
	Message *GetScheduledEventsUUID403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *GetScheduledEventsUUID403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type GetScheduledEventsUUID200ApplicationJSON struct {
	Resource shared.Event `json:"resource"`
}

type GetScheduledEventsUUIDResponse struct {
	ContentType                                    string
	ErrorResponse                                  *GetScheduledEventsUUIDErrorResponse
	StatusCode                                     int
	RawResponse                                    *http.Response
	GetScheduledEventsUUID200ApplicationJSONObject *GetScheduledEventsUUID200ApplicationJSON
	GetScheduledEventsUUID403ApplicationJSONObject *GetScheduledEventsUUID403ApplicationJSON
}
