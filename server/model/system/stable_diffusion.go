package system

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github/stable-diffusion-go/server/global"
)

type LoraIDsSlice []int

func (v *LoraIDsSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan Array value:", value))
	}
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, v)
	}
	*v = make([]int, 0)
	return nil
}

func (v LoraIDsSlice) Value() (driver.Value, error) {
	if v == nil {
		return "[]", nil
	}
	return json.Marshal(v)
}

type StableDiffusion struct {
	global.Model
	Url     string       `json:"url" gorm:"comment:stable_diffusion url"`      // url
	Width   int          `json:"width" gorm:"comment:图片宽度;default:512"`        // 图片宽度
	Height  int          `json:"height" gorm:"comment:图片高度;default:512"`       // 图片高度
	LoraIds LoraIDsSlice `json:"loraIds" gorm:"comment:lora的Id;type:longtext"` // lora
}

func (StableDiffusion) TableName() string {
	return "stable_diffusion"
}
