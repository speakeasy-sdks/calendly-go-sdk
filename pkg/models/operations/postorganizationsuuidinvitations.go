package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type PostOrganizationsUUIDInvitationsSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type PostOrganizationsUUIDInvitationsPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type PostOrganizationsUUIDInvitationsRequestBody struct {
	Email string `json:"email"`
}

type PostOrganizationsUUIDInvitationsRequest struct {
	PathParams PostOrganizationsUUIDInvitationsPathParams
	Request    PostOrganizationsUUIDInvitationsRequestBody `request:"mediaType=application/json"`
	Security   PostOrganizationsUUIDInvitationsSecurity
}

type PostOrganizationsUUIDInvitationsErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostOrganizationsUUIDInvitationsErrorResponse
// Error Object
type PostOrganizationsUUIDInvitationsErrorResponse struct {
	Details []PostOrganizationsUUIDInvitationsErrorResponseDetails `json:"details,omitempty"`
	Message string                                                 `json:"message"`
	Title   string                                                 `json:"title"`
}

type PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum string

const (
	PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnumYouAlreadySentAllTheInvitationsYouReAllottedBasedUponTheNumberOfSeatsPurchasedWithYourAccountPleasePurchaseMoreSeatsToSendMoreInvitations PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum = "You already sent all the invitations you're allotted based upon the number of seats purchased with your account. Please purchase more seats to send more invitations."
	PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnumYouAlreadySentAllTheInvitationsAllottedToYouWithATrialAccount                                                                             PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum = "You already sent all the invitations allotted to you with a trial account."
	PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnumYouDoNotHavePermission                                                                                                                    PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum = "You do not have permission"
	PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnumYouCannotPerformThisActionForAnOrganizationWithScimEnabled                                                                                PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum = "You cannot perform this action for an organization with SCIM enabled."
)

type PostOrganizationsUUIDInvitations403ApplicationJSONTitleEnum string

const (
	PostOrganizationsUUIDInvitations403ApplicationJSONTitleEnumPermissionDenied PostOrganizationsUUIDInvitations403ApplicationJSONTitleEnum = "Permission Denied"
)

type PostOrganizationsUUIDInvitations403ApplicationJSON struct {
	Message *PostOrganizationsUUIDInvitations403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *PostOrganizationsUUIDInvitations403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type PostOrganizationsUUIDInvitations201ApplicationJSON struct {
	Resource shared.OrganizationInvitation `json:"resource"`
}

type PostOrganizationsUUIDInvitationsResponse struct {
	ContentType                                              string
	ErrorResponse1                                           *PostOrganizationsUUIDInvitationsErrorResponse
	ErrorResponse                                            *shared.ErrorResponse
	StatusCode                                               int
	RawResponse                                              *http.Response
	PostOrganizationsUUIDInvitations201ApplicationJSONObject *PostOrganizationsUUIDInvitations201ApplicationJSON
	PostOrganizationsUUIDInvitations403ApplicationJSONObject *PostOrganizationsUUIDInvitations403ApplicationJSON
}
