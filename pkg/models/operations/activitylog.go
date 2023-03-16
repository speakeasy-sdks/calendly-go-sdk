package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
	"time"
)

type ActivityLogSortEnum string

const (
	ActivityLogSortEnumActionAsc            ActivityLogSortEnum = "action:asc"
	ActivityLogSortEnumActionDesc           ActivityLogSortEnum = "action:desc"
	ActivityLogSortEnumActorDisplayNameAsc  ActivityLogSortEnum = "actor.display_name:asc"
	ActivityLogSortEnumActorDisplayNameDesc ActivityLogSortEnum = "actor.display_name:desc"
	ActivityLogSortEnumActorURIAsc          ActivityLogSortEnum = "actor.uri:asc"
	ActivityLogSortEnumActorURIDesc         ActivityLogSortEnum = "actor.uri:desc"
	ActivityLogSortEnumNamespaceAsc         ActivityLogSortEnum = "namespace:asc"
	ActivityLogSortEnumNamespaceDesc        ActivityLogSortEnum = "namespace:desc"
	ActivityLogSortEnumOccurredAtAsc        ActivityLogSortEnum = "occurred_at:asc"
	ActivityLogSortEnumOccurredAtDesc       ActivityLogSortEnum = "occurred_at:desc"
)

type ActivityLogQueryParams struct {
	Action        []string              `queryParam:"style=form,explode=false,name=action"`
	Actor         []string              `queryParam:"style=form,explode=false,name=actor"`
	Count         *int64                `queryParam:"style=form,explode=true,name=count"`
	MaxOccurredAt *time.Time            `queryParam:"style=form,explode=true,name=max_occurred_at"`
	MinOccurredAt *time.Time            `queryParam:"style=form,explode=true,name=min_occurred_at"`
	Namespace     []string              `queryParam:"style=form,explode=false,name=namespace"`
	Organization  string                `queryParam:"style=form,explode=true,name=organization"`
	PageToken     *string               `queryParam:"style=form,explode=true,name=page_token"`
	SearchTerm    *string               `queryParam:"style=form,explode=true,name=search_term"`
	Sort          []ActivityLogSortEnum `queryParam:"style=form,explode=false,name=sort"`
}

type ActivityLogRequest struct {
	QueryParams ActivityLogQueryParams
}

type ActivityLog403ApplicationJSONMessageEnum string

const (
	ActivityLog403ApplicationJSONMessageEnumPleaseUpgradeYourCalendlyAccountToEnterprise ActivityLog403ApplicationJSONMessageEnum = "Please upgrade your Calendly account to Enterprise."
	ActivityLog403ApplicationJSONMessageEnumYouDoNotHavePermissionToAccessThisResource   ActivityLog403ApplicationJSONMessageEnum = "You do not have permission to access this resource."
)

type ActivityLog403ApplicationJSONTitleEnum string

const (
	ActivityLog403ApplicationJSONTitleEnumPermissionDenied ActivityLog403ApplicationJSONTitleEnum = "Permission Denied"
)

type ActivityLog403ApplicationJSON struct {
	Message *ActivityLog403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *ActivityLog403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type ActivityLogErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// ActivityLogErrorResponse
// Error Object
type ActivityLogErrorResponse struct {
	Details []ActivityLogErrorResponseDetails `json:"details,omitempty"`
	Message string                            `json:"message"`
	Title   string                            `json:"title"`
}

// ActivityLog200ApplicationJSON
// Activity log response
type ActivityLog200ApplicationJSON struct {
	Collection           []shared.Entry    `json:"collection"`
	ExceedsMaxTotalCount bool              `json:"exceeds_max_total_count"`
	LastEventTime        *time.Time        `json:"last_event_time,omitempty"`
	Pagination           shared.Pagination `json:"pagination"`
	TotalCount           int64             `json:"total_count"`
}

type ActivityLogResponse struct {
	ContentType                         string
	ErrorResponse                       *ActivityLogErrorResponse
	StatusCode                          int
	RawResponse                         *http.Response
	ActivityLog200ApplicationJSONObject *ActivityLog200ApplicationJSON
	ActivityLog403ApplicationJSONObject *ActivityLog403ApplicationJSON
}
