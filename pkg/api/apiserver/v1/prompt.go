package v1

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

type PromptResponse struct {
	Answer string `json:"answer"`
}
