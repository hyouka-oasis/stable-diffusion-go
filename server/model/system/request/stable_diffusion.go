package request

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
)

type StableDiffusionQueryParams struct {
	request.PageInfo
}

type StableDiffusionCreateParams struct {
	global.Model
	Url     string `json:"url"`    // url
	Width   int    `json:"width"`  // 图片宽度
	Height  int    `json:"height"` // 图片高度
	LoraIds []uint `json:"loraIds"`
}
