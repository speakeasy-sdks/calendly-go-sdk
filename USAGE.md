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
            Oauth2: &shared.SchemeOauth2{
                Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
            },
        }),
    )

    req := operations.ActivityLogRequest{
        QueryParams: operations.ActivityLogQueryParams{
            Action: []string{
                "deserunt",
                "porro",
                "nulla",
            },
            Actor: []string{
                "https://api.calendly.com/users/EBHAAFHDCAEQTSEZ",
                "https://api.calendly.com/users/EBHAAFHDCAEQTSEZ",
                "https://api.calendly.com/users/EBHAAFHDCAEQTSEZ",
            },
            Count: 857946,
            MaxOccurredAt: "2022-09-07T04:05:52.921Z",
            MinOccurredAt: "2022-05-19T19:20:58.319Z",
            Namespace: []string{
                "fuga",
                "facilis",
            },
            Organization: "https://api.calendly.com/organizations/EBHAAFHDCAEQTSEZ",
            PageToken: "eum",
            SearchTerm: "iusto",
            Sort: []ActivityLogSortEnum{
                "occurred_at:asc",
                "action:asc",
            },
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