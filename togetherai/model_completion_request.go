// Code generated by github.com/zchee/go-openapi-tools/cmd/oapi-generator. DO NOT EDIT.

package togetherai

// CompletionRequest represents a model of CompletionRequest.
type CompletionRequest struct {
	Echo              bool     `json:"echo,omitempty"`
	FrequencyPenalty  float32  `json:"frequency_penalty,omitempty"`
	LogitBias         float32  `json:"logit_bias,omitempty"`
	Logprobs          int32    `json:"logprobs,omitempty"`
	MaxTokens         int32    `json:"max_tokens,omitempty"`
	MinP              float32  `json:"min_p,omitempty"`
	Model             string   `json:"model"`
	PresencePenalty   float32  `json:"presence_penalty,omitempty"`
	Prompt            string   `json:"prompt"`
	RepetitionPenalty float32  `json:"repetition_penalty,omitempty"`
	SafetyModel       string   `json:"safety_model,omitempty"`
	Seed              int32    `json:"seed,omitempty"`
	Stop              []string `json:"stop,omitempty"`
	Stream            bool     `json:"stream,omitempty"`
	Temperature       float32  `json:"temperature,omitempty"`
	TopK              int32    `json:"top_k,omitempty"`
	TopP              float32  `json:"top_p,omitempty"`
}

// GetEcho returns the Echo field value if set, zero value otherwise.
func (c *CompletionRequest) GetEcho() (ret bool) {
	if c == nil {
		return ret
	}
	return c.Echo
}

// GetFrequencyPenalty returns the FrequencyPenalty field value if set, zero value otherwise.
func (c *CompletionRequest) GetFrequencyPenalty() (ret float32) {
	if c == nil {
		return ret
	}
	return c.FrequencyPenalty
}

// GetLogitBias returns the LogitBias field value if set, zero value otherwise.
func (c *CompletionRequest) GetLogitBias() (ret float32) {
	if c == nil {
		return ret
	}
	return c.LogitBias
}

// GetLogprobs returns the Logprobs field value if set, zero value otherwise.
func (c *CompletionRequest) GetLogprobs() (ret int32) {
	if c == nil {
		return ret
	}
	return c.Logprobs
}

// GetMaxTokens returns the MaxTokens field value if set, zero value otherwise.
func (c *CompletionRequest) GetMaxTokens() (ret int32) {
	if c == nil {
		return ret
	}
	return c.MaxTokens
}

// GetMinP returns the MinP field value if set, zero value otherwise.
func (c *CompletionRequest) GetMinP() (ret float32) {
	if c == nil {
		return ret
	}
	return c.MinP
}

// GetModel returns the Model field value if set, zero value otherwise.
func (c *CompletionRequest) GetModel() (ret string) {
	if c == nil {
		return ret
	}
	return c.Model
}

// GetPresencePenalty returns the PresencePenalty field value if set, zero value otherwise.
func (c *CompletionRequest) GetPresencePenalty() (ret float32) {
	if c == nil {
		return ret
	}
	return c.PresencePenalty
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (c *CompletionRequest) GetPrompt() (ret string) {
	if c == nil {
		return ret
	}
	return c.Prompt
}

// GetRepetitionPenalty returns the RepetitionPenalty field value if set, zero value otherwise.
func (c *CompletionRequest) GetRepetitionPenalty() (ret float32) {
	if c == nil {
		return ret
	}
	return c.RepetitionPenalty
}

// GetSafetyModel returns the SafetyModel field value if set, zero value otherwise.
func (c *CompletionRequest) GetSafetyModel() (ret string) {
	if c == nil {
		return ret
	}
	return c.SafetyModel
}

// GetSeed returns the Seed field value if set, zero value otherwise.
func (c *CompletionRequest) GetSeed() (ret int32) {
	if c == nil {
		return ret
	}
	return c.Seed
}

// GetStop returns the Stop field value if set, zero value otherwise.
func (c *CompletionRequest) GetStop() (ret []string) {
	if c == nil {
		return ret
	}
	return c.Stop
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (c *CompletionRequest) GetStream() (ret bool) {
	if c == nil {
		return ret
	}
	return c.Stream
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (c *CompletionRequest) GetTemperature() (ret float32) {
	if c == nil {
		return ret
	}
	return c.Temperature
}

// GetTopK returns the TopK field value if set, zero value otherwise.
func (c *CompletionRequest) GetTopK() (ret int32) {
	if c == nil {
		return ret
	}
	return c.TopK
}

// GetTopP returns the TopP field value if set, zero value otherwise.
func (c *CompletionRequest) GetTopP() (ret float32) {
	if c == nil {
		return ret
	}
	return c.TopP
}
