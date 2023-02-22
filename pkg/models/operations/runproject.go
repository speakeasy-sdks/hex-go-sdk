package operations

import (
	"github.com/speakeasy-sdks/hex-go-sdk/pkg/models/shared"
)

type RunProjectPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
}

type RunProjectRequest struct {
	PathParams RunProjectPathParams
	Request    *shared.RunProjectRequestBody `request:"mediaType=application/json"`
}

type RunProjectResponse struct {
	ContentType                       string
	ProjectRunResponsePayload         *shared.ProjectRunResponsePayload
	RunProject422ApplicationJSONAnyOf *interface{}
	StatusCode                        int
	TsoaErrorResponsePayload          *shared.TsoaErrorResponsePayload
}
