package source

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
)

func ChatgptOllama(text string, ollamaConfig system.SettingsOllamaConfig) (prompt string, err error) {
	// 使用OpenAI API调用chatGPT进行翻译
	config := openai.DefaultConfig("ollama")
	//config.BaseURL = ollamaConfig.Url
	//model := ollamaConfig.ModelName
	config.BaseURL = global.Config.Ollama.Url
	model := global.Config.Ollama.Model
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
