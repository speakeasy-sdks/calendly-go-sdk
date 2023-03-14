package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
	"time"
)

type PostDataComplianceDeletionEventsRequestBody struct {
	EndTime   time.Time `json:"end_time"`
	StartTime time.Time `json:"start_time"`
}

type PostDataComplianceDeletionEventsRequest struct {
	Request PostDataComplianceDeletionEventsRequestBody `request:"mediaType=application/json"`
}

type PostDataComplianceDeletionEventsErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostDataComplianceDeletionEventsErrorResponse
// Error Object
type PostDataComplianceDeletionEventsErrorResponse struct {
	Details []PostDataComplianceDeletionEventsErrorResponseDetails `json:"details,omitempty"`
	Message string                                                 `json:"message"`
	Title   string                                                 `json:"title"`
}

type PostDataComplianceDeletionEventsResponse struct {
	ContentType                                              string
	ErrorResponse                                            *PostDataComplianceDeletionEventsErrorResponse
	ErrorResponse1                                           *shared.ErrorResponse
	StatusCode                                               int
	RawResponse                                              *http.Response
	PostDataComplianceDeletionEvents202ApplicationJSONObject map[string]interface{}
}
