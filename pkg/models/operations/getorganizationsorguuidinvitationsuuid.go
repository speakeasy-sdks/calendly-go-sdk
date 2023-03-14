package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetOrganizationsOrgUUIDInvitationsUUIDPathParams struct {
	OrgUUID string `pathParam:"style=simple,explode=false,name=org_uuid"`
	UUID    string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetOrganizationsOrgUUIDInvitationsUUIDRequest struct {
	PathParams GetOrganizationsOrgUUIDInvitationsUUIDPathParams
}

type GetOrganizationsOrgUUIDInvitationsUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetOrganizationsOrgUUIDInvitationsUUIDErrorResponse
// Error Object
type GetOrganizationsOrgUUIDInvitationsUUIDErrorResponse struct {
	Details []GetOrganizationsOrgUUIDInvitationsUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                                       `json:"message"`
	Title   string                                                       `json:"title"`
}

type GetOrganizationsOrgUUIDInvitationsUUID200ApplicationJSON struct {
	Resource shared.OrganizationInvitation `json:"resource"`
}

type GetOrganizationsOrgUUIDInvitationsUUIDResponse struct {
	ContentType                                                    string
	ErrorResponse                                                  *GetOrganizationsOrgUUIDInvitationsUUIDErrorResponse
	StatusCode                                                     int
	RawResponse                                                    *http.Response
	GetOrganizationsOrgUUIDInvitationsUUID200ApplicationJSONObject *GetOrganizationsOrgUUIDInvitationsUUID200ApplicationJSON
}
