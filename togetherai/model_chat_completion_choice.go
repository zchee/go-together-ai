// Code generated by github.com/zchee/go-openapi-tools/cmd/oapi-generator. DO NOT EDIT.

package togetherai

// ChatCompletionChoice represents a model of ChatCompletionChoice.
type ChatCompletionChoice struct {
	Delta        map[string]any `json:"delta"`
	FinishReason string         `json:"finish_reason"`
	Index        int32          `json:"index"`
	Logprobs     map[string]any `json:"logprobs,omitempty"`
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (c *ChatCompletionChoice) GetDelta() (ret map[string]any) {
	if c == nil {
		return ret
	}
	return c.Delta
}

// GetFinishReason returns the FinishReason field value if set, zero value otherwise.
func (c *ChatCompletionChoice) GetFinishReason() (ret string) {
	if c == nil {
		return ret
	}
	return c.FinishReason
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (c *ChatCompletionChoice) GetIndex() (ret int32) {
	if c == nil {
		return ret
	}
	return c.Index
}

// GetLogprobs returns the Logprobs field value if set, zero value otherwise.
func (c *ChatCompletionChoice) GetLogprobs() (ret map[string]any) {
	if c == nil {
		return ret
	}
	return c.Logprobs
}
