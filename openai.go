package goopenai

import (
	"io"
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

func (o *OpenAI) ListModels() (types.ModelsResponse, error) {
	req, err := o.newRequest("GET", "https://api.openai.com/v1/models", nil)
	if err != nil {
		return types.ModelsResponse{}, err
	}

	resp, err := o.Client.Do(req)
	if err != nil {
		return types.ModelsResponse{}, err
	}

	return types.DecodeModelsResponse(resp.Body)
}

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

	return types.DecodeCompletionResponse(resp.Body)
}

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

	return types.DecodeEditResponse(resp.Body)
}

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

	return types.DecodeImageResponse(resp.Body)
}

func (o *OpenAI) CreateCompletionSimple(prompt, model string, maxTokens int) (types.CompletionResponse, error) {
	if model == "" {
		model = "text-davinci-003"
	}
	if maxTokens == 0 {
		maxTokens = 256
	}
	completionRequest := types.CompletionRequest{
		Prompt:    prompt,
		Model:     model,
		MaxTokens: maxTokens,
	}

	return o.CreateCompletion(&completionRequest)
}
