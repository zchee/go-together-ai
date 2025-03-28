// Code generated by github.com/zchee/go-openapi-tools/cmd/oapi-generator. DO NOT EDIT.

package togetherai

// ToolChoice represents a model of ToolChoice.
type ToolChoice struct {
	Function map[string]any `json:"function"`
	ID       string         `json:"id"`
	Index    float32        `json:"index"`
	Type     string         `json:"type"`
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (t *ToolChoice) GetFunction() (ret map[string]any) {
	if t == nil {
		return ret
	}
	return t.Function
}

// GetID returns the ID field value if set, zero value otherwise.
func (t *ToolChoice) GetID() (ret string) {
	if t == nil {
		return ret
	}
	return t.ID
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (t *ToolChoice) GetIndex() (ret float32) {
	if t == nil {
		return ret
	}
	return t.Index
}

// GetType returns the Type field value if set, zero value otherwise.
func (t *ToolChoice) GetType() (ret string) {
	if t == nil {
		return ret
	}
	return t.Type
}
