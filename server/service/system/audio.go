package system

import (
	"github/stable-diffusion-go/server/model/system"
)

type AudioService struct{}

// CreateAudioAndSrt 批量文字转图片
func (s *AudioService) CreateAudioAndSrt(params system.AudioConfig) error {
	var settings system.Settings
	return nil
}
