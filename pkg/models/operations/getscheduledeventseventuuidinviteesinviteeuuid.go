package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetScheduledEventsEventUUIDInviteesInviteeUUIDPathParams struct {
	EventUUID   string `pathParam:"style=simple,explode=false,name=event_uuid"`
	InviteeUUID string `pathParam:"style=simple,explode=false,name=invitee_uuid"`
}

type GetScheduledEventsEventUUIDInviteesInviteeUUIDRequest struct {
	PathParams GetScheduledEventsEventUUIDInviteesInviteeUUIDPathParams
}

type GetScheduledEventsEventUUIDInviteesInviteeUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetScheduledEventsEventUUIDInviteesInviteeUUIDErrorResponse
// Error Object
type GetScheduledEventsEventUUIDInviteesInviteeUUIDErrorResponse struct {
	Details []GetScheduledEventsEventUUIDInviteesInviteeUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                                               `json:"message"`
	Title   string                                                               `json:"title"`
}

type GetScheduledEventsEventUUIDInviteesInviteeUUID200ApplicationJSON struct {
	Resource shared.Invitee `json:"resource"`
}

type GetScheduledEventsEventUUIDInviteesInviteeUUIDResponse struct {
	ContentType                                                            string
	ErrorResponse                                                          *GetScheduledEventsEventUUIDInviteesInviteeUUIDErrorResponse
	ErrorResponse1                                                         *shared.ErrorResponse
	StatusCode                                                             int
	RawResponse                                                            *http.Response
	GetScheduledEventsEventUUIDInviteesInviteeUUID200ApplicationJSONObject *GetScheduledEventsEventUUIDInviteesInviteeUUID200ApplicationJSON
}
