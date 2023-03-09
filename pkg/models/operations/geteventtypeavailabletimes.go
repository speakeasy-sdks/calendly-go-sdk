package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetEventTypeAvailableTimesSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetEventTypeAvailableTimesQueryParams struct {
	EndTime   *string `queryParam:"style=form,explode=true,name=end_time"`
	EventType *string `queryParam:"style=form,explode=true,name=event_type"`
	StartTime *string `queryParam:"style=form,explode=true,name=start_time"`
}

type GetEventTypeAvailableTimesRequest struct {
	QueryParams GetEventTypeAvailableTimesQueryParams
	Security    GetEventTypeAvailableTimesSecurity
}

type GetEventTypeAvailableTimesErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetEventTypeAvailableTimesErrorResponse
// Error Object
type GetEventTypeAvailableTimesErrorResponse struct {
	Details []GetEventTypeAvailableTimesErrorResponseDetails `json:"details,omitempty"`
	Message string                                           `json:"message"`
	Title   string                                           `json:"title"`
}

// GetEventTypeAvailableTimes200ApplicationJSON
// Service Response
type GetEventTypeAvailableTimes200ApplicationJSON struct {
	Collection []shared.EventTypeAvailableTime `json:"collection"`
}

type GetEventTypeAvailableTimesResponse struct {
	ContentType                                        string
	ErrorResponse                                      *GetEventTypeAvailableTimesErrorResponse
	StatusCode                                         int
	RawResponse                                        *http.Response
	GetEventTypeAvailableTimes200ApplicationJSONObject *GetEventTypeAvailableTimes200ApplicationJSON
}
