// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetUserBusyTimesQueryParams struct {
	// End time of the requested availability range
	EndTime string `queryParam:"style=form,explode=true,name=end_time"`
	// Start time of the requested availability range
	StartTime string `queryParam:"style=form,explode=true,name=start_time"`
	// The uri associated with the user
	User string `queryParam:"style=form,explode=true,name=user"`
}

type GetUserBusyTimesRequest struct {
	QueryParams GetUserBusyTimesQueryParams
}

type GetUserBusyTimes403ApplicationJSONMessageEnum string

const (
	GetUserBusyTimes403ApplicationJSONMessageEnumThisUserIsNotInYourOrganization GetUserBusyTimes403ApplicationJSONMessageEnum = "This user is not in your organization"
)

type GetUserBusyTimes403ApplicationJSONTitleEnum string

const (
	GetUserBusyTimes403ApplicationJSONTitleEnumPermissionDenied GetUserBusyTimes403ApplicationJSONTitleEnum = "Permission Denied"
)

// GetUserBusyTimes403ApplicationJSON - Permission Denied
type GetUserBusyTimes403ApplicationJSON struct {
	Message *GetUserBusyTimes403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *GetUserBusyTimes403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type GetUserBusyTimesErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetUserBusyTimesErrorResponse - Error Object
type GetUserBusyTimesErrorResponse struct {
	Details []GetUserBusyTimesErrorResponseDetails `json:"details,omitempty"`
	Message string                                 `json:"message"`
	Title   string                                 `json:"title"`
}

// GetUserBusyTimes200ApplicationJSON - Service Response
type GetUserBusyTimes200ApplicationJSON struct {
	// The set of internal and external scheduled calendar events matching the criteria
	Collection []shared.UserBusyTime `json:"collection"`
}

type GetUserBusyTimesResponse struct {
	ContentType string
	// Request is not valid
	ErrorResponse *GetUserBusyTimesErrorResponse
	// Unable to access external calendar
	ErrorResponse1 *shared.ErrorResponse
	StatusCode     int
	RawResponse    *http.Response
	// OK
	GetUserBusyTimes200ApplicationJSONObject *GetUserBusyTimes200ApplicationJSON
	// Permission Denied
	GetUserBusyTimes403ApplicationJSONObject *GetUserBusyTimes403ApplicationJSON
}
