package goopenai

import (
	"os"
	"testing"

	"github.com/CasualCodersProjects/gopenai/types"
)

func TestNewInstance(t *testing.T) {
	apiKey := "test"
	openai := NewOpenAI(&OpenAIOpts{APIKey: apiKey})

	if openai.APIKey != apiKey {
		t.Errorf("Expected APIKey to be %s, got %s", apiKey, openai.APIKey)
	}
}

func TestNewInstanceWithEnv(t *testing.T) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	openai := NewOpenAI(&OpenAIOpts{})
	if openai.APIKey != apiKey {
		t.Errorf("Expected APIKey to be %s, got %s", apiKey, openai.APIKey)
	}
}

func TestListModels(t *testing.T) {
	openai := NewOpenAI(&OpenAIOpts{})
	models, err := openai.ListModels()
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if len(models.Data) == 0 {
		t.Errorf("Expected models to be greater than 0, got %d", len(models.Data))
	}
}

func TestCreateCompletion(t *testing.T) {
	openai := NewOpenAI(&OpenAIOpts{})
	completionRequest := types.NewDefaultCompletionRequest("This is a test")
	completion, err := openai.CreateCompletion(completionRequest)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if completion.ID == "" {
		t.Errorf("Expected completion ID to be set, got %s", completion.ID)
	}
	if len(completion.Choices) != 1 {
		t.Errorf("Expected completion choices to be 1, got %d", len(completion.Choices))
	}
}

func TestCreateEdit(t *testing.T) {
	openai := NewOpenAI(&OpenAIOpts{})
	editRequest := types.NewDefaultEditRequest("this iss a ttest", "Fix the spelling mistakes")
	edits, err := openai.CreateEdit(editRequest)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if len(edits.Edits) != 1 {
		t.Errorf("Expected edits choices to be 1, got %d", len(edits.Edits))
	}
}

func TestCreateImage(t *testing.T) {
	openai := NewOpenAI(&OpenAIOpts{})
	imageRequest := types.NewDefaultImageRequest("The quick brown fox jumps over the lazy dog.")
	image, err := openai.CreateImage(imageRequest)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if len(image.Data) == 1 {
		t.Errorf("Expected image ID to be set, got %s", image.Data)
	}
}
