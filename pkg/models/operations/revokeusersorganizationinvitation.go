package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type RevokeUsersOrganizationInvitationPathParams struct {
	OrgUUID string `pathParam:"style=simple,explode=false,name=org_uuid"`
	UUID    string `pathParam:"style=simple,explode=false,name=uuid"`
}

type RevokeUsersOrganizationInvitationRequest struct {
	PathParams RevokeUsersOrganizationInvitationPathParams
}

type RevokeUsersOrganizationInvitationErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// RevokeUsersOrganizationInvitationErrorResponse
// Error Object
type RevokeUsersOrganizationInvitationErrorResponse struct {
	Details []RevokeUsersOrganizationInvitationErrorResponseDetails `json:"details,omitempty"`
	Message string                                                  `json:"message"`
	Title   string                                                  `json:"title"`
}

type RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum string

const (
	RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnumYouCannotPerformThisActionForAnOrganizationWithScimEnabled RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum = "You cannot perform this action for an organization with SCIM enabled."
	RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnumYouDoNotHavePermission                                     RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum = "You do not have permission"
	RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnumYouDoNotHavePermissionToAccessThisResource                 RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum = "You do not have permission to access this resource."
)

type RevokeUsersOrganizationInvitation403ApplicationJSONTitleEnum string

const (
	RevokeUsersOrganizationInvitation403ApplicationJSONTitleEnumPermissionDenied RevokeUsersOrganizationInvitation403ApplicationJSONTitleEnum = "Permission Denied"
)

type RevokeUsersOrganizationInvitation403ApplicationJSON struct {
	Message *RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *RevokeUsersOrganizationInvitation403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type RevokeUsersOrganizationInvitationResponse struct {
	ContentType                                               string
	ErrorResponse1                                            *RevokeUsersOrganizationInvitationErrorResponse
	ErrorResponse                                             *shared.ErrorResponse
	StatusCode                                                int
	RawResponse                                               *http.Response
	RevokeUsersOrganizationInvitation403ApplicationJSONObject *RevokeUsersOrganizationInvitation403ApplicationJSON
}
