package core

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"github/stable-diffusion-go/server/global"
	"log"
)

func ChatgptOllama(message string) (prompt string, err error) {
	// 使用OpenAI API调用chatGPT进行翻译
	config := openai.DefaultConfig("ollama")
	config.BaseURL = global.Config.Ollama.Url
	model := global.Config.Ollama.Model
	client := openai.NewClientWithConfig(config)
	//file, err := os.ReadFile(global.SdPrompt)
	//if err != nil {
	//	log.Fatal("读取文件失败", err)
	//	return
	//}
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
