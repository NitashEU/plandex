package shared

var StrongModelPack ModelPack

var BuiltInModelPacks = []*ModelPack{
	&StrongModelPack,
}

var DefaultModelPack *ModelPack = &StrongModelPack

func init() {
	StrongModelPack = ModelPack{
		Name:        "strong",
		Description: "For difficult tasks where slower responses and builds are ok. Uses o1 for architecture and planning, claude-3.7-sonnet for implementation, prioritizes reliability over speed for builds. Supports up to 160k input context.",
		Planner: PlannerRoleConfig{
			ModelRoleConfig:    *gemini25pro(ModelRolePlanner, nil),
			PlannerModelConfig: getPlannerModelConfig(ModelProviderOpenRouter, "google/gemini-2.5-pro-exp-03-25"),
		},
		Architect:        gemini25pro(ModelRoleArchitect, nil),
		Coder:            claude37SonnetThinking(ModelRoleCoder, nil),
		PlanSummary:      *openaio3miniLow(ModelRolePlanSummary, nil),
		Builder:          *openaio3miniHigh(ModelRoleBuilder, nil),
		WholeFileBuilder: openaio3miniHigh(ModelRoleWholeFileBuilder, nil),
		Namer:            *openaio3miniHigh(ModelRoleName, nil),
		CommitMsg:        *openaio3miniHigh(ModelRoleCommitMsg, nil),
		ExecStatus:       *openaio3miniMedium(ModelRoleExecStatus, nil),
	}
}

type modelConfig struct {
	largeContextFallback *ModelRoleConfig
	largeOutputFallback  *ModelRoleConfig
	// errorFallback        *ModelRoleConfig
	strongModel *ModelRoleConfig
}

func getModelConfig(role ModelRole, provider ModelProvider, modelId ModelId, fallbacks *modelConfig) *ModelRoleConfig {
	if fallbacks == nil {
		fallbacks = &modelConfig{}
	}

	return &ModelRoleConfig{
		Role:            role,
		BaseModelConfig: GetAvailableModel(provider, modelId).BaseModelConfig,
		Temperature:     DefaultConfigByRole[role].Temperature,
		TopP:            DefaultConfigByRole[role].TopP,

		LargeContextFallback: fallbacks.largeContextFallback,
		LargeOutputFallback:  fallbacks.largeOutputFallback,
		// ErrorFallback:        fallbacks.errorFallback,
		StrongModel: fallbacks.strongModel,
	}
}

func claude37Sonnet(role ModelRole, fallbacks *modelConfig) *ModelRoleConfig {
	return getModelConfig(role, ModelProviderOpenRouter, "anthropic/claude-3.7-sonnet", fallbacks)
}

func claude37SonnetThinking(role ModelRole, fallbacks *modelConfig) *ModelRoleConfig {
	return getModelConfig(role, ModelProviderOpenRouter, "anthropic/claude-3.7-sonnet:thinking", fallbacks)
}

func openai41(role ModelRole, fallbacks *modelConfig) *ModelRoleConfig {
	return getModelConfig(role, ModelProviderOpenAI, "openai/gpt-4.1", fallbacks)
}

func openaio3miniHigh(role ModelRole, fallbacks *modelConfig) *ModelRoleConfig {
	return getModelConfig(role, ModelProviderOpenAI, "openai/o3-mini-high", fallbacks)
}

func openaio3miniMedium(role ModelRole, fallbacks *modelConfig) *ModelRoleConfig {
	return getModelConfig(role, ModelProviderOpenAI, "openai/o3-mini-medium", fallbacks)
}

func openaio3miniLow(role ModelRole, fallbacks *modelConfig) *ModelRoleConfig {
	return getModelConfig(role, ModelProviderOpenAI, "openai/o3-mini-low", fallbacks)
}

func openaio3(role ModelRole, fallbacks *modelConfig) *ModelRoleConfig {
	return getModelConfig(role, ModelProviderOpenAI, "openai/o3", fallbacks)
}

func gemini25pro(role ModelRole, fallbacks *modelConfig) *ModelRoleConfig {
	return getModelConfig(role, ModelProviderOpenRouter, "google/gemini-2.5-pro-exp-03-25", fallbacks)
}
