package operations

import (
	"github.com/speakeasy-sdks/hex-go-sdk/pkg/models/shared"
)

type CancelRunPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
	RunID     string `pathParam:"style=simple,explode=false,name=runId"`
}

type CancelRunRequest struct {
	PathParams CancelRunPathParams
}

type CancelRunResponse struct {
	ContentType              string
	StatusCode               int
	TsoaErrorResponsePayload *shared.TsoaErrorResponsePayload
}
