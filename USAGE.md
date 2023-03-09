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
    s := sdk.New()
    
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
            MaxOccurredAt: "2022-07-24T22:01:41.287Z",
            MinOccurredAt: "2022-07-16T18:24:49.225Z",
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