# OpenAI Binding for Go

This is a super simple wrapper for OpenAI's GPT Completion, GPT Edits, and DALL-E Image generation APIs.

# Getting Started

## Installing

```
go get -u github.com/CasualCodersProjects/gopenai.git
```

## Usage

Initializing the OpenAI object.

```go
openAI := gopenai.NewOpenAI(&gopenai.OpenAIOpts{})
```

Passing in an API Key.

```go
openAI := gopenai.NewOpenAI(&gopenai.OpenAIOpts{
    APIKey: "your-api-key-here",
})
```

Creating a Simple Completion.

```go
// if passed an empty string for the model and zero for the tokens, 
// will default to "text-davinci-003" and 256
results, err := openAI.CreateCompletionSimple("Write me a story about a woman named Alice", "", 0)
if err != nil {
    panic(err)
}
```

Creating a completion with more options.

```go
request := types.NewDefaultCompletionRequest("Your prompt here\n")
request.MaxTokens = 1024
request.Stop = []string{"\n"}
request.N = 3

results, err := openAI.CreateCompletion(request)
if err != nil {
    panic(err)
}
```

Creating an Edit.

```go
request := types.NewDefaultEditRequest("Misspell this senstence", "Correct the spelling mistakes.")

results, err := openAI.CreateEdit(request)
if err != nil {
    panic(err)
}
```

Creating an Image.

```go
request := types.NewDefaultImageRequest("A dog in a hat")

results, err := openAI.CreateImage(request)
if err != nil {
    panic(err)
}
```

# FAQ

## Will you support other OpenAI APIs?

Not unless we need them! This library is for another Go project of ours, we figured we'd open source it so that others can use it too! It fits our needs, so updates will only be made on an as-needed basis. However, anyone is welcome to submit a Pull Request and add anything else they would also find useful!

## Are you affiliated with OpenAI?

No. This is a community project.
