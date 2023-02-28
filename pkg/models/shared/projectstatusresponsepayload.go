package shared

import (
	"time"
)

// ProjectStatusResponsePayload
// Response format returned by the getRunStatus endpoint
type ProjectStatusResponsePayload struct {
	ElapsedTime    float64              `json:"elapsedTime"`
	EndTime        time.Time            `json:"endTime"`
	ProjectID      string               `json:"projectId"`
	ProjectVersion int                  `json:"projectVersion"`
	RunID          string               `json:"runId"`
	RunURL         string               `json:"runUrl"`
	StartTime      time.Time            `json:"startTime"`
	Status         ProjectRunStatusEnum `json:"status"`
	TraceID        string               `json:"traceId"`
}
