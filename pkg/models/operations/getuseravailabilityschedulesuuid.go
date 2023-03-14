package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetUserAvailabilitySchedulesUUIDPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetUserAvailabilitySchedulesUUIDRequest struct {
	PathParams GetUserAvailabilitySchedulesUUIDPathParams
}

type GetUserAvailabilitySchedulesUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetUserAvailabilitySchedulesUUIDErrorResponse
// Error Object
type GetUserAvailabilitySchedulesUUIDErrorResponse struct {
	Details []GetUserAvailabilitySchedulesUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                                 `json:"message"`
	Title   string                                                 `json:"title"`
}

type GetUserAvailabilitySchedulesUUID403ApplicationJSONMessageEnum string

const (
	GetUserAvailabilitySchedulesUUID403ApplicationJSONMessageEnumThisUserIsNotInYourOrganization GetUserAvailabilitySchedulesUUID403ApplicationJSONMessageEnum = "This user is not in your organization"
)

type GetUserAvailabilitySchedulesUUID403ApplicationJSONTitleEnum string

const (
	GetUserAvailabilitySchedulesUUID403ApplicationJSONTitleEnumPermissionDenied GetUserAvailabilitySchedulesUUID403ApplicationJSONTitleEnum = "Permission Denied"
)

type GetUserAvailabilitySchedulesUUID403ApplicationJSON struct {
	Message *GetUserAvailabilitySchedulesUUID403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *GetUserAvailabilitySchedulesUUID403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type GetUserAvailabilitySchedulesUUID200ApplicationJSON struct {
	Resource shared.AvailabilitySchedule `json:"resource"`
}

type GetUserAvailabilitySchedulesUUIDResponse struct {
	ContentType                                              string
	ErrorResponse                                            *GetUserAvailabilitySchedulesUUIDErrorResponse
	StatusCode                                               int
	RawResponse                                              *http.Response
	GetUserAvailabilitySchedulesUUID200ApplicationJSONObject *GetUserAvailabilitySchedulesUUID200ApplicationJSON
	GetUserAvailabilitySchedulesUUID403ApplicationJSONObject *GetUserAvailabilitySchedulesUUID403ApplicationJSON
}
