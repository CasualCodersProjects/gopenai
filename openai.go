package goopenai

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/CasualCodersProjects/gopenai/types"
)

type OpenAI struct {
	APIKey string
	Client *http.Client
}

type OpenAIOpts struct {
	APIKey string
}

// NewOpenAI creates a new instance of the OpenAI API
func NewOpenAI(options *OpenAIOpts) OpenAI {
	apiKey := options.APIKey
	if apiKey == "" {
		apiKey = os.Getenv("OPENAI_API_KEY")
	}
	return OpenAI{
		APIKey: apiKey,
		Client: http.DefaultClient,
	}
}

func (o *OpenAI) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Authorization", "Bearer "+o.APIKey)

	return req, nil
}

// ListModels returns a list of models
func (o *OpenAI) ListModels() (types.ModelsResponse, error) {
	req, err := o.newRequest("GET", "https://api.openai.com/v1/models", nil)
	if err != nil {
		return types.ModelsResponse{}, err
	}

	resp, err := o.Client.Do(req)
	if err != nil {
		return types.ModelsResponse{}, err
	}

	defer resp.Body.Close()

	return types.DecodeModelsResponse(resp.Body)
}

// CreateCompletion creates a completion
func (o *OpenAI) CreateCompletion(completionRequest *types.CompletionRequest) (types.CompletionResponse, error) {
	body, err := types.Encode(completionRequest)
	if err != nil {
		return types.CompletionResponse{}, err
	}

	req, err := o.newRequest("POST", "https://api.openai.com/v1/completions", body)
	if err != nil {
		return types.CompletionResponse{}, err
	}

	resp, err := o.Client.Do(req)
	if err != nil {
		return types.CompletionResponse{}, err
	}

	defer resp.Body.Close()

	return types.DecodeCompletionResponse(resp.Body)
}

// CreateEdit creates an edit
func (o *OpenAI) CreateEdit(editRequest *types.EditRequest) (types.EditResponse, error) {
	body, err := types.Encode(editRequest)
	if err != nil {
		return types.EditResponse{}, err
	}

	req, err := o.newRequest("POST", "https://api.openai.com/v1/edits", body)
	if err != nil {
		return types.EditResponse{}, err
	}

	resp, err := o.Client.Do(req)
	if err != nil {
		return types.EditResponse{}, err
	}

	defer resp.Body.Close()

	return types.DecodeEditResponse(resp.Body)
}

// CreateImage creates an image
func (o *OpenAI) CreateImage(imageRequest *types.ImageRequest) (types.ImageResponse, error) {
	body, err := types.Encode(imageRequest)
	if err != nil {
		return types.ImageResponse{}, err
	}

	req, err := o.newRequest("POST", "https://api.openai.com/v1/images", body)
	if err != nil {
		return types.ImageResponse{}, err
	}

	resp, err := o.Client.Do(req)
	if err != nil {
		return types.ImageResponse{}, err
	}

	defer resp.Body.Close()

	return types.DecodeImageResponse(resp.Body)
}

// CreateCompletionSimple creates a completion with a simple interface
func (o *OpenAI) CreateCompletionSimple(prompt, model string, maxTokens int) (types.CompletionResponse, error) {

	completionRequest := types.NewDefaultCompletionRequest(prompt)

	if model != "" {
		completionRequest.Model = model
	}

	if maxTokens > 0 {
		completionRequest.MaxTokens = maxTokens
	}

	return o.CreateCompletion(completionRequest)
}

// CreateChat creates a chat
func (o *OpenAI) CreateChat(chatRequest *types.ChatRequest) (types.ChatResponse, error) {
	body, err := types.Encode(chatRequest)
	if err != nil {
		return types.ChatResponse{}, err
	}

	req, err := o.newRequest("POST", "https://api.openai.com/v1/chat/completions", body)
	if err != nil {
		return types.ChatResponse{}, err
	}

	resp, err := o.Client.Do(req)
	if err != nil {
		return types.ChatResponse{}, err
	}

	defer resp.Body.Close()

	return types.DecodeChatResponse(resp.Body)
}

// CreateChatSimple creates a chat with a simple interface
func (o *OpenAI) CreateChatSimple(prompt string, maxTokens int) (types.ChatResponse, error) {
	chatRequest := types.NewDefaultChatRequest(prompt)

	if maxTokens > 0 {
		chatRequest.MaxTokens = maxTokens
	}

	return o.CreateChat(chatRequest)
}

func (o *OpenAI) CreateTranscription(file []byte, filename string) (types.AudioTranscriptionResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return types.AudioTranscriptionResponse{}, err
	}

	_, err = part.Write(file)
	if err != nil {
		return types.AudioTranscriptionResponse{}, err
	}

	err = writer.WriteField("model", "whisper-1")
	if err != nil {
		return types.AudioTranscriptionResponse{}, err
	}

	err = writer.Close()
	if err != nil {
		return types.AudioTranscriptionResponse{}, err
	}

	// didn't do newRequest because it doesn't support multipart
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/audio/transcriptions", body)
	if err != nil {
		return types.AudioTranscriptionResponse{}, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+o.APIKey)

	resp, err := o.Client.Do(req)
	if err != nil {
		return types.AudioTranscriptionResponse{}, err
	}

	defer resp.Body.Close()

	return types.DecodeAudioTranscriptionResponse(resp.Body)
}
