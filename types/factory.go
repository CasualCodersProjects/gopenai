package types

// NewDefaultCompletionRequest returns a CompletionRequest with default values
func NewDefaultCompletionRequest(prompt string) *CompletionRequest {
	return &CompletionRequest{
		Model:            "text-davinci-003",
		MaxTokens:        256,
		Prompt:           prompt,
		Suffix:           "",
		Temperature:      1,
		TopP:             1,
		N:                1,
		Stream:           false,
		Stop:             nil,
		Echo:             false,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		BestOf:           1,
		User:             "",
	}
}

// NewDefaultEditRequest returns a EditRequest with default values
func NewDefaultEditRequest(input, instructions string) *EditRequest {
	return &EditRequest{
		Model:        "text-davinci-edit-001",
		Input:        input,
		Instructions: instructions,
		Temperature:  1,
		N:            1,
		TopP:         1,
	}
}

// NewDefaultImageRequest returns a ImageRequest with default values
func NewDefaultImageRequest(prompt string) *ImageRequest {
	return &ImageRequest{
		Prompt:         prompt,
		N:              1,
		Size:           "256x256",
		ResponseFormat: "url",
		User:           "",
	}
}

// NewDefaultChatRequest returns a ChatRequest with default values
func NewDefaultChatRequest(prompt string) *ChatRequest {
	messages := []ChatMessage{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: prompt},
	}

	return &ChatRequest{
		Model:            "gpt-3.5-turbo",
		Messages:         messages,
		Temperature:      1,
		TopP:             1,
		N:                1,
		Stream:           false,
		Stop:             nil,
		MaxTokens:        1024,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		User:             "",
	}
}
