package shared

import (
	"github.com/davecgh/go-spew/spew"
)

/*
'MaxTokens' is the absolute input limit for the provider.

'MaxOutputTokens' is the absolute output limit for the provider.

'ReservedOutputTokens' is how much we set aside in context for the model to use in its output. It's more of a realistic output limit, since for some models, the hard maximum 'MaxTokens' is actually equal to the input limit, which would leave no room for input.

The effective input limit is 'MaxTokens' - 'ReservedOutputTokens'.

For example, OpenAI o3-mini has a MaxTokens of 200k and a MaxOutputTokens of 100k. But in practice, we are very unlikely to use all the output tokens, and we want to leave more space for input. So we set ReservedOutputTokens to 40k, allowing ~25k for reasoning tokens, as well as ~15k for real output tokens, which is enough for most use cases. The new effective input limit is therefore 200k - 40k = 160k. However, these are not passed through as hard limits. So if we have a smaller amount of input (under 100k) the model could still use up to the full 100k output tokens if necessary.

For models with a low output limit, we just set ReservedOutputTokens to the MaxOutputTokens.

When checking for sufficient credits on Plandex Cloud, we use MaxOutputTokens-InputTokens, since this is the maximum that could hypothetically be used.

'DefaultMaxConvoTokens' is the default maximum number of conversation tokens that are allowed before we start using gradual summarization to shorten the conversation.

'ModelName' is the name of the model on the provider's side.

'ModelId' is the identifier for the model on the Plandex side—it must be unique per provider. We have this so that models with the same name and provider, but different settings can be differentiated.

'ModelCompatibility' is used to check for feature support (like image support).

'BaseUrl' is the base URL for the provider.

'PreferredModelOutputFormat' is the preferred output format for the model—currently either 'ModelOutputFormatToolCallJson' or 'ModelOutputFormatXml' — OpenAI models like JSON (and benefit from strict JSON schemas), while most other providers are unreliable for JSON generation and do better with XML, even if they claim to support JSON.

'RoleParamsDisabled' is used to disable role-based parameters like temperature, top_p, etc. for the model—OpenAI early releases often don't allow changes to these.

'SystemPromptDisabled' is used to disable the system prompt for the model—OpenAI early releases sometimes don't allow system prompts.

'ReasoningEffortEnabled' is used to enable reasoning effort for the model (like OpenAI's o3-mini).

'ReasoningEffort' is the reasoning effort for the model, when 'ReasoningEffortEnabled' is true.

'PredictedOutputEnabled' is used to enable predicted output for the model (currently only supported by gpt-4o).

'ApiKeyEnvVar' is the environment variable that contains the API key for the model.
*/

