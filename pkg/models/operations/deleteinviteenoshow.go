package operations

import (
	"net/http"
)

type DeleteInviteeNoShowPathParams struct {
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type DeleteInviteeNoShowRequest struct {
	PathParams DeleteInviteeNoShowPathParams
}

type DeleteInviteeNoShowErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// DeleteInviteeNoShowErrorResponse
// Error Object
type DeleteInviteeNoShowErrorResponse struct {
	Details []DeleteInviteeNoShowErrorResponseDetails `json:"details,omitempty"`
	Message string                                    `json:"message"`
	Title   string                                    `json:"title"`
}

type DeleteInviteeNoShowResponse struct {
	ContentType   string
	ErrorResponse *DeleteInviteeNoShowErrorResponse
	StatusCode    int
	RawResponse   *http.Response
}
