package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type PostUsersUUIDWebhooksRequestBodyEventsEnum string

const (
	PostUsersUUIDWebhooksRequestBodyEventsEnumInviteeCanceled              PostUsersUUIDWebhooksRequestBodyEventsEnum = "invitee.canceled"
	PostUsersUUIDWebhooksRequestBodyEventsEnumInviteeCreated               PostUsersUUIDWebhooksRequestBodyEventsEnum = "invitee.created"
	PostUsersUUIDWebhooksRequestBodyEventsEnumRoutingFormSubmissionCreated PostUsersUUIDWebhooksRequestBodyEventsEnum = "routing_form_submission.created"
)

type PostUsersUUIDWebhooksRequestBodyScopeEnum string

const (
	PostUsersUUIDWebhooksRequestBodyScopeEnumOrganization PostUsersUUIDWebhooksRequestBodyScopeEnum = "organization"
	PostUsersUUIDWebhooksRequestBodyScopeEnumUser         PostUsersUUIDWebhooksRequestBodyScopeEnum = "user"
)

type PostUsersUUIDWebhooksRequestBody struct {
	Events       []PostUsersUUIDWebhooksRequestBodyEventsEnum `json:"events"`
	Organization string                                       `json:"organization"`
	Scope        PostUsersUUIDWebhooksRequestBodyScopeEnum    `json:"scope"`
	SigningKey   *string                                      `json:"signing_key,omitempty"`
	URL          string                                       `json:"url"`
	User         *string                                      `json:"user,omitempty"`
}

type PostUsersUUIDWebhooksRequest struct {
	Request PostUsersUUIDWebhooksRequestBody `request:"mediaType=application/json"`
}

type PostUsersUUIDWebhooksErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostUsersUUIDWebhooksErrorResponse
// Error Object
type PostUsersUUIDWebhooksErrorResponse struct {
	Details []PostUsersUUIDWebhooksErrorResponseDetails `json:"details,omitempty"`
	Message string                                      `json:"message"`
	Title   string                                      `json:"title"`
}

type PostUsersUUIDWebhooks403ApplicationJSONMessageEnum string

const (
	PostUsersUUIDWebhooks403ApplicationJSONMessageEnumPleaseUpgradeYourCalendlyAccountToProfessional PostUsersUUIDWebhooks403ApplicationJSONMessageEnum = "Please upgrade your Calendly account to Professional"
	PostUsersUUIDWebhooks403ApplicationJSONMessageEnumYouDoNotHavePermissionToAccessThisResource     PostUsersUUIDWebhooks403ApplicationJSONMessageEnum = "You do not have permission to access this resource."
	PostUsersUUIDWebhooks403ApplicationJSONMessageEnumYouDoNotHavePermission                         PostUsersUUIDWebhooks403ApplicationJSONMessageEnum = "You do not have permission"
)

type PostUsersUUIDWebhooks403ApplicationJSONTitleEnum string

const (
	PostUsersUUIDWebhooks403ApplicationJSONTitleEnumPermissionDenied PostUsersUUIDWebhooks403ApplicationJSONTitleEnum = "Permission Denied"
)

type PostUsersUUIDWebhooks403ApplicationJSON struct {
	Message *PostUsersUUIDWebhooks403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *PostUsersUUIDWebhooks403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type PostUsersUUIDWebhooks201ApplicationJSON struct {
	Resource shared.WebhookSubscription `json:"resource"`
}

type PostUsersUUIDWebhooksResponse struct {
	ContentType                                   string
	ErrorResponse                                 *PostUsersUUIDWebhooksErrorResponse
	StatusCode                                    int
	RawResponse                                   *http.Response
	PostUsersUUIDWebhooks201ApplicationJSONObject *PostUsersUUIDWebhooks201ApplicationJSON
	PostUsersUUIDWebhooks403ApplicationJSONObject *PostUsersUUIDWebhooks403ApplicationJSON
}
