package openai

import (
	"bongo/config"
	"context"
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

var client *gogpt.Client
var ctx context.Context

func Setup() {
	client = gogpt.NewClient(config.OpenAIToken)
	ctx = context.Background()
}

func SendPrompt(p string) (string, error) {
	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   2000,
		Temperature: 0.7,
		Prompt:      p,
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	fmt.Println(resp.Choices[0].Text)

	return resp.Choices[0].Text, nil
}
