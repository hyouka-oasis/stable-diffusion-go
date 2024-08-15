package system

import (
	"context"
	"github.com/ollama/ollama/api"
)

type OllamaService struct{}

// GetOllamaModelList 获取ollama模型列表
func (s *OllamaService) GetOllamaModelList() (list *api.ListResponse, err error) {
	ctx := context.Background()
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return
	}
	list, err = client.List(ctx)
	if err != nil {
		return
	}
	return
}
