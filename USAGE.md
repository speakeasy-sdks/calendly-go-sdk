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
        WithSecurity(        shared.Security{
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
                "vero",
                "perspiciatis",
                "nulla",
            },
            Count: 423655,
            MaxOccurredAt: "2022-07-31T10:58:38.487Z",
            MinOccurredAt: "2022-07-23T07:21:46.425Z",
            Namespace: []string{
                "iusto",
                "ullam",
            },
            Organization: "saepe",
            PageToken: "inventore",
            SearchTerm: "sapiente",
            Sort: []ActivityLogSortEnum{
                "actor.display_name:desc",
                "actor.uri:asc",
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