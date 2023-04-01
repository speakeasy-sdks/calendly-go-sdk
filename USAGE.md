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