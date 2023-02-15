package shared

type ProjectRunsResponsePayload struct {
	NextPage     map[string]interface{}         `json:"nextPage"`
	PreviousPage map[string]interface{}         `json:"previousPage"`
	Runs         []ProjectStatusResponsePayload `json:"runs"`
	TraceID      string                         `json:"traceId"`
}
