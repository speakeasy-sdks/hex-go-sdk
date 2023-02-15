package shared

type TsoaErrorResponsePayload struct {
	Details *string `json:"details,omitempty"`
	Reason  string  `json:"reason"`
	TraceID string  `json:"traceId"`
}
