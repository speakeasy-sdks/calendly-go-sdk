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

    ctx := context.Background()    
    req := operations.ListScheduledEventsRequest{
        Count: 5488.14,
        InviteeEmail: "alice@example.com",
        MaxStartTime: "provident",
        MinStartTime: "distinctio",
        Organization: "https://api.calendly.com/organizations/EBHAAFHDCAEQTSEZ",
        PageToken: "quibusdam",
        Sort: "unde",
        Status: "canceled",
        User: "https://api.calendly.com/users/EBHAAFHDCAEQTSEZ",
    }

    res, err := s.ScheduledEvents.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListScheduledEvents200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### ActivityLog

* `List` - List activity log entries

### Availability

* `Get` - Get User Availability Schedule
* `GetAvailability` - List User Availability Schedules
* `GetBusyTimes` - List User Busy Times

### DataCompliance

* `CreateDeletionEvent` - Delete Scheduled Event Data
* `DeleteInviteeData` - Delete Invitee Data

### EventTypes

* `Get` - Get Event Type
* `GetAvailableTimes` - List Event Type Available Times
* `List` - List User's Event Types

### Organizations

* `DeleteMemberships` - Remove User from Organization
* `GetInvitations` - Get Organization Invitation
* `InviteUser` - Invite User to Organization
* `ListInvitations` - List Organization Invitations
* `ListMemberships` - List Organization Memberships
* `RevokeInvite` - Revoke User's Organization Invitation

### RoutingForms

* `GetSubmissions` - List Routing Form Submissions
* `GetSubmissionsByUUID` - Get Routing Form Submission
* `GetByUUID` - Get Routing Form
* `List` - List Routing Forms

### ScheduledEvents

* `Cancel` - Cancel Event
* `Cancel` - Cancel Event
* `Cancel` - Cancel Event
* `CreateNoShow` - Create Invitee No Show
* `GetEventByUUID` - Get Event
* `GetInvitees` - List Event Invitees
* `GetInviteesByUUID` - Get Event Invitee
* `GetNoShow` - Get Invitee No Show
* `List` - List Events
* `UnmarkNoShow` - Delete Invitee No Show

### SchedulingLinks

* `Create` - Create Single-Use Scheduling Link

### Shares

* `Create` - Create Share

### Users

* `Get` - Get user
* `GetMemberships` - Get Organization Membership
* `Me` - Get current user

### Webhooks

* `Create` - Create Webhook Subscription
* `Delete` - Delete Webhook Subscription
* `Get` - Get Webhook Subscription
* `List` - List Webhook Subscriptions
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
