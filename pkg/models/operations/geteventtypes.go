package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetEventTypesSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetEventTypesQueryParams struct {
	Active       *bool    `queryParam:"style=form,explode=true,name=active"`
	AdminManaged *bool    `queryParam:"style=form,explode=true,name=admin_managed"`
	Count        *float64 `queryParam:"style=form,explode=true,name=count"`
	Organization *string  `queryParam:"style=form,explode=true,name=organization"`
	PageToken    *string  `queryParam:"style=form,explode=true,name=page_token"`
	Sort         *string  `queryParam:"style=form,explode=true,name=sort"`
	User         *string  `queryParam:"style=form,explode=true,name=user"`
}

type GetEventTypesRequest struct {
	QueryParams GetEventTypesQueryParams
	Security    GetEventTypesSecurity
}

type GetEventTypesErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetEventTypesErrorResponse
// Error Object
type GetEventTypesErrorResponse struct {
	Details []GetEventTypesErrorResponseDetails `json:"details,omitempty"`
	Message string                              `json:"message"`
	Title   string                              `json:"title"`
}

type GetEventTypes403ApplicationJSONMessageEnum string

const (
	GetEventTypes403ApplicationJSONMessageEnumThisUserIsNotInYourOrganization            GetEventTypes403ApplicationJSONMessageEnum = "This user is not in your organization"
	GetEventTypes403ApplicationJSONMessageEnumYouDoNotHavePermission                     GetEventTypes403ApplicationJSONMessageEnum = "You do not have permission"
	GetEventTypes403ApplicationJSONMessageEnumYouDoNotHavePermissionToAccessThisResource GetEventTypes403ApplicationJSONMessageEnum = "You do not have permission to access this resource."
)

type GetEventTypes403ApplicationJSONTitleEnum string

const (
	GetEventTypes403ApplicationJSONTitleEnumPermissionDenied GetEventTypes403ApplicationJSONTitleEnum = "Permission Denied"
)

type GetEventTypes403ApplicationJSON struct {
	Message *GetEventTypes403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *GetEventTypes403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

// GetEventTypes200ApplicationJSON
// Service response
type GetEventTypes200ApplicationJSON struct {
	Collection []shared.EventType `json:"collection"`
	Pagination shared.Pagination  `json:"pagination"`
}

type GetEventTypesResponse struct {
	ContentType                           string
	ErrorResponse                         *GetEventTypesErrorResponse
	StatusCode                            int
	RawResponse                           *http.Response
	GetEventTypes200ApplicationJSONObject *GetEventTypes200ApplicationJSON
	GetEventTypes403ApplicationJSONObject *GetEventTypes403ApplicationJSON
}
