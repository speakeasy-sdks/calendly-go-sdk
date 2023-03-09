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
    s := sdk.New(sdk.WithSecurity(
        shared.Security{
            Oauth2: &shared.SchemeOauth2{
                Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
            },
        },
    ))
    
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
            MaxOccurredAt: "2022-07-24T22:46:31.860Z",
            MinOccurredAt: "2022-07-16T19:09:39.798Z",
            Namespace: []string{
                "iusto",
                "ullam",
            },
            Organization: "saepe",
            PageToken: "inventore",
            SearchTerm: "sapiente",
            Sort: []ActivityLogSortEnum{
                "actor.uri:asc",
                "actor.uri:desc",
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