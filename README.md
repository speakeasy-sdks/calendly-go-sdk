# github.com/speakeasy-sdks/calendly-go-sdk

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/calendly-go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/calendly-go-sdk"
    "github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/calendly-go-sdk/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            Oauth2: sdk.String("Bearer YOUR_ACCESS_TOKEN_HERE"),
        }),
    )

    req := operations.ActivityLogRequest{
        Action: []string{
            "provident",
            "distinctio",
            "quibusdam",
        },
        Actor: []string{
            "https://api.calendly.com/users/EBHAAFHDCAEQTSEZ",
            "https://api.calendly.com/users/EBHAAFHDCAEQTSEZ",
            "https://api.calendly.com/users/EBHAAFHDCAEQTSEZ",
        },
        Count: 857946,
        MaxOccurredAt: "2021-04-22T12:08:58.275Z",
        MinOccurredAt: "2022-05-18T09:34:54.894Z",
        Namespace: []string{
            "suscipit",
            "iure",
            "magnam",
        },
        Organization: "https://api.calendly.com/organizations/EBHAAFHDCAEQTSEZ",
        PageToken: "debitis",
        SearchTerm: "ipsa",
        Sort: []ActivityLogSortEnum{
            "actor.display_name:asc",
            "actor.display_name:desc",
            "actor.uri:asc",
            "namespace:desc",
        },
    }

    ctx := context.Background()
    res, err := s.ActivityLog.ActivityLog(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ActivityLog200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### ActivityLog

* `ActivityLog` - List activity log entries

### Availability

* `GetUserAvailabilitySchedules` - List User Availability Schedules
* `GetUserAvailabilitySchedulesUUID` - Get User Availability Schedule
* `GetUserBusyTimes` - List User Busy Times

### DataCompliance

* `PostDataComplianceDeletionEvents` - Delete Scheduled Event Data
* `PostDataComplianceDeletionInvitees` - Delete Invitee Data

### EventTypes

* `GetEventTypesUUID` - Get Event Type
* `GetEventTypeAvailableTimes` - List Event Type Available Times
* `GetEventTypes` - List User's Event Types

### Organizations

* `DeleteOrganizationsUUIDMemberships` - Remove User from Organization
* `GetOrganizationMemberships` - List Organization Memberships
* `GetOrganizationsOrgUUIDInvitationsUUID` - Get Organization Invitation
* `GetOrganizationsUUIDInvitations` - List Organization Invitations
* `GetOrganizationsUUIDMemberships` - Get Organization Membership
* `PostOrganizationsUUIDInvitations` - Invite User to Organization
* `RevokeUsersOrganizationInvitation` - Revoke User's Organization Invitation

### RoutingForms

* `GetRoutingFormSubmissions` - List Routing Form Submissions
* `GetRoutingFormSubmissionsUUID` - Get Routing Form Submission
* `GetRoutingForms` - List Routing Forms
* `GetRoutingFormsUUID` - Get Routing Form

### ScheduledEvents

* `DeleteInviteeNoShow` - Delete Invitee No Show
* `GetScheduledEventsEventUUIDInviteesInviteeUUID` - Get Event Invitee
* `GetScheduledEventsUUID` - Get Event
* `GetInviteeNoShow` - Get Invitee No Show
* `GetInvitees` - List Event Invitees
* `GetScheduledEvents` - List Events
* `PostScheduledEventsUUIDCancellationJSON` - Cancel Event
* `PostScheduledEventsUUIDCancellationMultipart` - Cancel Event
* `PostScheduledEventsUUIDCancellationRaw` - Cancel Event
* `PostInviteeNoShow` - Create Invitee No Show

### SchedulingLinks

* `PostSchedulingLinks` - Create Single-Use Scheduling Link

### Shares

* `PostShares` - Create Share

### Users

* `GetMyUserAccount` - Get current user
* `GetUser` - Get user

### Webhooks

* `DeleteUsersUserUUIDWebhooksWebhookUUID` - Delete Webhook Subscription
* `GetUsersUserUUIDWebhooksWebhookUUID` - Get Webhook Subscription
* `GetWebhooks` - List Webhook Subscriptions
* `PostUsersUUIDWebhooks` - Create Webhook Subscription
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
