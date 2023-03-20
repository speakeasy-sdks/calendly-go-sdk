package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetUserBusyTimesQueryParams struct {
	EndTime   string `queryParam:"style=form,explode=true,name=end_time"`
	StartTime string `queryParam:"style=form,explode=true,name=start_time"`
	User      string `queryParam:"style=form,explode=true,name=user"`
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

type GetUserBusyTimes403ApplicationJSON struct {
	Message *GetUserBusyTimes403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *GetUserBusyTimes403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type GetUserBusyTimesErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetUserBusyTimesErrorResponse
// Error Object
type GetUserBusyTimesErrorResponse struct {
	Details []GetUserBusyTimesErrorResponseDetails `json:"details,omitempty"`
	Message string                                 `json:"message"`
	Title   string                                 `json:"title"`
}

// GetUserBusyTimes200ApplicationJSON
// Service Response
type GetUserBusyTimes200ApplicationJSON struct {
	Collection []shared.UserBusyTime `json:"collection"`
}

type GetUserBusyTimesResponse struct {
	ContentType                              string
	ErrorResponse                            *GetUserBusyTimesErrorResponse
	ErrorResponse1                           *shared.ErrorResponse
	StatusCode                               int
	RawResponse                              *http.Response
	GetUserBusyTimes200ApplicationJSONObject *GetUserBusyTimes200ApplicationJSON
	GetUserBusyTimes403ApplicationJSONObject *GetUserBusyTimes403ApplicationJSON
}
