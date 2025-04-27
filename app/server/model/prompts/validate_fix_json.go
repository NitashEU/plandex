package prompts

import (
	"github.com/sashabaranov/go-openai"
	jsonschema "github.com/sashabaranov/go-openai/jsonschema"
)

const SysValidateFixJson = `Please perform a JSON function call named "validate_fix" on the provided code. The function should return an object with a single key "replacements", which is an array of objects each containing:
- "old": the original text snippet that needs to be replaced.
- "new": the corrected text snippet.
Output only the JSON function call with its arguments and nothing else.`

var ValidateFixFn = openai.FunctionDefinition{
	Name:        "validate_fix",
	Description: "Return JSON object with replacements for fixing syntax errors and applying updates",
	Parameters: &jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"replacements": {
				Type:        jsonschema.Array,
				Description: "List of replacement operations",
				Items: &jsonschema.Definition{
					Type: jsonschema.Object,
					Properties: map[string]jsonschema.Definition{
						"old": {Type: jsonschema.String, Description: "Original text to be replaced"},
						"new": {Type: jsonschema.String, Description: "Replacement text"},
					},
					Required: []string{"old", "new"},
				},
			},
		},
		Required: []string{"replacements"},
	},
}

type ReplacementJson struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type ValidateFixRes struct {
	Replacements []ReplacementJson `json:"replacements"`
}
