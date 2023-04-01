// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type RevokeUsersOrganizationInvitationRequest struct {
	// The organization’s unique identifier
	OrgUUID string `pathParam:"style=simple,explode=false,name=org_uuid"`
	// The organization invitation's unique identifier
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum string

const (
	RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnumYouCannotPerformThisActionForAnOrganizationWithScimEnabled RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum = "You cannot perform this action for an organization with SCIM enabled."
	RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnumYouDoNotHavePermission                                     RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum = "You do not have permission"
	RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnumYouDoNotHavePermissionToAccessThisResource                 RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum = "You do not have permission to access this resource."
)

func (e *RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "You cannot perform this action for an organization with SCIM enabled.":
		fallthrough
	case "You do not have permission":
		fallthrough
	case "You do not have permission to access this resource.":
		*e = RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum: %s", s)
	}
}

type RevokeUsersOrganizationInvitation403ApplicationJSONTitleEnum string

const (
	RevokeUsersOrganizationInvitation403ApplicationJSONTitleEnumPermissionDenied RevokeUsersOrganizationInvitation403ApplicationJSONTitleEnum = "Permission Denied"
)

func (e *RevokeUsersOrganizationInvitation403ApplicationJSONTitleEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Permission Denied":
		*e = RevokeUsersOrganizationInvitation403ApplicationJSONTitleEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for RevokeUsersOrganizationInvitation403ApplicationJSONTitleEnum: %s", s)
	}
}

// RevokeUsersOrganizationInvitation403ApplicationJSON - Permission Denied
type RevokeUsersOrganizationInvitation403ApplicationJSON struct {
	Message *RevokeUsersOrganizationInvitation403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *RevokeUsersOrganizationInvitation403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type RevokeUsersOrganizationInvitationErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// RevokeUsersOrganizationInvitationErrorResponse - Error Object
type RevokeUsersOrganizationInvitationErrorResponse struct {
	Details []RevokeUsersOrganizationInvitationErrorResponseDetails `json:"details,omitempty"`
	Message string                                                  `json:"message"`
	Title   string                                                  `json:"title"`
}

type RevokeUsersOrganizationInvitationResponse struct {
	ContentType string
	// Cannot authenticate caller
	ErrorResponse1 *RevokeUsersOrganizationInvitationErrorResponse
	// Bad Request
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// Permission Denied
	RevokeUsersOrganizationInvitation403ApplicationJSONObject *RevokeUsersOrganizationInvitation403ApplicationJSON
}
