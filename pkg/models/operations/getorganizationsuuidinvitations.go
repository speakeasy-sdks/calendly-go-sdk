package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetOrganizationsUUIDInvitationsPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetOrganizationsUUIDInvitationsStatusEnum string

const (
	GetOrganizationsUUIDInvitationsStatusEnumPending  GetOrganizationsUUIDInvitationsStatusEnum = "pending"
	GetOrganizationsUUIDInvitationsStatusEnumAccepted GetOrganizationsUUIDInvitationsStatusEnum = "accepted"
	GetOrganizationsUUIDInvitationsStatusEnumDeclined GetOrganizationsUUIDInvitationsStatusEnum = "declined"
)

type GetOrganizationsUUIDInvitationsQueryParams struct {
	Count     *float64                                   `queryParam:"style=form,explode=true,name=count"`
	Email     *string                                    `queryParam:"style=form,explode=true,name=email"`
	PageToken *string                                    `queryParam:"style=form,explode=true,name=page_token"`
	Sort      *string                                    `queryParam:"style=form,explode=true,name=sort"`
	Status    *GetOrganizationsUUIDInvitationsStatusEnum `queryParam:"style=form,explode=true,name=status"`
}

type GetOrganizationsUUIDInvitationsRequest struct {
	PathParams  GetOrganizationsUUIDInvitationsPathParams
	QueryParams GetOrganizationsUUIDInvitationsQueryParams
}

type GetOrganizationsUUIDInvitationsErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetOrganizationsUUIDInvitationsErrorResponse
// Error Object
type GetOrganizationsUUIDInvitationsErrorResponse struct {
	Details []GetOrganizationsUUIDInvitationsErrorResponseDetails `json:"details,omitempty"`
	Message string                                                `json:"message"`
	Title   string                                                `json:"title"`
}

type GetOrganizationsUUIDInvitations200ApplicationJSON struct {
	Collection []shared.OrganizationInvitation `json:"collection"`
	Pagination shared.Pagination               `json:"pagination"`
}

type GetOrganizationsUUIDInvitationsResponse struct {
	ContentType                                             string
	ErrorResponse                                           *GetOrganizationsUUIDInvitationsErrorResponse
	StatusCode                                              int
	RawResponse                                             *http.Response
	GetOrganizationsUUIDInvitations200ApplicationJSONObject *GetOrganizationsUUIDInvitations200ApplicationJSON
}
