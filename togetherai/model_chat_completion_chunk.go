// Code generated by github.com/zchee/go-openapi-tools/cmd/oapi-generator. DO NOT EDIT.

package togetherai

// ChatCompletionChunk represents a model of ChatCompletionChunk.
type ChatCompletionChunk struct {
	Choices           []map[string]any `json:"choices"`
	Created           int32            `json:"created"`
	ID                string           `json:"id"`
	Model             string           `json:"model"`
	Object            string           `json:"object"`
	SystemFingerprint string           `json:"system_fingerprint,omitempty"`
}

// GetChoices returns the Choices field value if set, zero value otherwise.
func (c *ChatCompletionChunk) GetChoices() (ret []map[string]any) {
	if c == nil {
		return ret
	}
	return c.Choices
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (c *ChatCompletionChunk) GetCreated() (ret int32) {
	if c == nil {
		return ret
	}
	return c.Created
}

// GetID returns the ID field value if set, zero value otherwise.
func (c *ChatCompletionChunk) GetID() (ret string) {
	if c == nil {
		return ret
	}
	return c.ID
}

// GetModel returns the Model field value if set, zero value otherwise.
func (c *ChatCompletionChunk) GetModel() (ret string) {
	if c == nil {
		return ret
	}
	return c.Model
}

// GetObject returns the Object field value if set, zero value otherwise.
func (c *ChatCompletionChunk) GetObject() (ret string) {
	if c == nil {
		return ret
	}
	return c.Object
}

// GetSystemFingerprint returns the SystemFingerprint field value if set, zero value otherwise.
func (c *ChatCompletionChunk) GetSystemFingerprint() (ret string) {
	if c == nil {
		return ret
	}
	return c.SystemFingerprint
}
