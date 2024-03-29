// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
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

func (e *ActivityLogSortEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "action:asc":
		fallthrough
	case "action:desc":
		fallthrough
	case "actor.display_name:asc":
		fallthrough
	case "actor.display_name:desc":
		fallthrough
	case "actor.uri:asc":
		fallthrough
	case "actor.uri:desc":
		fallthrough
	case "namespace:asc":
		fallthrough
	case "namespace:desc":
		fallthrough
	case "occurred_at:asc":
		fallthrough
	case "occurred_at:desc":
		*e = ActivityLogSortEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ActivityLogSortEnum: %s", s)
	}
}

type ActivityLogRequest struct {
	// The action(s) associated with the entries
	Action []string `queryParam:"style=form,explode=false,name=action"`
	// Return entries from the user(s) associated with the provided URIs
	Actor []string `queryParam:"style=form,explode=false,name=actor"`
	// The number of rows to return
	Count *int64 `queryParam:"style=form,explode=true,name=count"`
	// Include entries that occurred prior to this time (sample time format: "2020-01-02T03:04:05.678Z"). This time should use the UTC timezone.
	MaxOccurredAt *time.Time `queryParam:"style=form,explode=true,name=max_occurred_at"`
	// Include entries that occurred after this time (sample time format: "2020-01-02T03:04:05.678Z"). This time should use the UTC timezone.
	MinOccurredAt *time.Time `queryParam:"style=form,explode=true,name=min_occurred_at"`
	// The categories of the entries
	Namespace []string `queryParam:"style=form,explode=false,name=namespace"`
	// Return activity log entries from the organization associated with this URI
	Organization string `queryParam:"style=form,explode=true,name=organization"`
	// The token to pass to get the next portion of the collection
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
	// Filters entries based on the search term.
	//
	// Supported operators:
	//   - `|` - to allow filtering by one term or another. Example: `this | that`
	//   - `+` - to allow filtering by one term and another. Example: `this + that`
	//   - `"` - to allow filtering by an exact search term. Example: `"email@website.com"`
	//   - `-` - to omit specific terms from results. Example: `Added -User`
	//   - `()` - to allow specifying precedence during a search. Example: `(this + that) OR (person + place)`
	//   - `*` - to allow prefix searching. Example `*@other-website.com`
	//
	SearchTerm *string `queryParam:"style=form,explode=true,name=search_term"`
	// Order results by the specified field and direction. List of {field}:{direction} values.
	Sort []ActivityLogSortEnum `queryParam:"style=form,explode=false,name=sort"`
}

type ActivityLog403ApplicationJSONMessageEnum string

const (
	ActivityLog403ApplicationJSONMessageEnumPleaseUpgradeYourCalendlyAccountToEnterprise ActivityLog403ApplicationJSONMessageEnum = "Please upgrade your Calendly account to Enterprise."
	ActivityLog403ApplicationJSONMessageEnumYouDoNotHavePermissionToAccessThisResource   ActivityLog403ApplicationJSONMessageEnum = "You do not have permission to access this resource."
)

func (e *ActivityLog403ApplicationJSONMessageEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Please upgrade your Calendly account to Enterprise.":
		fallthrough
	case "You do not have permission to access this resource.":
		*e = ActivityLog403ApplicationJSONMessageEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ActivityLog403ApplicationJSONMessageEnum: %s", s)
	}
}

type ActivityLog403ApplicationJSONTitleEnum string

const (
	ActivityLog403ApplicationJSONTitleEnumPermissionDenied ActivityLog403ApplicationJSONTitleEnum = "Permission Denied"
)

func (e *ActivityLog403ApplicationJSONTitleEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Permission Denied":
		*e = ActivityLog403ApplicationJSONTitleEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ActivityLog403ApplicationJSONTitleEnum: %s", s)
	}
}

// ActivityLog403ApplicationJSON - Permission Denied
type ActivityLog403ApplicationJSON struct {
	Message *ActivityLog403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *ActivityLog403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type ActivityLogErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// ActivityLogErrorResponse - Error Object
type ActivityLogErrorResponse struct {
	Details []ActivityLogErrorResponseDetails `json:"details,omitempty"`
	Message string                            `json:"message"`
	Title   string                            `json:"title"`
}

// ActivityLog200ApplicationJSON - Activity log response
type ActivityLog200ApplicationJSON struct {
	// The set of activity log entries matching the criteria
	Collection []shared.Entry `json:"collection"`
	// If there are more search results than the total_count field indicates, pagination will continue to return results past the total_count field value.
	ExceedsMaxTotalCount bool `json:"exceeds_max_total_count"`
	// The date and time of the newest entry (format: "2020-01-02T03:04:05.678Z") in the collection array.
	LastEventTime *time.Time        `json:"last_event_time,omitempty"`
	Pagination    shared.Pagination `json:"pagination"`
	// Total number of records based on search criteria
	TotalCount int64 `json:"total_count"`
}

type ActivityLogResponse struct {
	ContentType string
	// Request is not valid
	ErrorResponse *ActivityLogErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// OK
	ActivityLog200ApplicationJSONObject *ActivityLog200ApplicationJSON
	// Permission Denied
	ActivityLog403ApplicationJSONObject *ActivityLog403ApplicationJSON
}
