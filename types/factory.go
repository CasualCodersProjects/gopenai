package types

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

func NewDefaultImageRequest(prompt string) *ImageRequest {
	return &ImageRequest{
		Prompt:         prompt,
		N:              1,
		Size:           "256x256",
		ResponseFormat: "url",
		User:           "",
	}
}
