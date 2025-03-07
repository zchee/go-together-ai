// Code generated by github.com/zchee/go-openapi-tools/cmd/oapi-generator. DO NOT EDIT.

package togetherai

// ModelInfo represents a model of ModelInfo.
type ModelInfo struct {
	ContextLength int32          `json:"context_length,omitempty"`
	Created       int32          `json:"created"`
	DisplayName   string         `json:"display_name,omitempty"`
	ID            string         `json:"id"`
	License       string         `json:"license,omitempty"`
	Link          string         `json:"link,omitempty"`
	Object        string         `json:"object"`
	Organization  string         `json:"organization,omitempty"`
	Pricing       map[string]any `json:"pricing,omitempty"`
}

// GetContextLength returns the ContextLength field value if set, zero value otherwise.
func (m *ModelInfo) GetContextLength() (ret int32) {
	if m == nil {
		return ret
	}
	return m.ContextLength
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (m *ModelInfo) GetCreated() (ret int32) {
	if m == nil {
		return ret
	}
	return m.Created
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (m *ModelInfo) GetDisplayName() (ret string) {
	if m == nil {
		return ret
	}
	return m.DisplayName
}

// GetID returns the ID field value if set, zero value otherwise.
func (m *ModelInfo) GetID() (ret string) {
	if m == nil {
		return ret
	}
	return m.ID
}

// GetLicense returns the License field value if set, zero value otherwise.
func (m *ModelInfo) GetLicense() (ret string) {
	if m == nil {
		return ret
	}
	return m.License
}

// GetLink returns the Link field value if set, zero value otherwise.
func (m *ModelInfo) GetLink() (ret string) {
	if m == nil {
		return ret
	}
	return m.Link
}

// GetObject returns the Object field value if set, zero value otherwise.
func (m *ModelInfo) GetObject() (ret string) {
	if m == nil {
		return ret
	}
	return m.Object
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (m *ModelInfo) GetOrganization() (ret string) {
	if m == nil {
		return ret
	}
	return m.Organization
}

// GetPricing returns the Pricing field value if set, zero value otherwise.
func (m *ModelInfo) GetPricing() (ret map[string]any) {
	if m == nil {
		return ret
	}
	return m.Pricing
}
