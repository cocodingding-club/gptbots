package chatgpt

import (
	"context"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	"os"
)

func Ask(message string) string {
	env := os.Getenv("GPT_ENV")
	if "" == env {
		env = "local"
	}
	godotenv.Load(".env." + env)

	token := os.Getenv("GPT_TOKEN")
	client := openai.NewClient(token)
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
