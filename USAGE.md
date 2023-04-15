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