var AvailableModels = []*AvailableModel{
	// Direct OpenAI models
	{
		Description:           "OpenAI o3-mini-high",
		DefaultMaxConvoTokens: 10000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenAI,
			ModelName:                  "o3-mini",
			ModelId:                    "openai/o3-mini-high",
			MaxTokens:                  200000,
			MaxOutputTokens:            100000,
			ReservedOutputTokens:       30000,
			ApiKeyEnvVar:               OpenAIEnvVar,
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    OpenAIV1BaseUrl,
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
			RoleParamsDisabled:         true,
			ReasoningEffortEnabled:     true,
			ReasoningEffort:            ReasoningEffortHigh,
		},
	},
	{
		Description:           "OpenAI o3-mini-medium",
		DefaultMaxConvoTokens: 10000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenAI,
			ModelName:                  "o3-mini",
			ModelId:                    "openai/o3-mini-medium",
			MaxTokens:                  200000,
			MaxOutputTokens:            100000,
			ReservedOutputTokens:       40000, // 25k for reasoning, 15k for output
			ApiKeyEnvVar:               OpenAIEnvVar,
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    OpenAIV1BaseUrl,
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
			RoleParamsDisabled:         true,
			ReasoningEffortEnabled:     true,
			ReasoningEffort:            ReasoningEffortMedium,
		},
	},
	{
		Description:           "OpenAI o3-mini-low",
		DefaultMaxConvoTokens: 10000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenAI,
			ModelName:                  "o3-mini",
			ModelId:                    "openai/o3-mini-low",
			MaxTokens:                  200000,
			MaxOutputTokens:            100000,
			ReservedOutputTokens:       40000, // 25k for reasoning, 15k for output
			ApiKeyEnvVar:               OpenAIEnvVar,
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    OpenAIV1BaseUrl,
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
			RoleParamsDisabled:         true,
			ReasoningEffortEnabled:     true,
			ReasoningEffort:            ReasoningEffortLow,
		},
	},
	{
		Description:           "OpenAI o1",
		DefaultMaxConvoTokens: 15000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenAI,
			ModelName:                  "o1",
			ModelId:                    "openai/o1",
			MaxTokens:                  200000,
			MaxOutputTokens:            100000,
			ReservedOutputTokens:       40000, // 25k for reasoning, 15k for output
			ApiKeyEnvVar:               OpenAIEnvVar,
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    OpenAIV1BaseUrl,
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
			SystemPromptDisabled:       true,
			RoleParamsDisabled:         true,
		},
	},

	// o1-pro is not supported yet - more work is needed to support the responses API
	// {
	// 	Description:           "OpenAI o1-pro",
	// 	DefaultMaxConvoTokens: 15000,
	// 	BaseModelConfig: BaseModelConfig{
	// 		Provider:                   ModelProviderOpenAI,
	// 		ModelName:                  "o1-pro",
	// 		ModelId:                    "openai/o1-pro",
	// 		MaxTokens:                  200000,
	// 		MaxOutputTokens:            100000,
	// 		ReservedOutputTokens:       60000,
	// 		ApiKeyEnvVar:               OpenAIEnvVar,
	// 		ModelCompatibility:         fullCompatibility,
	// 		BaseUrl:                    OpenAIV1BaseUrl,
	// 		PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
	// 		RoleParamsDisabled:         true,
	// 		UsesOpenAIResponsesAPI:     true,
	// 	},
	// },

	{
		Description:           "OpenAI gpt-4.1",
		DefaultMaxConvoTokens: 15000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenAI,
			ModelName:                  "gpt-4.1",
			ModelId:                    "openai/gpt-4.1",
			MaxTokens:                  1047576,
			MaxOutputTokens:            32768,
			ReservedOutputTokens:       32768,
			ApiKeyEnvVar:               OpenAIEnvVar,
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    OpenAIV1BaseUrl,
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
		},
	},

	// OpenRouter models
	{
		Description:           "Anthropic Claude 3.7 Sonnet via OpenRouter",
		DefaultMaxConvoTokens: 15000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenRouter,
			ModelName:                  "anthropic/claude-3.7-sonnet",
			ModelId:                    "anthropic/claude-3.7-sonnet",
			MaxTokens:                  200000,
			MaxOutputTokens:            128000,
			ReservedOutputTokens:       20000,
			SupportsCacheControl:       true,
			ApiKeyEnvVar:               ApiKeyByProvider[ModelProviderOpenRouter],
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    BaseUrlByProvider[ModelProviderOpenRouter],
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
		},
	},
	{
		Description:           "Anthropic Claude 3.7 Sonnet (thinking) via OpenRouter",
		DefaultMaxConvoTokens: 15000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenRouter,
			ModelName:                  "anthropic/claude-3.7-sonnet:thinking",
			ModelId:                    "anthropic/claude-3.7-sonnet:thinking",
			MaxTokens:                  200000,
			MaxOutputTokens:            128000,
			ReservedOutputTokens:       40000,
			SupportsCacheControl:       true,
			ApiKeyEnvVar:               ApiKeyByProvider[ModelProviderOpenRouter],
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    BaseUrlByProvider[ModelProviderOpenRouter],
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
			IncludeReasoning:           true,
		},
	},
	{
		Description:           "Google Gemini Pro 2.5 Experimental via OpenRouter",
		DefaultMaxConvoTokens: 75000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenRouter,
			ModelName:                  "google/gemini-2.5-pro-preview-03-25",
			ModelId:                    "google/gemini-2.5-pro-preview-03-25",
			MaxTokens:                  1000000,
			MaxOutputTokens:            65535,
			ReservedOutputTokens:       65535,
			ApiKeyEnvVar:               ApiKeyByProvider[ModelProviderOpenRouter],
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    BaseUrlByProvider[ModelProviderOpenRouter],
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
		},
	},
	{
		Description:           "Google Gemini Flash 2.5 via OpenRouter",
		DefaultMaxConvoTokens: 75000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenRouter,
			ModelName:                  "google/gemini-2.5-flash-preview",
			ModelId:                    "google/gemini-2.5-flash-preview",
			MaxTokens:                  1000000,
			MaxOutputTokens:            8192,
			ReservedOutputTokens:       8192,
			ApiKeyEnvVar:               ApiKeyByProvider[ModelProviderOpenRouter],
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    BaseUrlByProvider[ModelProviderOpenRouter],
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
		},
	},

	// OpenAI models via OpenRouter
	{
		Description:           "OpenAI o3-mini-high via OpenRouter",
		DefaultMaxConvoTokens: 10000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenRouter,
			ModelName:                  "openai/o3-mini",
			ModelId:                    "openai/o3-mini-high",
			MaxTokens:                  200000,
			MaxOutputTokens:            100000,
			ReservedOutputTokens:       40000,
			ApiKeyEnvVar:               ApiKeyByProvider[ModelProviderOpenRouter],
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    BaseUrlByProvider[ModelProviderOpenRouter],
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
			SystemPromptDisabled:       true,
			RoleParamsDisabled:         true,
			ReasoningEffortEnabled:     true,
			ReasoningEffort:            ReasoningEffortHigh,
		},
	},
	{
		Description:           "OpenAI o3-mini-medium via OpenRouter",
		DefaultMaxConvoTokens: 10000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenRouter,
			ModelName:                  "openai/o3-mini",
			ModelId:                    "openai/o3-mini-medium",
			MaxTokens:                  200000,
			MaxOutputTokens:            100000,
			ReservedOutputTokens:       40000,
			ApiKeyEnvVar:               ApiKeyByProvider[ModelProviderOpenRouter],
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    BaseUrlByProvider[ModelProviderOpenRouter],
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
			SystemPromptDisabled:       true,
			RoleParamsDisabled:         true,
			ReasoningEffortEnabled:     true,
			ReasoningEffort:            ReasoningEffortMedium,
		},
	},
	{
		Description:           "OpenAI o3-mini-low via OpenRouter",
		DefaultMaxConvoTokens: 10000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenRouter,
			ModelName:                  "openai/o3-mini",
			ModelId:                    "openai/o3-mini-low",
			MaxTokens:                  200000,
			MaxOutputTokens:            100000,
			ReservedOutputTokens:       40000,
			ApiKeyEnvVar:               ApiKeyByProvider[ModelProviderOpenRouter],
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    BaseUrlByProvider[ModelProviderOpenRouter],
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
			SystemPromptDisabled:       true,
			RoleParamsDisabled:         true,
			ReasoningEffortEnabled:     true,
			ReasoningEffort:            ReasoningEffortLow,
		},
	},
	{
		Description:           "OpenAI o1 via OpenRouter",
		DefaultMaxConvoTokens: 15000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenRouter,
			ModelName:                  "openai/o1",
			ModelId:                    "openai/o1",
			MaxTokens:                  200000,
			MaxOutputTokens:            100000,
			ReservedOutputTokens:       40000,
			ApiKeyEnvVar:               ApiKeyByProvider[ModelProviderOpenRouter],
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    BaseUrlByProvider[ModelProviderOpenRouter],
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
			SystemPromptDisabled:       true,
			RoleParamsDisabled:         true,
		},
	},
	{
		Description:           "OpenAI gpt-4o via OpenRouter",
		DefaultMaxConvoTokens: 10000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenRouter,
			ModelName:                  "openai/gpt-4o",
			ModelId:                    "openai/gpt-4o",
			MaxTokens:                  128000,
			MaxOutputTokens:            16384,
			ReservedOutputTokens:       16384,
			ApiKeyEnvVar:               ApiKeyByProvider[ModelProviderOpenRouter],
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    BaseUrlByProvider[ModelProviderOpenRouter],
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
			PredictedOutputEnabled:     true,
		},
	},
	{
		Description:           "OpenAI gpt-4o-mini via OpenRouter",
		DefaultMaxConvoTokens: 10000,
		BaseModelConfig: BaseModelConfig{
			Provider:                   ModelProviderOpenRouter,
			ModelName:                  "openai/gpt-4o-mini",
			ModelId:                    "openai/gpt-4o-mini",
			MaxTokens:                  128000,
			MaxOutputTokens:            16384,
			ReservedOutputTokens:       16384,
			ApiKeyEnvVar:               ApiKeyByProvider[ModelProviderOpenRouter],
			ModelCompatibility:         fullCompatibility,
			BaseUrl:                    BaseUrlByProvider[ModelProviderOpenRouter],
			PreferredModelOutputFormat: ModelOutputFormatToolCallJson,
			PredictedOutputEnabled:     true,
		},
	},
}

