package system

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github/stable-diffusion-go/server/global"
)

type LoraIDsSlice []int

func (l LoraIDsSlice) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *LoraIDsSlice) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, l)
}

type StableDiffusion struct {
	global.Model
	Url     string       `json:"url" gorm:"comment:stable_diffusion url"` // url
	Width   int          `json:"width" gorm:"comment:图片宽度;default:512"`   // 图片宽度
	Height  int          `json:"height" gorm:"comment:图片高度;default:512"`  // 图片高度
	LoraIds LoraIDsSlice `json:"loraIds" gorm:"comment:lora的Id"`          // lora
}

func (StableDiffusion) TableName() string {
	return "stable_diffusion"
}
