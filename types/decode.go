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

// DecodeChatResponse decodes a response from the chat endpoint
func DecodeChatResponse(body io.Reader) (ChatResponse, error) {
	var chatResponse ChatResponse
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	jsonBytes := buf.Bytes()
	err := json.Unmarshal(jsonBytes, &chatResponse)
	return chatResponse, err
}

// DecodeAudioTranscriptionResponse decodes a response from the audio transcription endpoint
func DecodeAudioTranscriptionResponse(body io.Reader) (AudioTranscriptionResponse, error) {
	var audioTranscriptionResponse AudioTranscriptionResponse
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	jsonBytes := buf.Bytes()
	err := json.Unmarshal(jsonBytes, &audioTranscriptionResponse)
	return audioTranscriptionResponse, err
}
