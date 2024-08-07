package source

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"github/stable-diffusion-go/server/model/system"
)

func ChatgptOllama(text string, ollamaConfig system.SettingsOllamaConfig) (prompt string, err error) {
	config := openai.DefaultConfig("ollama")
	config.BaseURL = ollamaConfig.Url
	model := ollamaConfig.ModelName
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				//{
				//	Role:    openai.ChatMessageRoleSystem,
				//	Content: string(file),
				//},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)
	if err != nil {
		return "", errors.New("调用OpenAI API失败")
	}
	return resp.Choices[0].Message.Content, nil
}
