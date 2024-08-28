package source

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"github/stable-diffusion-go/server/model/system"
	"regexp"
)

var extractEnglish = regexp.MustCompile(`-\s*(.*?)\s*\(English\)`)

func OpenaiClient(ollamaConfig system.SettingsOllamaConfig, message *[]openai.ChatCompletionMessage) (prompt string, err error) {
	config := openai.DefaultConfig("ollama")
	config.BaseURL = ollamaConfig.Url
	model := ollamaConfig.ModelName
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: *message,
		},
	)
	if err != nil {
		return "", errors.New("调用OpenAI API失败" + err.Error())
	}
	return resp.Choices[0].Message.Content, nil
}

func ChatgptOllama(text string, ollamaConfig system.SettingsOllamaConfig, openContext bool, messageList *[]openai.ChatCompletionMessage) (prompt string, err error) {
	if openContext {
		*messageList = append(*messageList, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})
		content, err := OpenaiClient(ollamaConfig, messageList)
		if err != nil {
			return "", errors.New(err.Error())
		}
		*messageList = append(*messageList, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
		return content, nil
	} else {
		var message []openai.ChatCompletionMessage
		message = append(message, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})
		content, err := OpenaiClient(ollamaConfig, &message)
		if err != nil {
			return "", errors.New(err.Error())
		}
		return content, nil
	}
}
