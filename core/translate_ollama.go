package core

import (
	"ComicTweetsGo/global"
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
)

func ChatgptOllama(message string) (prompt string, err error) {
	// 使用OpenAI API调用chatGPT进行翻译
	config := openai.DefaultConfig("ollama")
	config.BaseURL = global.Config.Ollama.Url
	model := global.Config.Ollama.Model
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		log.Fatal("调用OpenAI API失败:", err)
		return
	}
	return resp.Choices[0].Message.Content, nil
}
