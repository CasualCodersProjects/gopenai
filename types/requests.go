package types

type CompletionRequest struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	MaxTokens        int      `json:"max_tokens"`
	Temperature      float64  `json:"temperature"`
	TopP             float64  `json:"top_p"`
	N                int      `json:"n"`
	Stream           bool     `json:"stream"`
	Stop             []string `json:"stop"`
	PresencePenalty  float64  `json:"presence_penalty"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	BestOf           int      `json:"best_of"`
	User             string   `json:"user"`
	Suffix           string   `json:"suffix"`
	Echo             bool     `json:"echo"`
}

type EditRequest struct {
	Model        string  `json:"model"`
	Input        string  `json:"input"`
	Instructions string  `json:"instruction"`
	Temperature  float64 `json:"temperature"`
	N            int     `json:"n"`
	TopP         float64 `json:"top_p"`
}

type ImageRequest struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
	User           string `json:"user"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest represents a request to the chat endpoint
type ChatRequest struct {
	Model            string        `json:"model"`
	Messages         []ChatMessage `json:"messages"`
	Temperature      float64       `json:"temperature"`
	TopP             float64       `json:"top_p"`
	N                int           `json:"n"`
	Stream           bool          `json:"stream"`
	Stop             []string      `json:"stop"`
	MaxTokens        int           `json:"max_tokens"`
	PresencePenalty  float64       `json:"presence_penalty"`
	FrequencyPenalty float64       `json:"frequency_penalty"`
	User             string        `json:"user"`
}
