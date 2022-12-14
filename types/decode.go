package types

import (
	"bytes"
	"encoding/json"
	"io"
)

// DecodeModelsResponse decodes a response from the models endpoint
func DecodeModelsResponse(body io.Reader) (ModelsResponse, error) {
	var modelsResponse ModelsResponse
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	jsonBytes := buf.Bytes()
	err := json.Unmarshal(jsonBytes, &modelsResponse)
	return modelsResponse, err
}

// DecodeCompletionResponse decodes a response from the completion endpoint
func DecodeCompletionResponse(body io.Reader) (CompletionResponse, error) {
	var completionResponse CompletionResponse
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	jsonBytes := buf.Bytes()
	err := json.Unmarshal(jsonBytes, &completionResponse)
	return completionResponse, err
}

// DecodeEditResponse decodes a response from the edit endpoint
func DecodeEditResponse(body io.Reader) (EditResponse, error) {
	var editResponse EditResponse
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	jsonBytes := buf.Bytes()
	err := json.Unmarshal(jsonBytes, &editResponse)
	return editResponse, err
}

// DecodeImageResponse decodes a response from the image endpoint
func DecodeImageResponse(body io.Reader) (ImageResponse, error) {
	var imageResponse ImageResponse
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	jsonBytes := buf.Bytes()
	err := json.Unmarshal(jsonBytes, &imageResponse)
	return imageResponse, err
}
