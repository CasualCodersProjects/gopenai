package types

import (
	"bytes"
	"encoding/json"
	"io"
)

func DecodeModelsResponse(body io.Reader) (ModelsResponse, error) {
	var modelsResponse ModelsResponse
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	jsonBytes := buf.Bytes()
	err := json.Unmarshal(jsonBytes, &modelsResponse)
	return modelsResponse, err
}

func DecodeCompletionResponse(body io.Reader) (CompletionResponse, error) {
	var completionResponse CompletionResponse
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	jsonBytes := buf.Bytes()
	err := json.Unmarshal(jsonBytes, &completionResponse)
	return completionResponse, err
}

func DecodeEditResponse(body io.Reader) (EditResponse, error) {
	var editResponse EditResponse
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	jsonBytes := buf.Bytes()
	err := json.Unmarshal(jsonBytes, &editResponse)
	return editResponse, err
}

func DecodeImageResponse(body io.Reader) (ImageResponse, error) {
	var imageResponse ImageResponse
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	jsonBytes := buf.Bytes()
	err := json.Unmarshal(jsonBytes, &imageResponse)
	return imageResponse, err
}