var AvailableModelsByComposite = map[string]*AvailableModel{}

func init() {
	for _, model := range AvailableModels {
		if model.Description == "" {
			spew.Dump(model)
			panic("description is not set")
		}

		if model.Provider == "" {
			spew.Dump(model)
			panic("model provider is not set")
		}
		if model.ModelId == "" {
			spew.Dump(model)
			panic("model id is not set")
		}

		if model.DefaultMaxConvoTokens == 0 {
			spew.Dump(model)
			panic("default max convo tokens is not set")
		}

		if model.MaxTokens == 0 {
			spew.Dump(model)
			panic("max tokens is not set")
		}

		if model.MaxOutputTokens == 0 {
			spew.Dump(model)
			panic("max output tokens is not set")
		}

		if model.ReservedOutputTokens == 0 {
			spew.Dump(model)
			panic("reserved output tokens is not set")
		}

		if model.ApiKeyEnvVar == "" {
			spew.Dump(model)
			panic("api key env var is not set")
		}

		if model.BaseUrl == "" {
			spew.Dump(model)
			panic("base url is not set")
		}

		if model.PreferredModelOutputFormat == "" {
			spew.Dump(model)
			panic("preferred model output format is not set")
		}

		compositeKey := string(model.Provider) + "/" + string(model.ModelId)
		AvailableModelsByComposite[compositeKey] = model
	}
}

func GetAvailableModel(provider ModelProvider, modelId ModelId) *AvailableModel {
	compositeKey := string(provider) + "/" + string(modelId)
	return AvailableModelsByComposite[compositeKey]
}
