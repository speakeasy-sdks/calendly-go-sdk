// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetEventTypesRequest struct {
	// Return only active event types if true, only inactive if false, or all event types if this parameter is omitted.
	Active *bool `queryParam:"style=form,explode=true,name=active"`
	// Return only admin managed event types if true, exclude admin managed event types if false, or include all event types if this parameter is omitted.
	AdminManaged *bool `queryParam:"style=form,explode=true,name=admin_managed"`
	// The number of rows to return
	Count *float64 `queryParam:"style=form,explode=true,name=count"`
	// View available personal, team, and organization event types associated with the organization's URI.
	Organization *string `queryParam:"style=form,explode=true,name=organization"`
	// The token to pass to get the next or previous portion of the collection
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
	// Order results by the specified field and direction. Accepts comma-separated list of {field}:{direction} values.
	// Supported fields are: name.
	// Sort direction is specified as: asc, desc.
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
	// View available personal, team, and organization event types associated with the user's URI.
	User *string `queryParam:"style=form,explode=true,name=user"`
}

type GetEventTypes403ApplicationJSONMessageEnum string

const (
	GetEventTypes403ApplicationJSONMessageEnumThisUserIsNotInYourOrganization            GetEventTypes403ApplicationJSONMessageEnum = "This user is not in your organization"
	GetEventTypes403ApplicationJSONMessageEnumYouDoNotHavePermission                     GetEventTypes403ApplicationJSONMessageEnum = "You do not have permission"
	GetEventTypes403ApplicationJSONMessageEnumYouDoNotHavePermissionToAccessThisResource GetEventTypes403ApplicationJSONMessageEnum = "You do not have permission to access this resource."
)

func (e *GetEventTypes403ApplicationJSONMessageEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "This user is not in your organization":
		fallthrough
	case "You do not have permission":
		fallthrough
	case "You do not have permission to access this resource.":
		*e = GetEventTypes403ApplicationJSONMessageEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEventTypes403ApplicationJSONMessageEnum: %s", s)
	}
}

type GetEventTypes403ApplicationJSONTitleEnum string

const (
	GetEventTypes403ApplicationJSONTitleEnumPermissionDenied GetEventTypes403ApplicationJSONTitleEnum = "Permission Denied"
)

func (e *GetEventTypes403ApplicationJSONTitleEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Permission Denied":
		*e = GetEventTypes403ApplicationJSONTitleEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEventTypes403ApplicationJSONTitleEnum: %s", s)
	}
}

// GetEventTypes403ApplicationJSON - Permission Denied
type GetEventTypes403ApplicationJSON struct {
	Message *GetEventTypes403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *GetEventTypes403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type GetEventTypesErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetEventTypesErrorResponse - Error Object
type GetEventTypesErrorResponse struct {
	Details []GetEventTypesErrorResponseDetails `json:"details,omitempty"`
	Message string                              `json:"message"`
	Title   string                              `json:"title"`
}

// GetEventTypes200ApplicationJSON - Service response
type GetEventTypes200ApplicationJSON struct {
	Collection []shared.EventType `json:"collection"`
	Pagination shared.Pagination  `json:"pagination"`
}

type GetEventTypesResponse struct {
	ContentType string
	// Request is not valid
	ErrorResponse *GetEventTypesErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// OK
	GetEventTypes200ApplicationJSONObject *GetEventTypes200ApplicationJSON
	// Permission Denied
	GetEventTypes403ApplicationJSONObject *GetEventTypes403ApplicationJSON
}
