package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteUsersUserUUIDWebhooksWebhookUUIDSecurity struct {
	Oauth2              *shared.SchemeOauth2              `security:"scheme,type=oauth2"`
	PersonalAccessToken *shared.SchemePersonalAccessToken `security:"scheme,type=http,subtype=bearer"`
}

type DeleteUsersUserUUIDWebhooksWebhookUUIDPathParams struct {
	WebhookUUID string `pathParam:"style=simple,explode=false,name=webhook_uuid"`
}

type DeleteUsersUserUUIDWebhooksWebhookUUIDRequest struct {
	PathParams DeleteUsersUserUUIDWebhooksWebhookUUIDPathParams
	Security   DeleteUsersUserUUIDWebhooksWebhookUUIDSecurity
}

type DeleteUsersUserUUIDWebhooksWebhookUUIDErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// DeleteUsersUserUUIDWebhooksWebhookUUIDErrorResponse
// Error Object
type DeleteUsersUserUUIDWebhooksWebhookUUIDErrorResponse struct {
	Details []DeleteUsersUserUUIDWebhooksWebhookUUIDErrorResponseDetails `json:"details,omitempty"`
	Message string                                                       `json:"message"`
	Title   string                                                       `json:"title"`
}

type DeleteUsersUserUUIDWebhooksWebhookUUIDResponse struct {
	ContentType   string
	ErrorResponse *DeleteUsersUserUUIDWebhooksWebhookUUIDErrorResponse
	StatusCode    int
	RawResponse   *http.Response
}
