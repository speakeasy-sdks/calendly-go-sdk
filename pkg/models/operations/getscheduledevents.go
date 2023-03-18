package operations

import (
	"github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
	"net/http"
)

type GetScheduledEventsStatusEnum string

const (
	GetScheduledEventsStatusEnumActive   GetScheduledEventsStatusEnum = "active"
	GetScheduledEventsStatusEnumCanceled GetScheduledEventsStatusEnum = "canceled"
)

type GetScheduledEventsQueryParams struct {
	Count        *float64                      `queryParam:"style=form,explode=true,name=count"`
	InviteeEmail *string                       `queryParam:"style=form,explode=true,name=invitee_email"`
	MaxStartTime *string                       `queryParam:"style=form,explode=true,name=max_start_time"`
	MinStartTime *string                       `queryParam:"style=form,explode=true,name=min_start_time"`
	Organization *string                       `queryParam:"style=form,explode=true,name=organization"`
	PageToken    *string                       `queryParam:"style=form,explode=true,name=page_token"`
	Sort         *string                       `queryParam:"style=form,explode=true,name=sort"`
	Status       *GetScheduledEventsStatusEnum `queryParam:"style=form,explode=true,name=status"`
	User         *string                       `queryParam:"style=form,explode=true,name=user"`
}

type GetScheduledEventsRequest struct {
	QueryParams GetScheduledEventsQueryParams
}

type GetScheduledEvents403ApplicationJSONMessageEnum string

const (
	GetScheduledEvents403ApplicationJSONMessageEnumYouDoNotHavePermissionToAccessThisResource                                      GetScheduledEvents403ApplicationJSONMessageEnum = "You do not have permission to access this resource."
	GetScheduledEvents403ApplicationJSONMessageEnumPleaseAlsoSpecifyOrganizationWhenRequestingEventsForAUserWithinYourOrganization GetScheduledEvents403ApplicationJSONMessageEnum = "Please also specify organization when requesting events for a user within your organization."
	GetScheduledEvents403ApplicationJSONMessageEnumThisUserIsNotInYourOrganization                                                 GetScheduledEvents403ApplicationJSONMessageEnum = "This user is not in your organization"
	GetScheduledEvents403ApplicationJSONMessageEnumYouDoNotHavePermission                                                          GetScheduledEvents403ApplicationJSONMessageEnum = "You do not have permission"
)

type GetScheduledEvents403ApplicationJSONTitleEnum string

const (
	GetScheduledEvents403ApplicationJSONTitleEnumPermissionDenied GetScheduledEvents403ApplicationJSONTitleEnum = "Permission Denied"
)

type GetScheduledEvents403ApplicationJSON struct {
	Message *GetScheduledEvents403ApplicationJSONMessageEnum `json:"message,omitempty"`
	Title   *GetScheduledEvents403ApplicationJSONTitleEnum   `json:"title,omitempty"`
}

type GetScheduledEventsErrorResponseDetails struct {
	Message   string  `json:"message"`
	Parameter *string `json:"parameter,omitempty"`
}

// GetScheduledEventsErrorResponse
// Error Object
type GetScheduledEventsErrorResponse struct {
	Details []GetScheduledEventsErrorResponseDetails `json:"details,omitempty"`
	Message string                                   `json:"message"`
	Title   string                                   `json:"title"`
}

// GetScheduledEvents200ApplicationJSON
// Service response
type GetScheduledEvents200ApplicationJSON struct {
	Collection []shared.Event    `json:"collection"`
	Pagination shared.Pagination `json:"pagination"`
}

type GetScheduledEventsResponse struct {
	ContentType                                string
	ErrorResponse                              *GetScheduledEventsErrorResponse
	StatusCode                                 int
	RawResponse                                *http.Response
	GetScheduledEvents200ApplicationJSONObject *GetScheduledEvents200ApplicationJSON
	GetScheduledEvents403ApplicationJSONObject *GetScheduledEvents403ApplicationJSON
}
