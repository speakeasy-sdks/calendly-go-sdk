package operations

import (
	"net/http"
)

type DeleteUsersUserUUIDWebhooksWebhookUUIDPathParams struct {
	WebhookUUID string `pathParam:"style=simple,explode=false,name=webhook_uuid"`
}

type DeleteUsersUserUUIDWebhooksWebhookUUIDRequest struct {
	PathParams DeleteUsersUserUUIDWebhooksWebhookUUIDPathParams
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
