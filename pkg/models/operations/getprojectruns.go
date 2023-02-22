package operations

import (
	"github.com/speakeasy-sdks/hex-go-sdk/pkg/models/shared"
)

type GetProjectRunsPathParams struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
}

type GetProjectRunsQueryParams struct {
	Limit        *int                         `queryParam:"style=form,explode=true,name=limit"`
	Offset       *int                         `queryParam:"style=form,explode=true,name=offset"`
	StatusFilter *shared.ProjectRunStatusEnum `queryParam:"style=form,explode=true,name=statusFilter"`
}

type GetProjectRunsRequest struct {
	PathParams  GetProjectRunsPathParams
	QueryParams GetProjectRunsQueryParams
}

type GetProjectRunsResponse struct {
	ContentType                string
	ProjectRunsResponsePayload *shared.ProjectRunsResponsePayload
	StatusCode                 int
	TsoaErrorResponsePayload   *shared.TsoaErrorResponsePayload
}
