package shared

import (
	"time"
)

type RoutingFormStatusEnum string

const (
	RoutingFormStatusEnumDraft     RoutingFormStatusEnum = "draft"
	RoutingFormStatusEnumPublished RoutingFormStatusEnum = "published"
)

// RoutingForm
// Information about a routing form.
type RoutingForm struct {
	CreatedAt    time.Time             `json:"created_at"`
	Name         string                `json:"name"`
	Organization string                `json:"organization"`
	Questions    []Question            `json:"questions"`
	Status       RoutingFormStatusEnum `json:"status"`
	UpdatedAt    time.Time             `json:"updated_at"`
	URI          string                `json:"uri"`
}
