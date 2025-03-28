// Code generated by github.com/zchee/go-openapi-tools/cmd/oapi-generator. DO NOT EDIT.

package togetherai

// CreateEndpointRequest represents a model of CreateEndpointRequest.
type CreateEndpointRequest struct {
	Autoscaling                map[string]any `json:"autoscaling"`
	DisablePromptCache         bool           `json:"disable_prompt_cache,omitempty"`
	DisableSpeculativeDecoding bool           `json:"disable_speculative_decoding,omitempty"`
	DisplayName                string         `json:"display_name,omitempty"`
	Hardware                   string         `json:"hardware"`
	Model                      string         `json:"model"`
	State                      string         `json:"state,omitempty"`
}

// GetAutoscaling returns the Autoscaling field value if set, zero value otherwise.
func (c *CreateEndpointRequest) GetAutoscaling() (ret map[string]any) {
	if c == nil {
		return ret
	}
	return c.Autoscaling
}

// GetDisablePromptCache returns the DisablePromptCache field value if set, zero value otherwise.
func (c *CreateEndpointRequest) GetDisablePromptCache() (ret bool) {
	if c == nil {
		return ret
	}
	return c.DisablePromptCache
}

// GetDisableSpeculativeDecoding returns the DisableSpeculativeDecoding field value if set, zero value otherwise.
func (c *CreateEndpointRequest) GetDisableSpeculativeDecoding() (ret bool) {
	if c == nil {
		return ret
	}
	return c.DisableSpeculativeDecoding
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (c *CreateEndpointRequest) GetDisplayName() (ret string) {
	if c == nil {
		return ret
	}
	return c.DisplayName
}

// GetHardware returns the Hardware field value if set, zero value otherwise.
func (c *CreateEndpointRequest) GetHardware() (ret string) {
	if c == nil {
		return ret
	}
	return c.Hardware
}

// GetModel returns the Model field value if set, zero value otherwise.
func (c *CreateEndpointRequest) GetModel() (ret string) {
	if c == nil {
		return ret
	}
	return c.Model
}

// GetState returns the State field value if set, zero value otherwise.
func (c *CreateEndpointRequest) GetState() (ret string) {
	if c == nil {
		return ret
	}
	return c.State
}
