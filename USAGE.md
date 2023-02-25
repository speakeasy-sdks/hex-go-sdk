<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/hex-go-sdk"
    "github.com/speakeasy-sdks/hex-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/hex-go-sdk/pkg/models/operations"
)

func main() {
    opts := []hex.SDKOption{
        hex.WithSecurity(
            shared.Security{
                BearerAuth: shared.SchemeBearerAuth{
                    Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
                },
            },
        ),
    }

    s := hex.New(opts...)
    
    req := operations.CancelRunRequest{
        PathParams: operations.CancelRunPathParams{
            ProjectID: "unde",
            RunID: "deserunt",
        },
    }

    ctx := context.Background()
    res, err := s.CancelRun(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->