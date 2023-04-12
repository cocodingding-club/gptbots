package chatgpt

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
)

func Ask(message string) string {
	client := openai.NewClient("token")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return "ChatCompletion error: %v\n"
	}

	return resp.Choices[0].Message.Content
}
