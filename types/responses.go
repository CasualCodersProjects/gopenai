package types

// Model represents a model
type Model struct {
	ID         string      `json:"id"`
	Object     string      `json:"object"`
	OwnedBy    string      `json:"owned_by"`
	Permission interface{} `json:"permission"`
}

// ModelsResponse response for listing models
type ModelsResponse struct {
	Data   []Model `json:"data"`
	Object string  `json:"object"`
}

// Choice represents a choice that the API returns
type Choice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	LogProbs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

// Usage represents the usage of the API
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// CompletionResponse response for completion
type CompletionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// Edit response for search
type Edit struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

// EditResponse response for search
type EditResponse struct {
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Edits   []Edit `json:"choices"`
	Usage   Usage  `json:"usage"`
}

// Image response for an image
type Image struct {
	URL string `json:"url"`
}

// ImageResponse response for an image generation request
type ImageResponse struct {
	Created int     `json:"created"`
	Data    []Image `json:"data"`
}

type ChatResponseChoice struct {
	Index        int         `json:"index"`
	FinishReason string      `json:"finish_reason"`
	Message      ChatMessage `json:"message"`
}

// ChatResponse response for a chat request
type ChatResponse struct {
	ID      string               `json:"id"`
	Object  string               `json:"object"`
	Created int                  `json:"created"`
	Choices []ChatResponseChoice `json:"choices"`
	Usage   Usage                `json:"usage"`
}

// AudioTranscriptionResponse response for a transcription request
type AudioTranscriptionResponse struct {
	Text string `json:"text"`
}
