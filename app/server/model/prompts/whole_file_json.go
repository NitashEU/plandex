package prompts

import (
	shared "plandex-shared"

	openai "github.com/sashabaranov/go-openai"
	jsonschema "github.com/sashabaranov/go-openai/jsonschema"
)

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

func GetWholeFileJsonPrompt(filePath string, preBuildStateWithLineNums shared.LineNumberedTextType, changesWithLineNumsType shared.LineNumberedTextType, changesDesc string, comments string) (string, int) {
	s := getBuildPromptHead(filePath, preBuildStateWithLineNums, changesDesc, changesWithLineNumsType)

	headNumTokens := shared.GetNumTokensEstimate(s)

	s += "## Comments\n\n"

	if comments != "" {
		s += comments + "\n\n"
	} else {
		s += CommentClassifierPrompt + "\n\n"
	}

	s += WholeFilePromptJson

	return s, headNumTokens
}

const WholeFilePromptJson = `
## Whole File 

Output the *entire merged file* with the *proposed updates* correctly applied. ALL reference comments will be replaced by the appropriate code from the *original file*. You will correctly merge the code from the *original file* with the *proposed updates* and output the entire file.

ALL identified reference comments MUST be replaced by the appropriate code from the *original file*. You MUST correctly merge the code from the *original file* with the *proposed updates* and output the *entire* resulting file. The resulting file MUST NOT include any reference comments.

The resulting file MUST be syntactically and semantically correct. All code structures must be properly balanced.

Respond by invoking the function "wholeFile" with a single argument:

{
  "wholeFile": "<entire file content here>"
}

Do NOT include line numbers in the <PlandexWholeFile> element. Do NOT include reference comments in the <PlandexWholeFile> element. Output the ENTIRE file, no matter how long it is, with NO EXCEPTIONS. Include the resulting file *only* with no other text. Do NOT wrap the file output in triple backticks or any other formatting, except for the <PlandexWholeFile> element tags.

Do NOT include any additional text after the <PlandexWholeFile> element. The output must end after </PlandexWholeFile>. DO NOT use the string <PlandexWholeFile> anywhere else in the output. ONLY use it to start the <PlandexWholeFile> element.

Do NOT UNDER ANY CIRCUMSTANCES *remove or change* any code that is not part of the changes in the *proposed updates*. ALL OTHER code from the *original file* must be reproduced *exactly* as it is in the *original file*. Do NOT remove comments, logging statements, commented out code, or anything else that is not part of the changes in the *proposed updates*. Your job is *only* to *apply* the changes in the *proposed updates* to the *original file*, not to make additional changes of *any kind*.

The ABSOLUTE MOST IMPORTANT THING is to leave all existing code that is not DIRECTLY part of the changes in the *proposed updates* *exactly* as it is in the *original file*. Do NOT remove any code that is not part of the changes in the *proposed updates*. Do NOT include any reference comments in the output; replace them with the appropriate code from the *original file*. Be ABSOLUTELY CERTAIN you have not left anything out which belongs in the final result.
`
