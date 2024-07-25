package system

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/system"
)

type StableDiffusionService struct{}

// GetStableDiffusionConfigList 获取stable-diffusion列表
func (s *StableDiffusionService) GetStableDiffusionConfigList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.StableDiffusion{})
	err = db.Count(&total).Error
	var stableDiffusionList []system.StableDiffusion
	if err != nil {
		return stableDiffusionList, total, err
	}
	db = db.Limit(limit).Offset(offset)
	OrderStr := "id desc"
	//if order != "" {
	//	orderMap := make(map[string]bool, 5)
	//	orderMap["id"] = true
	//	orderMap["path"] = true
	//	orderMap["api_group"] = true
	//	orderMap["description"] = true
	//	orderMap["method"] = true
	//	if !orderMap[order] {
	//		err = fmt.Errorf("非法的排序字段: %v", order)
	//		return apiList, total, err
	//	}
	//	OrderStr = order
	//	if desc {
	//		OrderStr = order + " desc"
	//	}
	//}
	err = db.Order(OrderStr).Find(&stableDiffusionList).Error
	return stableDiffusionList, total, err
}

// CreateStableDiffusionConfig 新增stable-diffusion配置
func (s *StableDiffusionService) CreateStableDiffusionConfig(config system.StableDiffusion) (err error) {
	err = global.DB.Create(&config).Error
	return err
}
