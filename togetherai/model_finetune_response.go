// Code generated by github.com/zchee/go-openapi-tools/cmd/oapi-generator. DO NOT EDIT.

package togetherai

// FinetuneResponse represents a model of FinetuneResponse.
type FinetuneResponse struct {
	BatchSize            int32          `json:"batch_size,omitempty"`
	CreatedAt            string         `json:"created_at,omitempty"`
	EpochsCompleted      int32          `json:"epochs_completed,omitempty"`
	EvalSteps            int32          `json:"eval_steps,omitempty"`
	ID                   string         `json:"id"`
	JobID                string         `json:"job_id,omitempty"`
	LearningRate         float32        `json:"learning_rate,omitempty"`
	LrScheduler          map[string]any `json:"lr_scheduler,omitempty"`
	MaxGradNorm          float32        `json:"max_grad_norm,omitempty"`
	Model                string         `json:"model,omitempty"`
	ModelOutputName      string         `json:"model_output_name,omitempty"`
	ModelOutputPath      string         `json:"model_output_path,omitempty"`
	NCheckpoints         int32          `json:"n_checkpoints,omitempty"`
	NEpochs              int32          `json:"n_epochs,omitempty"`
	NEvals               int32          `json:"n_evals,omitempty"`
	ParamCount           int32          `json:"param_count,omitempty"`
	QueueDepth           int32          `json:"queue_depth,omitempty"`
	Status               string         `json:"status"`
	TokenCount           int32          `json:"token_count,omitempty"`
	TotalPrice           int32          `json:"total_price,omitempty"`
	TrainingFile         string         `json:"training_file,omitempty"`
	TrainingType         map[string]any `json:"training_type,omitempty"`
	TrainingfileNumlines int32          `json:"trainingfile_numlines,omitempty"`
	TrainingfileSize     int32          `json:"trainingfile_size,omitempty"`
	UpdatedAt            string         `json:"updated_at,omitempty"`
	ValidationFile       string         `json:"validation_file,omitempty"`
	WandbProjectName     string         `json:"wandb_project_name,omitempty"`
	WandbURL             string         `json:"wandb_url,omitempty"`
	WarmupRatio          float32        `json:"warmup_ratio,omitempty"`
	WeightDecay          float32        `json:"weight_decay,omitempty"`
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (f *FinetuneResponse) GetBatchSize() (ret int32) {
	if f == nil {
		return ret
	}
	return f.BatchSize
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (f *FinetuneResponse) GetCreatedAt() (ret string) {
	if f == nil {
		return ret
	}
	return f.CreatedAt
}

// GetEpochsCompleted returns the EpochsCompleted field value if set, zero value otherwise.
func (f *FinetuneResponse) GetEpochsCompleted() (ret int32) {
	if f == nil {
		return ret
	}
	return f.EpochsCompleted
}

// GetEvalSteps returns the EvalSteps field value if set, zero value otherwise.
func (f *FinetuneResponse) GetEvalSteps() (ret int32) {
	if f == nil {
		return ret
	}
	return f.EvalSteps
}

// GetID returns the ID field value if set, zero value otherwise.
func (f *FinetuneResponse) GetID() (ret string) {
	if f == nil {
		return ret
	}
	return f.ID
}

// GetJobID returns the JobID field value if set, zero value otherwise.
func (f *FinetuneResponse) GetJobID() (ret string) {
	if f == nil {
		return ret
	}
	return f.JobID
}

// GetLearningRate returns the LearningRate field value if set, zero value otherwise.
func (f *FinetuneResponse) GetLearningRate() (ret float32) {
	if f == nil {
		return ret
	}
	return f.LearningRate
}

// GetLrScheduler returns the LrScheduler field value if set, zero value otherwise.
func (f *FinetuneResponse) GetLrScheduler() (ret map[string]any) {
	if f == nil {
		return ret
	}
	return f.LrScheduler
}

// GetMaxGradNorm returns the MaxGradNorm field value if set, zero value otherwise.
func (f *FinetuneResponse) GetMaxGradNorm() (ret float32) {
	if f == nil {
		return ret
	}
	return f.MaxGradNorm
}

// GetModel returns the Model field value if set, zero value otherwise.
func (f *FinetuneResponse) GetModel() (ret string) {
	if f == nil {
		return ret
	}
	return f.Model
}

// GetModelOutputName returns the ModelOutputName field value if set, zero value otherwise.
func (f *FinetuneResponse) GetModelOutputName() (ret string) {
	if f == nil {
		return ret
	}
	return f.ModelOutputName
}

// GetModelOutputPath returns the ModelOutputPath field value if set, zero value otherwise.
func (f *FinetuneResponse) GetModelOutputPath() (ret string) {
	if f == nil {
		return ret
	}
	return f.ModelOutputPath
}

// GetNCheckpoints returns the NCheckpoints field value if set, zero value otherwise.
func (f *FinetuneResponse) GetNCheckpoints() (ret int32) {
	if f == nil {
		return ret
	}
	return f.NCheckpoints
}

// GetNEpochs returns the NEpochs field value if set, zero value otherwise.
func (f *FinetuneResponse) GetNEpochs() (ret int32) {
	if f == nil {
		return ret
	}
	return f.NEpochs
}

// GetNEvals returns the NEvals field value if set, zero value otherwise.
func (f *FinetuneResponse) GetNEvals() (ret int32) {
	if f == nil {
		return ret
	}
	return f.NEvals
}

// GetParamCount returns the ParamCount field value if set, zero value otherwise.
func (f *FinetuneResponse) GetParamCount() (ret int32) {
	if f == nil {
		return ret
	}
	return f.ParamCount
}

// GetQueueDepth returns the QueueDepth field value if set, zero value otherwise.
func (f *FinetuneResponse) GetQueueDepth() (ret int32) {
	if f == nil {
		return ret
	}
	return f.QueueDepth
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (f *FinetuneResponse) GetStatus() (ret string) {
	if f == nil {
		return ret
	}
	return f.Status
}

// GetTokenCount returns the TokenCount field value if set, zero value otherwise.
func (f *FinetuneResponse) GetTokenCount() (ret int32) {
	if f == nil {
		return ret
	}
	return f.TokenCount
}

// GetTotalPrice returns the TotalPrice field value if set, zero value otherwise.
func (f *FinetuneResponse) GetTotalPrice() (ret int32) {
	if f == nil {
		return ret
	}
	return f.TotalPrice
}

// GetTrainingFile returns the TrainingFile field value if set, zero value otherwise.
func (f *FinetuneResponse) GetTrainingFile() (ret string) {
	if f == nil {
		return ret
	}
	return f.TrainingFile
}

// GetTrainingType returns the TrainingType field value if set, zero value otherwise.
func (f *FinetuneResponse) GetTrainingType() (ret map[string]any) {
	if f == nil {
		return ret
	}
	return f.TrainingType
}

// GetTrainingfileNumlines returns the TrainingfileNumlines field value if set, zero value otherwise.
func (f *FinetuneResponse) GetTrainingfileNumlines() (ret int32) {
	if f == nil {
		return ret
	}
	return f.TrainingfileNumlines
}

// GetTrainingfileSize returns the TrainingfileSize field value if set, zero value otherwise.
func (f *FinetuneResponse) GetTrainingfileSize() (ret int32) {
	if f == nil {
		return ret
	}
	return f.TrainingfileSize
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (f *FinetuneResponse) GetUpdatedAt() (ret string) {
	if f == nil {
		return ret
	}
	return f.UpdatedAt
}

// GetValidationFile returns the ValidationFile field value if set, zero value otherwise.
func (f *FinetuneResponse) GetValidationFile() (ret string) {
	if f == nil {
		return ret
	}
	return f.ValidationFile
}

// GetWandbProjectName returns the WandbProjectName field value if set, zero value otherwise.
func (f *FinetuneResponse) GetWandbProjectName() (ret string) {
	if f == nil {
		return ret
	}
	return f.WandbProjectName
}

// GetWandbURL returns the WandbURL field value if set, zero value otherwise.
func (f *FinetuneResponse) GetWandbURL() (ret string) {
	if f == nil {
		return ret
	}
	return f.WandbURL
}

// GetWarmupRatio returns the WarmupRatio field value if set, zero value otherwise.
func (f *FinetuneResponse) GetWarmupRatio() (ret float32) {
	if f == nil {
		return ret
	}
	return f.WarmupRatio
}

// GetWeightDecay returns the WeightDecay field value if set, zero value otherwise.
func (f *FinetuneResponse) GetWeightDecay() (ret float32) {
	if f == nil {
		return ret
	}
	return f.WeightDecay
}
