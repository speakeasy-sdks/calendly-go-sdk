package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetWebhooksScopeEnum string

const (
	GetWebhooksScopeEnumOrganization GetWebhooksScopeEnum = "organization"
	GetWebhooksScopeEnumUser         GetWebhooksScopeEnum = "user"
)

type GetWebhooksQueryParams struct {
	Count        *float64             `queryParam:"style=form,explode=true,name=count"`
	Organization string               `queryParam:"style=form,explode=true,name=organization"`
	PageToken    *string              `queryParam:"style=form,explode=true,name=page_token"`
	Scope        GetWebhooksScopeEnum `queryParam:"style=form,explode=true,name=scope"`
	Sort         *string              `queryParam:"style=form,explode=true,name=sort"`
	User         *string              `queryParam:"style=form,explode=true,name=user"`
}

type GetWebhooksRequest struct {
	QueryParams GetWebhooksQueryParams
}

type GetWebhooks403ApplicationJSONMessageEnum string

const (
	GetWebhooks403ApplicationJSONMessageEnumYouDoNotHavePermission                     GetWebhooks403ApplicationJSONMessageEnum = "You do not have permission"
	GetWebhooks403ApplicationJSONMessageEnumYouDoNotHavePermissionToAccessThisResource GetWebhooks403ApplicationJSONMessageEnum = "You do not have permission to access this resource."
	GetWebhooks403ApplicationJSONMessageEnumUnauthorized                               GetWebhooks403ApplicationJSONMessageEnum = "Unauthorized"
)

type GetWebhooks403ApplicationJSONTitleEnum string

const (
	GetWebhooks403ApplicationJSONTitleEnumPermissionDenied GetWebhooks403ApplicationJSONTitleEnum = "Permission Denied"
)

type GetWebhooks403ApplicationJSON struct {
	Message *GetWebhooks403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *GetWebhooks403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type GetWebhooksErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetWebhooksErrorResponse
// Error Object
type GetWebhooksErrorResponse struct {
	Details []GetWebhooksErrorResponseDetails `json:"details,omitempty"`
	Message string                            `json:"message"`
	Title   string                            `json:"title"`
}

type GetWebhooks200ApplicationJSON struct {
	Collection []shared.WebhookSubscription `json:"collection"`
	Pagination shared.Pagination            `json:"pagination"`
}

type GetWebhooksResponse struct {
	ContentType                         string
	ErrorResponse                       *GetWebhooksErrorResponse
	StatusCode                          int
	RawResponse                         *http.Response
	GetWebhooks200ApplicationJSONObject *GetWebhooks200ApplicationJSON
	GetWebhooks403ApplicationJSONObject *GetWebhooks403ApplicationJSON
}
