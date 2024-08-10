package source

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"github/stable-diffusion-go/server/model/system"
	"regexp"
)

var extractEnglish = regexp.MustCompile(`-\s*(.*?)\s*\(English\)`)

func ChatgptOllama(text string, ollamaConfig system.SettingsOllamaConfig, message *[]openai.ChatCompletionMessage) (prompt string, err error) {
	config := openai.DefaultConfig("ollama")
	config.BaseURL = ollamaConfig.Url
	model := ollamaConfig.ModelName
	client := openai.NewClientWithConfig(config)
	*message = append(*message, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: text,
	})
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: *message,
		},
	)
	if err != nil {
		return "", errors.New("调用OpenAI API失败")
	}
	*message = append(*message, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: resp.Choices[0].Message.Content,
	})
	return resp.Choices[0].Message.Content, nil
}
