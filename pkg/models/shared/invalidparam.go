package shared

// InvalidParam
// Information about an input parameter pass to the API call that could not be parsed.
type InvalidParam struct {
	DataType      string `json:"dataType"`
	InputCellType string `json:"inputCellType"`
	ParamName     string `json:"paramName"`
	ParamValue    string `json:"paramValue"`
}
