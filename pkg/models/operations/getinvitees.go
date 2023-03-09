package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetInviteesSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetInviteesPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetInviteesStatusEnum string

const (
	GetInviteesStatusEnumActive   GetInviteesStatusEnum = "active"
	GetInviteesStatusEnumCanceled GetInviteesStatusEnum = "canceled"
)

type GetInviteesQueryParams struct {
	Count     *float64               `queryParam:"style=form,explode=true,name=count"`
	Email     *string                `queryParam:"style=form,explode=true,name=email"`
	PageToken *string                `queryParam:"style=form,explode=true,name=page_token"`
	Sort      *string                `queryParam:"style=form,explode=true,name=sort"`
	Status    *GetInviteesStatusEnum `queryParam:"style=form,explode=true,name=status"`
}

type GetInviteesRequest struct {
	PathParams  GetInviteesPathParams
	QueryParams GetInviteesQueryParams
	Security    GetInviteesSecurity
}

type GetInviteesErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetInviteesErrorResponse
// Error Object
type GetInviteesErrorResponse struct {
	Details []GetInviteesErrorResponseDetails `json:"details,omitempty"`
	Message string                            `json:"message"`
	Title   string                            `json:"title"`
}

type GetInvitees403ApplicationJSONMessageEnum string

const (
	GetInvitees403ApplicationJSONMessageEnumYouAreNotAllowedToViewThisEvent GetInvitees403ApplicationJSONMessageEnum = "You are not allowed to view this event"
)

type GetInvitees403ApplicationJSONTitleEnum string

const (
	GetInvitees403ApplicationJSONTitleEnumPermissionDenied GetInvitees403ApplicationJSONTitleEnum = "Permission Denied"
)

type GetInvitees403ApplicationJSON struct {
	Message *GetInvitees403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *GetInvitees403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

// GetInvitees200ApplicationJSON
// Service response
type GetInvitees200ApplicationJSON struct {
	Collection []shared.Invitee  `json:"collection"`
	Pagination shared.Pagination `json:"pagination"`
}

type GetInviteesResponse struct {
	ContentType                         string
	ErrorResponse                       *GetInviteesErrorResponse
	StatusCode                          int
	RawResponse                         *http.Response
	GetInvitees200ApplicationJSONObject *GetInvitees200ApplicationJSON
	GetInvitees403ApplicationJSONObject *GetInvitees403ApplicationJSON
}
