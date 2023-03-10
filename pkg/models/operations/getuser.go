package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetUserPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type GetUserRequest struct {
	PathParams GetUserPathParams
}

type GetUserErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetUserErrorResponse
// Error Object
type GetUserErrorResponse struct {
	Details []GetUserErrorResponseDetails `json:"details,omitempty"`
	Message string                        `json:"message"`
	Title   string                        `json:"title"`
}

// GetUser200ApplicationJSON
// Service response
type GetUser200ApplicationJSON struct {
	Resource shared.User `json:"resource"`
}

type GetUserResponse struct {
	ContentType                     string
	ErrorResponse                   *GetUserErrorResponse
	StatusCode                      int
	RawResponse                     *http.Response
	GetUser200ApplicationJSONObject *GetUser200ApplicationJSON
}
