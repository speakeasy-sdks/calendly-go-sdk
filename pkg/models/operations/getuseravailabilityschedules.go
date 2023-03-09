package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetUserAvailabilitySchedulesSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetUserAvailabilitySchedulesQueryParams struct {
	User *string `queryParam:"style=form,explode=true,name=user"`
}

type GetUserAvailabilitySchedulesRequest struct {
	QueryParams GetUserAvailabilitySchedulesQueryParams
	Security    GetUserAvailabilitySchedulesSecurity
}

type GetUserAvailabilitySchedulesErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetUserAvailabilitySchedulesErrorResponse
// Error Object
type GetUserAvailabilitySchedulesErrorResponse struct {
	Details []GetUserAvailabilitySchedulesErrorResponseDetails `json:"details,omitempty"`
	Message string                                             `json:"message"`
	Title   string                                             `json:"title"`
}

type GetUserAvailabilitySchedules403ApplicationJSONMessageEnum string

const (
	GetUserAvailabilitySchedules403ApplicationJSONMessageEnumThisUserIsNotInYourOrganization GetUserAvailabilitySchedules403ApplicationJSONMessageEnum = "This user is not in your organization"
)

type GetUserAvailabilitySchedules403ApplicationJSONTitleEnum string

const (
	GetUserAvailabilitySchedules403ApplicationJSONTitleEnumPermissionDenied GetUserAvailabilitySchedules403ApplicationJSONTitleEnum = "Permission Denied"
)

type GetUserAvailabilitySchedules403ApplicationJSON struct {
	Message *GetUserAvailabilitySchedules403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *GetUserAvailabilitySchedules403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type GetUserAvailabilitySchedules200ApplicationJSON struct {
	Collection []shared.AvailabilitySchedule `json:"collection"`
}

type GetUserAvailabilitySchedulesResponse struct {
	ContentType                                          string
	ErrorResponse                                        *GetUserAvailabilitySchedulesErrorResponse
	StatusCode                                           int
	RawResponse                                          *http.Response
	GetUserAvailabilitySchedules200ApplicationJSONObject *GetUserAvailabilitySchedules200ApplicationJSON
	GetUserAvailabilitySchedules403ApplicationJSONObject *GetUserAvailabilitySchedules403ApplicationJSON
}
