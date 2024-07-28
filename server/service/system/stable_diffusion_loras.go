package system

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/example"
	"github/stable-diffusion-go/server/model/system"
)

type StableDiffusionLorasService struct{}

// GetStableDiffusionLorasList 获取stable-diffusion列表
func (s *StableDiffusionLorasService) GetStableDiffusionLorasList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.StableDiffusionLoras{})
	err = db.Count(&total).Error
	var stableDiffusionList []*system.StableDiffusionLoras
	if err != nil {
		return stableDiffusionList, total, err
	}
	db = db.Limit(limit).Offset(offset)
	OrderStr := "id desc"
	err = db.Order(OrderStr).Find(&stableDiffusionList).Error
	fmt.Println(stableDiffusionList, "stableDiffusionList")
	for _, stableDiffusion := range stableDiffusionList {
		if stableDiffusion.ImageId != 0 {
			var file example.ExaFileUploadAndDownload
			err = global.DB.Model(&example.ExaFileUploadAndDownload{}).Where("id = ?", stableDiffusion.ImageId).Find(&file).Error
			if err == nil {
				stableDiffusion.Url = file.Url
			}
		}
	}
	return &stableDiffusionList, total, err
}

// CreateStableDiffusionLoras 创建loras
func (s *StableDiffusionLorasService) CreateStableDiffusionLoras(stableDiffusionLoras system.StableDiffusionLoras) error {
	return global.DB.Create(&stableDiffusionLoras).Error
}
