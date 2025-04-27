package prompts

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

// SysCommitMsgJson is the system prompt instructing the model to generate a commit message via a function call
const SysCommitMsgJson = `You are a tool that generates a concise, descriptive git commit message summarizing the pending changes.
Respond by calling the function "commitMsg" exactly once with a JSON object containing a single property "commitMsg" set to the generated commit message. Do not include any additional text.`

// CommitMsgFn defines the structure of the JSON function call for generating a commit message
var CommitMsgFn = openai.FunctionDefinition{
	Name:        "commitMsg",
	Description: "Generate a concise commit message summarizing the pending changes",
	Parameters: &jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"commitMsg": {
				Type:        jsonschema.String,
				Description: "The commit message summarizing the changes",
			},
		},
		Required: []string{"commitMsg"},
	},
}

// CommitMsgRes represents the JSON response structure for the commit message function call
type CommitMsgRes struct {
	// CommitMsg is the generated commit message
	CommitMsg string `json:"commitMsg"`
}
