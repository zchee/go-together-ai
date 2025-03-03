// Code generated by github.com/zchee/go-openapi-tools/cmd/oapi-generator. DO NOT EDIT.

package togetherai

// CompletionChunk represents a model of CompletionChunk.
type CompletionChunk struct {
	ID    string         `json:"id"`
	Seed  int32          `json:"seed,omitempty"`
	Token map[string]any `json:"token"`
}

// GetID returns the ID field value if set, zero value otherwise.
func (c *CompletionChunk) GetID() (ret string) {
	if c == nil {
		return ret
	}
	return c.ID
}

// GetSeed returns the Seed field value if set, zero value otherwise.
func (c *CompletionChunk) GetSeed() (ret int32) {
	if c == nil {
		return ret
	}
	return c.Seed
}

// GetToken returns the Token field value if set, zero value otherwise.
func (c *CompletionChunk) GetToken() (ret map[string]any) {
	if c == nil {
		return ret
	}
	return c.Token
}
