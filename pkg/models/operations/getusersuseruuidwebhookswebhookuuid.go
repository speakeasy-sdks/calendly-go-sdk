package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetUsersUserUUIDWebhooksWebhookUUIDSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type GetUsersUserUUIDWebhooksWebhookUUIDPathParams struct {
	WebhookUUID string `pathParam:"style=simple,explode=false,name=webhook_uuid"`
}

type GetUsersUserUUIDWebhooksWebhookUUIDRequest struct {
	PathParams GetUsersUserUUIDWebhooksWebhookUUIDPathParams
	Security   GetUsersUserUUIDWebhooksWebhookUUIDSecurity
}

type GetUsersUserUUIDWebhooksWebhookUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetUsersUserUUIDWebhooksWebhookUUIDErrorResponse
// Error Object
type GetUsersUserUUIDWebhooksWebhookUUIDErrorResponse struct {
	Details []GetUsersUserUUIDWebhooksWebhookUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                                    `json:"message"`
	Title   string                                                    `json:"title"`
}

type GetUsersUserUUIDWebhooksWebhookUUID200ApplicationJSON struct {
	Resource shared.WebhookSubscription `json:"resource"`
}

type GetUsersUserUUIDWebhooksWebhookUUIDResponse struct {
	ContentType                                                 string
	ErrorResponse                                               *GetUsersUserUUIDWebhooksWebhookUUIDErrorResponse
	StatusCode                                                  int
	RawResponse                                                 *http.Response
	GetUsersUserUUIDWebhooksWebhookUUID200ApplicationJSONObject *GetUsersUserUUIDWebhooksWebhookUUID200ApplicationJSON
}
