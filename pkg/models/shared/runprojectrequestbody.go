package shared

// RunProjectRequestBody
// Request body format for triggering a project run
type RunProjectRequestBody struct {
	DryRun      *bool                  `json:"dryRun,omitempty"`
	InputParams map[string]interface{} `json:"inputParams,omitempty"`
	UpdateCache *bool                  `json:"updateCache,omitempty"`
}
