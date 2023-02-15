package shared

type InvalidParamPayload struct {
	Invalid  []InvalidParam `json:"invalid"`
	NotFound []string       `json:"notFound"`
	TraceID  string         `json:"traceId"`
}
