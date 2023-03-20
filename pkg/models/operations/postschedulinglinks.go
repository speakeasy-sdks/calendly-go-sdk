package operations

import (
	"net/http"
)

type PostSchedulingLinksRequestBodyOwnerTypeEnum string

const (
	PostSchedulingLinksRequestBodyOwnerTypeEnumEventType PostSchedulingLinksRequestBodyOwnerTypeEnum = "EventType"
)

type PostSchedulingLinksRequestBody struct {
	MaxEventCount float64                                     `json:"max_event_count"`
	Owner         string                                      `json:"owner"`
	OwnerType     PostSchedulingLinksRequestBodyOwnerTypeEnum `json:"owner_type"`
}

type PostSchedulingLinksRequest struct {
	Request PostSchedulingLinksRequestBody `request:"mediaType=application/json"`
}

type PostSchedulingLinksErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// PostSchedulingLinksErrorResponse
// Error Object
type PostSchedulingLinksErrorResponse struct {
	Details []PostSchedulingLinksErrorResponseDetails `json:"details,omitempty"`
	Message string                                    `json:"message"`
	Title   string                                    `json:"title"`
}

type PostSchedulingLinks201ApplicationJSONResourceOwnerTypeEnum string

const (
	PostSchedulingLinks201ApplicationJSONResourceOwnerTypeEnumEventType PostSchedulingLinks201ApplicationJSONResourceOwnerTypeEnum = "EventType"
)

type PostSchedulingLinks201ApplicationJSONResource struct {
	BookingURL string                                                      `json:"booking_url"`
	Owner      *string                                                     `json:"owner,omitempty"`
	OwnerType  *PostSchedulingLinks201ApplicationJSONResourceOwnerTypeEnum `json:"owner_type,omitempty"`
}

type PostSchedulingLinks201ApplicationJSON struct {
	Resource PostSchedulingLinks201ApplicationJSONResource `json:"resource"`
}

type PostSchedulingLinksResponse struct {
	ContentType                                 string
	ErrorResponse                               *PostSchedulingLinksErrorResponse
	StatusCode                                  int
	RawResponse                                 *http.Response
	PostSchedulingLinks201ApplicationJSONObject *PostSchedulingLinks201ApplicationJSON
}
