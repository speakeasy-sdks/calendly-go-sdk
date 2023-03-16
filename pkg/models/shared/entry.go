package shared

import (
	"time"
)

// Entry
// Object for a created activity log record
type Entry struct {
	Action             string                 `json:"action"`
	Actor              *Actor                 `json:"actor,omitempty"`
	Details            map[string]interface{} `json:"details"`
	FullyQualifiedName string                 `json:"fully_qualified_name"`
	Namespace          string                 `json:"namespace"`
	OccurredAt         time.Time              `json:"occurred_at"`
	Organization       string                 `json:"organization"`
	URI                string                 `json:"uri"`
}
