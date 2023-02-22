package shared

// ProjectRunResponsePayload
// Response format returned by the runProject endpoint
type ProjectRunResponsePayload struct {
	ProjectID      string `json:"projectId"`
	ProjectVersion int    `json:"projectVersion"`
	RunID          string `json:"runId"`
	RunStatusURL   string `json:"runStatusUrl"`
	RunURL         string `json:"runUrl"`
	TraceID        string `json:"traceId"`
}
