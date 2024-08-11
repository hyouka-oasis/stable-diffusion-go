package source

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github/stable-diffusion-go/server/model/system"
	"regexp"
)

var extractEnglish = regexp.MustCompile(`-\s*(.*?)\s*\(English\)`)

func ChatgptOllama(text string, ollamaConfig system.SettingsOllamaConfig, openContext bool, message *[]openai.ChatCompletionMessage) (prompt string, err error) {
	config := openai.DefaultConfig("ollama")
	config.BaseURL = ollamaConfig.Url
	model := ollamaConfig.ModelName
	client := openai.NewClientWithConfig(config)
	if openContext {
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
		fmt.Println(resp.Choices[0].Message.Content, "输出内容")
		*message = append(*message, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: resp.Choices[0].Message.Content,
		})
		return resp.Choices[0].Message.Content, nil
	} else {
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: model,
				Messages: []openai.ChatCompletionMessage{
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
}
