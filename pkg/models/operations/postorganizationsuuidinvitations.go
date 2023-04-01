// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type PostOrganizationsUUIDInvitationsRequestBody struct {
	// The email of the user being invited
	Email string `json:"email"`
}

type PostOrganizationsUUIDInvitationsRequest struct {
	RequestBody PostOrganizationsUUIDInvitationsRequestBody `request:"mediaType=application/json"`
	// The organization's unique identifier
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum string

const (
	PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnumYouAlreadySentAllTheInvitationsYouReAllottedBasedUponTheNumberOfSeatsPurchasedWithYourAccountPleasePurchaseMoreSeatsToSendMoreInvitations PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum = "You already sent all the invitations you're allotted based upon the number of seats purchased with your account. Please purchase more seats to send more invitations."
	PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnumYouAlreadySentAllTheInvitationsAllottedToYouWithATrialAccount                                                                             PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum = "You already sent all the invitations allotted to you with a trial account."
	PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnumYouDoNotHavePermission                                                                                                                    PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum = "You do not have permission"
	PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnumYouCannotPerformThisActionForAnOrganizationWithScimEnabled                                                                                PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum = "You cannot perform this action for an organization with SCIM enabled."
)

func (e *PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "You already sent all the invitations you're allotted based upon the number of seats purchased with your account. Please purchase more seats to send more invitations.":
		fallthrough
	case "You already sent all the invitations allotted to you with a trial account.":
		fallthrough
	case "You do not have permission":
		fallthrough
	case "You cannot perform this action for an organization with SCIM enabled.":
		*e = PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum: %s", s)
	}
}

type PostOrganizationsUUIDInvitations403ApplicationJSONTitleEnum string

const (
	PostOrganizationsUUIDInvitations403ApplicationJSONTitleEnumPermissionDenied PostOrganizationsUUIDInvitations403ApplicationJSONTitleEnum = "Permission Denied"
)

func (e *PostOrganizationsUUIDInvitations403ApplicationJSONTitleEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Permission Denied":
		*e = PostOrganizationsUUIDInvitations403ApplicationJSONTitleEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PostOrganizationsUUIDInvitations403ApplicationJSONTitleEnum: %s", s)
	}
}

// PostOrganizationsUUIDInvitations403ApplicationJSON - Permission Denied
type PostOrganizationsUUIDInvitations403ApplicationJSON struct {
	Message *PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *PostOrganizationsUUIDInvitations403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type PostOrganizationsUUIDInvitationsErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostOrganizationsUUIDInvitationsErrorResponse - Error Object
type PostOrganizationsUUIDInvitationsErrorResponse struct {
	Details []PostOrganizationsUUIDInvitationsErrorResponseDetails `json:"details,omitempty"`
	Message string                                                 `json:"message"`
	Title   string                                                 `json:"title"`
}

// PostOrganizationsUUIDInvitations201ApplicationJSON - Created
type PostOrganizationsUUIDInvitations201ApplicationJSON struct {
	// Organization Invitation object
	Resource shared.OrganizationInvitation `json:"resource"`
}

type PostOrganizationsUUIDInvitationsResponse struct {
	ContentType string
	// Cannot authenticate caller
	ErrorResponse1 *PostOrganizationsUUIDInvitationsErrorResponse
	// Bad Request
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// Created
	PostOrganizationsUUIDInvitations201ApplicationJSONObject *PostOrganizationsUUIDInvitations201ApplicationJSON
	// Permission Denied
	PostOrganizationsUUIDInvitations403ApplicationJSONObject *PostOrganizationsUUIDInvitations403ApplicationJSON
}
