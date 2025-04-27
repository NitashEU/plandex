package prompts

import (
	openai "github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

// SysExecStatusJson is the system prompt for JSON-based execution status checks.
const SysExecStatusJson = `You are an execution status assistant. Analyze the current message and conversation context to determine whether the current subtask has been completed. Respond strictly with a JSON object in the following format:
{
  "reasoning": "<detailed explanation of your decision>",
  "subtaskFinished": <true|false>
}`

// ExecStatusFn defines the function schema for checking subtask completion.
var ExecStatusFn = openai.FunctionDefinition{
    Name:        "exec_status_check",
    Description: "Determines if a subtask is completed based on the current message and conversation history.",
    Parameters: &jsonschema.Definition{
        Type: "object",
        Properties: map[string]jsonschema.Definition{
            "reasoning": {
                Type:        "string",
                Description: "The reasoning explaining why the subtask is or isn’t finished.",
            },
            "subtaskFinished": {
                Type:        "boolean",
                Description: "True if the subtask has been completed; otherwise false.",
            },
        },
        Required: []string{"reasoning", "subtaskFinished"},
    },
}

// ExecStatusRes represents the JSON response from the exec_status_check function.
type ExecStatusRes struct {
    Reasoning       string `json:"reasoning"`
    SubtaskFinished bool   `json:"subtaskFinished"`
}
