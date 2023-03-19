package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetInviteeNoShowPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetInviteeNoShowRequest struct {
	PathParams GetInviteeNoShowPathParams
}

type GetInviteeNoShowErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetInviteeNoShowErrorResponse
// Error Object
type GetInviteeNoShowErrorResponse struct {
	Details []GetInviteeNoShowErrorResponseDetails `json:"details,omitempty"`
	Message string                                 `json:"message"`
	Title   string                                 `json:"title"`
}

type GetInviteeNoShow200ApplicationJSON struct {
	Resource shared.InviteeNoShow `json:"resource"`
}

type GetInviteeNoShowResponse struct {
	ContentType                              string
	ErrorResponse                            *GetInviteeNoShowErrorResponse
	StatusCode                               int
	RawResponse                              *http.Response
	GetInviteeNoShow200ApplicationJSONObject *GetInviteeNoShow200ApplicationJSON
}
