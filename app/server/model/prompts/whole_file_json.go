package prompts

import (
	openai "github.com/sashabaranov/go-openai"
	jsonschema "github.com/sashabaranov/go-openai/jsonschema"
)

// SysWholeFileJson is the system prompt instructing the model to return the entire merged file
// content via a JSON function call without any additional text or formatting.
const SysWholeFileJson = `You are an AI coding assistant. You must provide the entire merged file with all proposed updates applied. Respond by invoking the function "wholeFile" with a single argument:

{
  "wholeFile": "<entire file content here>"
}

Do not include any additional text, formatting, or tags. Only return the function call in JSON format.`

// WholeFileFn is the OpenAI function definition guiding the model on how to structure the function call.
var WholeFileFn = openai.FunctionDefinition{
	Name:        "wholeFile",
	Description: "Returns the entire merged file content with proposed changes applied",
	Parameters: &jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"wholeFile": {
				Type:        jsonschema.String,
				Description: "The full file content with all changes applied",
			},
		},
		Required: []string{"wholeFile"},
	},
}

// WholeFileRes represents the JSON structure returned by the model when calling the "wholeFile" function.
type WholeFileRes struct {
	WholeFile string `json:"wholeFile"`
}
